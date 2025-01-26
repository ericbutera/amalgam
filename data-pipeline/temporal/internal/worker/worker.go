package worker

import (
	"context"

	clientHelper "github.com/ericbutera/amalgam/data-pipeline/temporal/internal/client"
	"github.com/ericbutera/amalgam/pkg/config/env"
	"github.com/ericbutera/amalgam/pkg/otel"
	"github.com/samber/lo"
	"go.temporal.io/sdk/client"
	"go.temporal.io/sdk/worker"
)

type Config struct {
	TaskQueue string `mapstructure:"task_queue"`
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
	IgnoredSpanNames []string `mapstructure:"ignored_span_names"`
}

func NewOtel(ctx context.Context) (func(context.Context) error, error) {
	config, err := env.New[OtelConfig]()
	if err != nil {
		return nil, err
	}
	shutdown, err := otel.Setup(ctx, config.IgnoredSpanNames)
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
		lo.Must0(shutdown(ctx))
		client.Close()
	}
	return w, closers, nil
}
