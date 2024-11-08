package main

import (
	"context"
	"log/slog"
	"os"

	app "github.com/ericbutera/amalgam/data-pipeline/temporal/feed"
	"github.com/ericbutera/amalgam/data-pipeline/temporal/feed/internal/config"
	"github.com/ericbutera/amalgam/data-pipeline/temporal/internal/bucket"
	"github.com/ericbutera/amalgam/data-pipeline/temporal/internal/client"
	"github.com/ericbutera/amalgam/data-pipeline/temporal/internal/feeds"
	helper "github.com/ericbutera/amalgam/data-pipeline/temporal/internal/worker"
	cfg "github.com/ericbutera/amalgam/pkg/config"
	otel "github.com/ericbutera/amalgam/pkg/otel"
	"github.com/samber/lo"
	"go.temporal.io/sdk/worker"
)

func main() {
	config := lo.Must(cfg.NewFromEnv[config.Config]())
	bucketConfig := lo.Must(bucket.NewConfig(config))
	bucketClient := lo.Must(bucket.NewMinioClient(bucketConfig))

	feeds := lo.Must(feeds.NewFeeds(config.RpcHost, config.RpcInsecure))
	defer feeds.Close()
	a := app.NewActivities(bucketClient, feeds)

	ctx := context.Background()
	shutdown := lo.Must(otel.Setup(ctx))
	defer helper.HandleShutdown(ctx, shutdown)

	client := lo.Must(client.NewTemporalClient(config.TemporalHost))
	defer client.Close()

	w := worker.New(client, config.TaskQueue, worker.Options{
		Interceptors: helper.NewInterceptors(otel.Tracer),
	})
	w.RegisterWorkflow(app.FetchFeedsWorkflow)
	w.RegisterWorkflow(app.FeedWorkflow)
	w.RegisterActivity(a)

	err := w.Run(worker.InterruptCh())
	if err != nil {
		slog.Error("unable to start worker", "error", err)
		os.Exit(1)
	}
}
