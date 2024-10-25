package main

import (
	"context"
	"log/slog"
	"os"

	app "github.com/ericbutera/amalgam/data-pipeline/temporal/feed"
	"github.com/ericbutera/amalgam/data-pipeline/temporal/feed/internal/config"
	"github.com/ericbutera/amalgam/data-pipeline/temporal/feed/internal/feeds"
	cfg "github.com/ericbutera/amalgam/pkg/config"
	"github.com/samber/lo"
	"go.temporal.io/sdk/client"
)

func main() {
	config := lo.Must(cfg.NewFromEnv[config.Config]())

	feeds := lo.Must(feeds.NewFeeds(config))
	defer feeds.Close()

	urls := lo.Must(feeds.GetFeeds())

	c := lo.Must(app.NewTemporalClient(config.TemporalHost))
	defer c.Close()

	opts := client.StartWorkflowOptions{
		TaskQueue: config.TaskQueue,
	}
	we, err := c.ExecuteWorkflow(
		context.Background(),
		opts,
		app.FeedWorkflow,
		urls,
	)
	if err != nil {
		slog.Error("unable to execute workflow", "error", err)
		os.Exit(1)
	}

	slog.Info("started workflow", "workflow_id", we.GetID())
}
