package client

// TODO move to pkg/clients/rpc

import (
	"context"
	"log/slog"
	"os"
	"time"

	"github.com/ericbutera/amalgam/pkg/config/env"
	pb "github.com/ericbutera/amalgam/pkg/feeds/v1"
	"github.com/grpc-ecosystem/go-grpc-middleware/v2/interceptors/logging"
	"github.com/grpc-ecosystem/go-grpc-middleware/v2/interceptors/timeout"
	"go.opentelemetry.io/contrib/instrumentation/google.golang.org/grpc/otelgrpc"
	"go.opentelemetry.io/otel/trace"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

const Timeout = 10 * time.Second

func New(target string, useInsecure bool) (pb.FeedServiceClient, Closer, error) {
	conn, err := NewConnection(target, useInsecure)
	if err != nil {
		return nil, nil, err
	}

	return pb.NewFeedServiceClient(conn), newCloser(conn), nil
}

type Config struct {
	RpcHost     string `mapstructure:"rpc_host"`
	RpcInsecure bool   `mapstructure:"rpc_insecure"`
}

func NewFromEnv() (pb.FeedServiceClient, Closer, error) {
	config, err := env.New[Config]()
	if err != nil {
		return nil, nil, err
	}
	return New(config.RpcHost, config.RpcInsecure)
}

func NewConnection(target string, useInsecure bool) (*grpc.ClientConn, error) {
	dialOpts := defaultDialOpts()

	if useInsecure {
		dialOpts = append(dialOpts, grpc.WithTransportCredentials(insecure.NewCredentials()))
	}

	return grpc.NewClient(
		target,
		dialOpts...,
	)
}

func NewClient(conn *grpc.ClientConn) pb.FeedServiceClient {
	return pb.NewFeedServiceClient(conn)
}

// use with defer to close the connection
type Closer func() error

func newCloser(conn *grpc.ClientConn) Closer {
	return func() error { return conn.Close() }
}

func defaultDialOpts() []grpc.DialOption {
	logger, logOpts := newLogger()
	return []grpc.DialOption{
		grpc.WithStatsHandler(otelgrpc.NewClientHandler()),
		grpc.WithChainUnaryInterceptor(
			timeout.UnaryClientInterceptor(Timeout),
			logging.UnaryClientInterceptor(logger, logOpts...),
		),
		grpc.WithChainStreamInterceptor(
			logging.StreamClientInterceptor(logger, logOpts...),
		),
	}
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
