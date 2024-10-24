package client

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

type client struct {
	Client pb.FeedServiceClient
	Conn   *grpc.ClientConn
}

func NewClient(target string, opts ...Option) (*client, error) {
	logger, logOpts := newLogger()

	creds := insecure.NewCredentials() // TODO: use secure by default!

	conn, err := grpc.NewClient(
		target,
		grpc.WithTransportCredentials(creds),
		grpc.WithStatsHandler(otelgrpc.NewClientHandler()),
		grpc.WithChainUnaryInterceptor(
			timeout.UnaryClientInterceptor(10*time.Second),
			logging.UnaryClientInterceptor(logger, logOpts...),
		),
		grpc.WithChainStreamInterceptor(
			logging.StreamClientInterceptor(logger, logOpts...),
		),
	)
	if err != nil {
		return nil, err
	}

	return &client{
		Client: pb.NewFeedServiceClient(conn),
		Conn:   conn,
	}, nil
}

type Option func(*client) error

// TODO: options!
// func WithTimeout(seconds int) Option {}

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
