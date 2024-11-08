package client

// TODO move to pkg/clients/rpc

import (
	"context"
	"log/slog"
	"os"
	"time"

	pb "github.com/ericbutera/amalgam/pkg/feeds/v1"
	"github.com/grpc-ecosystem/go-grpc-middleware/v2/interceptors/logging"
	"github.com/grpc-ecosystem/go-grpc-middleware/v2/interceptors/timeout"
	"go.opentelemetry.io/contrib/instrumentation/google.golang.org/grpc/otelgrpc"
	"go.opentelemetry.io/otel/trace"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

// TODO: combine common boilerplate code from rpc client & server

type Rpc struct {
	Client pb.FeedServiceClient
	Conn   *grpc.ClientConn
}

// TODO: find a cleaner way to handle Client and Conn. this return value is confusing
func NewClient(target string, useInsecure bool) (*Rpc, error) {
	logger, logOpts := newLogger()

	dialOpts := []grpc.DialOption{
		grpc.WithStatsHandler(otelgrpc.NewClientHandler()),
		grpc.WithChainUnaryInterceptor(
			timeout.UnaryClientInterceptor(10*time.Second),
			logging.UnaryClientInterceptor(logger, logOpts...),
		),
		grpc.WithChainStreamInterceptor(
			logging.StreamClientInterceptor(logger, logOpts...),
		),
	}

	if useInsecure {
		dialOpts = append(dialOpts, grpc.WithTransportCredentials(insecure.NewCredentials()))
	}

	conn, err := grpc.NewClient(
		target,
		dialOpts...,
	)
	if err != nil {
		return nil, err
	}

	return &Rpc{
		Client: pb.NewFeedServiceClient(conn),
		Conn:   conn,
	}, nil
}

func newLogger() (logging.Logger, []logging.Option) {
	logTraceID := func(ctx context.Context) logging.Fields {
		if span := trace.SpanContextFromContext(ctx); span.IsSampled() {
			return logging.Fields{"traceID", span.TraceID().String()}
		}
		return nil
	}
	logger := slog.New(slog.NewTextHandler(os.Stdout, nil)) // TODO: why does go insist on using stderr?
	opts := []logging.Option{
		logging.WithLogOnEvents(logging.StartCall, logging.FinishCall),
		logging.WithFieldsFromContext(logTraceID),
	}
	return interceptorLogger(logger), opts
}

// InterceptorLogger adapts slog logger to interceptor logger.
// This code is simple enough to be copied and not imported.
func interceptorLogger(l *slog.Logger) logging.Logger {
	return logging.LoggerFunc(func(ctx context.Context, lvl logging.Level, msg string, fields ...any) {
		l.Log(ctx, slog.Level(lvl), msg, fields...)
	})
}
