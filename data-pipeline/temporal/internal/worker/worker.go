package worker

import (
	"context"
	"log/slog"
	"strings"

	clientHelper "github.com/ericbutera/amalgam/data-pipeline/temporal/internal/client"
	"github.com/ericbutera/amalgam/pkg/config/env"
	"github.com/ericbutera/amalgam/pkg/otel"
	"go.temporal.io/sdk/client"
	"go.temporal.io/sdk/worker"
)

type Config struct {
	TaskQueue string `env:"TASK_QUEUE"`
}

func NewFromEnv(client client.Client) (worker.Worker, error) {
	config, err := env.New[Config]()
	if err != nil {
		return nil, err
	}
	w := worker.New(client, config.TaskQueue, worker.Options{
		Interceptors: NewInterceptors(otel.Tracer),
	})
	return w, nil
}

type OtelConfig struct {
	IgnoredSpanNames string `env:"IGNORED_SPAN_NAMES"`
}

func NewOtel(ctx context.Context) (func(context.Context) error, error) {
	config, err := env.New[OtelConfig]()
	if err != nil {
		return nil, err
	}
	shutdown, err := otel.Setup(ctx, strings.Split(config.IgnoredSpanNames, ","))
	if err != nil {
		return nil, err
	}
	return shutdown, nil
}

func New(ctx context.Context) (worker.Worker, func(), error) {
	client, err := clientHelper.NewTemporalClientFromEnv()
	if err != nil {
		return nil, nil, err
	}
	shutdown, err := NewOtel(ctx)
	if err != nil {
		return nil, nil, err
	}
	w, err := NewFromEnv(client)
	if err != nil {
		return nil, nil, err
	}
	closers := func() {
		slog.Info("Shutting down worker")
		if err := shutdown(ctx); err != nil {
			slog.Error("shutdown error", "error", err)
		}
		client.Close()
	}
	return w, closers, nil
}
