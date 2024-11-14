package server

import (
	"context"
	"fmt"
	"log/slog"
	"net"
	"net/http"
	"syscall"

	"github.com/ericbutera/amalgam/internal/db"
	"github.com/ericbutera/amalgam/internal/service"
	"github.com/ericbutera/amalgam/pkg/config/env"
	pb "github.com/ericbutera/amalgam/pkg/feeds/v1"
	"github.com/ericbutera/amalgam/pkg/otel"
	"github.com/ericbutera/amalgam/rpc/internal/config"
	grpc_server "github.com/ericbutera/amalgam/rpc/internal/server/grpc"
	metrics_server "github.com/ericbutera/amalgam/rpc/internal/server/metrics"
	"github.com/ericbutera/amalgam/rpc/internal/server/observability"
	"github.com/oklog/run"
	"google.golang.org/grpc"
	"google.golang.org/grpc/health"
	healthpb "google.golang.org/grpc/health/grpc_health_v1"
	"gorm.io/gorm"
)

// TODO: combine common boilerplate code from rpc client & server

type Server struct {
	config    *config.Config
	grpcSrv   *grpc.Server
	metricSrv *http.Server
	db        *gorm.DB
	service   service.Service
	shutdowns []func(context.Context) error
	pb.UnimplementedFeedServiceServer
}

func (s *Server) Serve(ctx context.Context) error {
	g := &run.Group{}

	g.Add(func() error {
		slog.Info("launching server", "port", s.config.Port)
		listener, err := net.Listen("tcp", fmt.Sprintf(":%s", s.config.Port))
		if err != nil {
			return err
		}
		defer listener.Close()
		return s.grpcSrv.Serve(listener)
	}, func(err error) {
		slog.Error("shutting down server", "err", err)
		s.grpcSrv.GracefulStop()
		s.grpcSrv.Stop()

		for _, shutdown := range s.shutdowns {
			if err := shutdown(ctx); err != nil {
				slog.Error("failed to shutdown", "err", err)
			}
		}
	})

	g.Add(func() error {
		slog.Info("serving metrics", "addr", s.metricSrv.Addr)
		return s.metricSrv.ListenAndServe()
	}, func(error) {
		if err := s.metricSrv.Close(); err != nil {
			slog.Error("failed to stop web server", "err", err)
		}
	})

	g.Add(run.SignalHandler(ctx, syscall.SIGINT, syscall.SIGTERM))

	return g.Run()
}

type Option func(*Server) error

func WithDb(db *gorm.DB) Option {
	return func(s *Server) error {
		s.service = service.NewGorm(db)
		s.db = db
		return nil
	}
}

func WithDbFromEnv() Option {
	return func(s *Server) error {
		db, err := db.NewFromEnv()
		if err != nil {
			return err
		}
		return WithDb(db)(s)
	}
}

func WithService(service service.Service) Option {
	return func(s *Server) error {
		s.service = service
		return nil
	}
}

func WithConfig(data *config.Config) Option {
	return func(s *Server) error {
		s.config = data
		return nil
	}
}

func WithOtel(ctx context.Context) Option {
	return func(s *Server) error {
		shutdown, err := otel.Setup(ctx)
		if err != nil {
			return err
		}
		s.shutdowns = append(s.shutdowns, shutdown)
		return nil
	}
}

func WithMetricServer(srv *http.Server) Option {
	return func(s *Server) error {
		s.metricSrv = srv
		return nil
	}
}

func WithGrpcServer(srv *grpc.Server) Option {
	return func(s *Server) error {
		s.grpcSrv = srv
		return nil
	}
}

func New(opts ...Option) (*Server, error) {
	server := Server{}

	for _, opt := range opts {
		if err := opt(&server); err != nil {
			return nil, err
		}
	}

	o := observability.New()

	if server.config == nil {
		config, err := env.New[config.Config]()
		if err != nil {
			return nil, err
		}
		server.config = config
	}
	if server.service == nil {
		db, err := db.NewFromEnv()
		if err != nil {
			return nil, err
		}
		server.service = service.NewGorm(db)
	}
	if server.metricSrv == nil {
		server.metricSrv = metrics_server.NewServer(o.Registry, server.config.MetricAddress)
	}
	if server.grpcSrv == nil {
		server.grpcSrv = grpc_server.NewServer(o.ServerMetrics, o.FeedMetrics)
	}

	pb.RegisterFeedServiceServer(server.grpcSrv, &server)

	// TODO: extract package
	healthService := health.NewServer()
	healthpb.RegisterHealthServer(server.grpcSrv, healthService)
	healthService.SetServingStatus("rpc", healthpb.HealthCheckResponse_SERVING)

	return &server, nil
}
