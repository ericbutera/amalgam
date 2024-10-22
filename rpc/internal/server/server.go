package server

import (
	"context"
	"fmt"
	"log/slog"
	"net"
	"net/http"
	"os"
	"runtime/debug"
	"syscall"

	"github.com/ericbutera/amalgam/internal/db"
	"github.com/ericbutera/amalgam/internal/service"
	pb "github.com/ericbutera/amalgam/pkg/rpc/proto"
	"github.com/grpc-ecosystem/go-grpc-middleware/v2/interceptors/logging"
	"github.com/grpc-ecosystem/go-grpc-middleware/v2/interceptors/recovery"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"gorm.io/gorm"

	grpcprom "github.com/grpc-ecosystem/go-grpc-middleware/providers/prometheus"
	"go.opentelemetry.io/contrib/instrumentation/google.golang.org/grpc/otelgrpc"
	"go.opentelemetry.io/otel/trace"

	"github.com/oklog/run"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/reflection"
	"google.golang.org/grpc/status"
)

type Server struct {
	metricAddress string
	port          string
	srv           *grpc.Server
	listener      net.Listener
	panicsTotal   prometheus.Counter
	metricSrv     *http.Server
	service       *service.Service
	pb.UnimplementedFeedServiceServer
}

func (s *Server) Serve() error {
	g := &run.Group{}
	g.Add(func() error {
		slog.Info("launching server", "port", s.port)
		return s.srv.Serve(s.listener)
	}, func(err error) {
		s.srv.GracefulStop()
		s.srv.Stop()
	})

	g.Add(func() error {
		slog.Info("serving metrics", "addr", s.metricSrv.Addr)
		return s.metricSrv.ListenAndServe()
	}, func(error) {
		if err := s.metricSrv.Close(); err != nil {
			slog.Error("failed to stop web server", "err", err)
		}
	})

	g.Add(run.SignalHandler(context.Background(), syscall.SIGINT, syscall.SIGTERM))

	if err := g.Run(); err != nil {
		return err
	}
	return nil
}

type Option func(*Server) error

func WithPort(port string) Option {
	return func(s *Server) error {
		s.port = port
		return nil
	}
}
func WithMetricAddress(addr string) Option {
	return func(s *Server) error {
		s.metricAddress = addr
		return nil
	}
}
func WithListener(lis net.Listener) Option {
	return func(s *Server) error {
		s.listener = lis
		return nil
	}
}
func WithDb(db *gorm.DB) Option {
	return func(s *Server) error {
		s.service = service.New(db)
		return nil
	}
}
func WithService(service *service.Service) Option {
	return func(s *Server) error {
		s.service = service
		return nil
	}
}

func New(opts ...Option) (*Server, error) {
	server := Server{}

	registry := prometheus.NewRegistry()
	srvMetrics := newServerMetrics()
	registry.MustRegister(srvMetrics)
	server.newPromMetrics(registry)
	logger, logOpts := newServerLogger()

	server.srv = grpc.NewServer(
		grpc.StatsHandler(otelgrpc.NewServerHandler()),
		grpc.ChainUnaryInterceptor(
			// Order matters e.g. tracing interceptor have to create span first for the later exemplars to work.
			srvMetrics.UnaryServerInterceptor(grpcprom.WithExemplarFromContext(exemplarFromContext)),
			logging.UnaryServerInterceptor(logger, logOpts...),
			recovery.UnaryServerInterceptor(recovery.WithRecoveryHandler(server.grpcPanicRecoveryHandler)),
		),
		grpc.ChainStreamInterceptor(
			srvMetrics.StreamServerInterceptor(grpcprom.WithExemplarFromContext(exemplarFromContext)),
			logging.StreamServerInterceptor(logger, logOpts...),
			recovery.StreamServerInterceptor(recovery.WithRecoveryHandler(server.grpcPanicRecoveryHandler)),
		),
	)

	for _, opt := range opts {
		if err := opt(&server); err != nil {
			return nil, err
		}
	}

	if server.port == "" {
		return nil, fmt.Errorf("port is required")
	}
	if server.listener == nil {
		listener, err := net.Listen("tcp", fmt.Sprintf(":%s", server.port))
		if err != nil {
			return nil, err
		}
		server.listener = listener
	}
	if server.service == nil {
		db, err := db.NewFromEnv()
		if err != nil {
			return nil, err
		}
		server.service = service.New(db)
	}

	server.metricSrv = newMetricsServer(registry, server.metricAddress)
	pb.RegisterFeedServiceServer(server.srv, &Server{})
	reflection.Register(server.srv)
	srvMetrics.InitializeMetrics(server.srv)

	return &server, nil
}

func newServerLogger() (logging.Logger, []logging.Option) {
	logger := slog.New(slog.NewTextHandler(os.Stdout, nil)) // TODO: why does go insist on using stderr?
	opts := []logging.Option{
		logging.WithLogOnEvents(logging.StartCall, logging.FinishCall),
		// Add any other option (check functions starting with logging.With).
	}
	return interceptorLogger(logger), opts
}

func newServerMetrics() *grpcprom.ServerMetrics {
	return grpcprom.NewServerMetrics(
		grpcprom.WithServerHandlingTimeHistogram(
			grpcprom.WithHistogramBuckets([]float64{0.001, 0.01, 0.1, 0.3, 0.6, 1, 3, 6, 9, 20, 30, 60, 90, 120}),
		),
	)
}

func newMetricsServer(registry *prometheus.Registry, address string) *http.Server {
	srv := &http.Server{Addr: address}
	m := http.NewServeMux()
	m.Handle("/metrics", promhttp.HandlerFor(
		registry,
		promhttp.HandlerOpts{
			EnableOpenMetrics: true, // Opt into OpenMetrics e.g. to support exemplars.
		},
	))
	srv.Handler = m
	return srv
}

func (s *Server) newPromMetrics(reg prometheus.Registerer) {
	s.panicsTotal = promauto.With(reg).NewCounter(prometheus.CounterOpts{
		Name: "grpc_req_panics_recovered_total",
		Help: "Total number of gRPC requests recovered from internal panic.",
	})
}

func exemplarFromContext(ctx context.Context) prometheus.Labels {
	if span := trace.SpanContextFromContext(ctx); span.IsSampled() {
		return prometheus.Labels{"traceID": span.TraceID().String()}
	}
	return nil
}

func (s *Server) grpcPanicRecoveryHandler(p any) (err error) {
	s.panicsTotal.Inc()
	slog.Error("recovered from panic", "panic", p, "stack", debug.Stack())
	return status.Errorf(codes.Internal, "%s", p)
}

// InterceptorLogger adapts slog logger to interceptor logger.
// This code is simple enough to be copied and not imported.
func interceptorLogger(l *slog.Logger) logging.Logger {
	return logging.LoggerFunc(func(ctx context.Context, lvl logging.Level, msg string, fields ...any) {
		l.Log(ctx, slog.Level(lvl), msg, fields...)
	})
}
