package main

import (
	"context"
	"log/slog"
	"os"

	app "github.com/ericbutera/amalgam/data-pipeline/temporal/feed"
	"github.com/ericbutera/amalgam/data-pipeline/temporal/feed/internal/bucket"
	"github.com/ericbutera/amalgam/data-pipeline/temporal/feed/internal/config"
	"github.com/ericbutera/amalgam/data-pipeline/temporal/feed/internal/feeds"
	cfg "github.com/ericbutera/amalgam/pkg/config"
	otel "github.com/ericbutera/amalgam/pkg/otel"
	"github.com/samber/lo"
	"go.opentelemetry.io/otel/trace"

	"go.temporal.io/sdk/contrib/opentelemetry"
	"go.temporal.io/sdk/interceptor"
	"go.temporal.io/sdk/worker"
)

func main() {
	config := lo.Must(cfg.NewFromEnv[config.Config]())

	bucketClient := lo.Must(bucket.NewMinioClient(config))
	feeds := lo.Must(feeds.NewFeeds(config))
	defer feeds.Close()
	a := app.NewActivities(bucketClient, feeds)

	ctx := context.Background()
	shutdown := lo.Must(otel.Setup(ctx))
	defer handleShutdown(ctx, shutdown)

	c := lo.Must(app.NewTemporalClient(config.TemporalHost))
	defer c.Close()

	w := worker.New(c, config.TaskQueue, worker.Options{
		Interceptors: newInterceptors(otel.Tracer),
	})
	w.RegisterWorkflow(app.FeedWorkflow)
	w.RegisterActivity(a)

	err := w.Run(worker.InterruptCh())
	if err != nil {
		slog.Error("unable to start worker", "error", err)
		os.Exit(1)
	}
}

func newInterceptors(tracer trace.Tracer) []interceptor.WorkerInterceptor {
	traceInterceptor := lo.Must(opentelemetry.NewTracer(opentelemetry.TracerOptions{
		Tracer: tracer,
	}))
	return []interceptor.WorkerInterceptor{
		interceptor.NewTracingInterceptor(traceInterceptor),
	}
}

func handleShutdown(ctx context.Context, shutdown func(context.Context) error) {
	slog.Info("shutting down otel")
	if err := shutdown(ctx); err != nil {
		slog.Error("failed to shutdown OpenTelemetry", "error", err)
	}
}
