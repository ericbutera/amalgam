package main

import (
	"context"
	"log/slog"
	"net/http"
	"os"

	"github.com/Khan/genqlient/graphql"
	"github.com/ericbutera/amalgam/internal/temporal/client"
	"github.com/ericbutera/amalgam/internal/temporal/feed_tasks"
	workerHelper "github.com/ericbutera/amalgam/internal/temporal/worker"
	"github.com/ericbutera/amalgam/pkg/config/env"
	"github.com/samber/lo"
	"go.temporal.io/sdk/worker"
)

func main() {
	ctx := context.Background()
	config := lo.Must(env.New[feed_tasks.Config]())
	graphClient := graphql.NewClient(config.GraphHost, &http.Client{})

	client := lo.Must(client.NewTemporalClient(config.TemporalHost))
	defer client.Close()

	shutdown := lo.Must(workerHelper.NewOtel(ctx))
	defer func() {
		if err := shutdown(ctx); err != nil {
			slog.Error("telemetry shutdown error", "error", err)
		}
	}()

	activities := feed_tasks.NewActivities(graphClient, client)

	w := lo.Must(workerHelper.NewFromEnv(client))
	w.RegisterWorkflow(feed_tasks.GenerateFeedsWorkflow)
	w.RegisterWorkflow(feed_tasks.GenerateFeedsWorkflowV2)
	w.RegisterWorkflow(feed_tasks.RefreshFeedsWorkflow)
	w.RegisterWorkflow(feed_tasks.AddFeedWorkflow)
	w.RegisterWorkflow(feed_tasks.AddFeedWorkflowV2)
	w.RegisterActivity(activities)

	err := w.Run(worker.InterruptCh())
	if err != nil {
		slog.Error("unable to start worker", "error", err)
		os.Exit(1) //nolint: gocritic
	}
}
