package grpc

import (
	"context"
	"log/slog"
	"runtime/debug"

	"github.com/bufbuild/protovalidate-go"
	"github.com/ericbutera/amalgam/internal/logger"
	"github.com/ericbutera/amalgam/services/rpc/internal/server/grpc/interceptors"
	"github.com/ericbutera/amalgam/services/rpc/internal/server/observability"
	grpcprom "github.com/grpc-ecosystem/go-grpc-middleware/providers/prometheus"
	"github.com/grpc-ecosystem/go-grpc-middleware/v2/interceptors/logging"
	protovalidate_middleware "github.com/grpc-ecosystem/go-grpc-middleware/v2/interceptors/protovalidate"
	"github.com/grpc-ecosystem/go-grpc-middleware/v2/interceptors/recovery"
	"github.com/prometheus/client_golang/prometheus"
	"go.opentelemetry.io/contrib/instrumentation/google.golang.org/grpc/otelgrpc"
	"go.opentelemetry.io/otel/trace"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/reflection"
	"google.golang.org/grpc/status"
)

func NewServer(srvMetrics *grpcprom.ServerMetrics, feedMetrics *observability.FeedMetrics) (*grpc.Server, error) {
	logger, logOpts := newLogger()
	recoveryHandler := grpcPanicRecoveryHandler(feedMetrics)

	validator, err := protovalidate.New()
	if err != nil {
		return nil, err
	}

	// Check grpc ecosystem before creating custom interceptors.
	// https://github.com/grpc-ecosystem/go-grpc-middleware/tree/main/interceptors
	// TODO: research effort to reward for consistent validationErr responses in protovalidate middleware
	srv := grpc.NewServer(
		grpc.StatsHandler(otelgrpc.NewServerHandler()),
		grpc.ChainUnaryInterceptor(
			// Order matters e.g. tracing interceptor have to create span first for the later exemplars to work.
			srvMetrics.UnaryServerInterceptor(grpcprom.WithExemplarFromContext(exemplarFromContext)),
			logging.UnaryServerInterceptor(logger, logOpts...),
			recovery.UnaryServerInterceptor(recovery.WithRecoveryHandler(recoveryHandler)),
			protovalidate_middleware.UnaryServerInterceptor(validator),
			interceptors.UnaryMetricMiddlewareHandler(feedMetrics),
		),
		grpc.ChainStreamInterceptor(
			srvMetrics.StreamServerInterceptor(grpcprom.WithExemplarFromContext(exemplarFromContext)),
			logging.StreamServerInterceptor(logger, logOpts...),
			recovery.StreamServerInterceptor(recovery.WithRecoveryHandler(recoveryHandler)),
			protovalidate_middleware.StreamServerInterceptor(validator),
			interceptors.StreamMetricMiddlewareHandler(feedMetrics),
		),
	)

	srvMetrics.InitializeMetrics(srv)
	reflection.Register(srv)

	return srv, nil
}

func newLogger() (logging.Logger, []logging.Option) {
	logTraceID := func(ctx context.Context) logging.Fields {
		if span := trace.SpanContextFromContext(ctx); span.IsSampled() {
			return logging.Fields{"traceID", span.TraceID().String()}
		}
		return nil
	}

	opts := []logging.Option{
		logging.WithLogOnEvents(logging.StartCall, logging.FinishCall),
		logging.WithFieldsFromContext(logTraceID),
	}
	return interceptorLogger(logger.New()), opts
}

// InterceptorLogger adapts slog logger to interceptor logger.
// This code is simple enough to be copied and not imported.
func interceptorLogger(l *slog.Logger) logging.Logger {
	return logging.LoggerFunc(func(ctx context.Context, lvl logging.Level, msg string, fields ...any) {
		// well this is ridiculous. ignore health check logs.
		for i := 0; i < len(fields); i += 2 {
			if key, ok := fields[i].(string); ok && key == "grpc.service" {
				if val, ok := fields[i+1].(string); ok && val == "grpc.health.v1.Health" {
					return
				}
				break
			}
		}
		l.Log(ctx, slog.Level(lvl), msg, fields...)
	})
}

func exemplarFromContext(ctx context.Context) prometheus.Labels {
	if span := trace.SpanContextFromContext(ctx); span.IsSampled() {
		return prometheus.Labels{"traceID": span.TraceID().String()}
	}
	return nil
}

func grpcPanicRecoveryHandler(feedMetrics *observability.FeedMetrics) recovery.RecoveryHandlerFunc {
	var panicsTotal prometheus.Counter
	if feedMetrics != nil && feedMetrics.PanicsTotal != nil {
		panicsTotal = feedMetrics.PanicsTotal
	}

	return func(p any) (err error) {
		if panicsTotal != nil {
			panicsTotal.Inc()
		}
		slog.Error("recovered from panic", "panic", p, "stack", debug.Stack())
		return status.Errorf(codes.Internal, "%s", p)
	}
}
