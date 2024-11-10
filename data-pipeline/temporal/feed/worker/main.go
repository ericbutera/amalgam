package main

import (
	"context"
	"log/slog"
	"os"

	app "github.com/ericbutera/amalgam/data-pipeline/temporal/feed"
	"github.com/ericbutera/amalgam/data-pipeline/temporal/feed/internal/config"
	"github.com/ericbutera/amalgam/data-pipeline/temporal/feed/internal/transforms"
	"github.com/ericbutera/amalgam/data-pipeline/temporal/internal/bucket"
	"github.com/ericbutera/amalgam/data-pipeline/temporal/internal/client"
	"github.com/ericbutera/amalgam/data-pipeline/temporal/internal/feeds"
	helper "github.com/ericbutera/amalgam/data-pipeline/temporal/internal/worker"
	"github.com/ericbutera/amalgam/internal/http/fetch"
	"github.com/ericbutera/amalgam/pkg/config/env"
	"github.com/ericbutera/amalgam/pkg/otel"
	"github.com/samber/lo"
	"go.temporal.io/sdk/worker"
)

func main() {
	config := lo.Must(env.New[config.Config]())

	transforms := transforms.New()

	fetcher := lo.Must(fetch.New())

	bucketConfig := lo.Must(bucket.NewConfig(config))
	bucketClient := lo.Must(bucket.NewMinio(bucketConfig))

	feeds := lo.Must(feeds.NewFeeds(config.RpcHost, config.RpcInsecure))
	defer feeds.Close()

	a := app.NewActivities(transforms, fetcher, bucketClient, feeds)

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
