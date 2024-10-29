package main

import (
	"log/slog"
	"os"

	app "github.com/ericbutera/amalgam/data-pipeline/temporal/feed"
	"github.com/ericbutera/amalgam/data-pipeline/temporal/feed/internal/bucket"
	"github.com/ericbutera/amalgam/data-pipeline/temporal/feed/internal/config"
	"github.com/ericbutera/amalgam/data-pipeline/temporal/feed/internal/feeds"
	cfg "github.com/ericbutera/amalgam/pkg/config"
	"github.com/samber/lo"

	"go.temporal.io/sdk/worker"
)

func main() {
	config := lo.Must(cfg.NewFromEnv[config.Config]())

	bucketClient := lo.Must(bucket.NewMinioClient(config))
	feeds := lo.Must(feeds.NewFeeds(config))
	defer feeds.Close()
	a := app.NewActivities(bucketClient, feeds)

	c := lo.Must(app.NewTemporalClient(config.TemporalHost))
	defer c.Close()
	w := worker.New(c, config.TaskQueue, worker.Options{})
	w.RegisterWorkflow(app.FeedWorkflow)
	w.RegisterActivity(a)

	err := w.Run(worker.InterruptCh())
	if err != nil {
		slog.Error("unable to start worker", "error", err)
		os.Exit(1)
	}
}
