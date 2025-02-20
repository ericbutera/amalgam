package main

import (
	"log/slog"
	"net/http"
	"os"

	"github.com/Khan/genqlient/graphql"
	"github.com/ericbutera/amalgam/data-pipeline/temporal/feed_tasks"
	"github.com/ericbutera/amalgam/data-pipeline/temporal/internal/client"
	"github.com/ericbutera/amalgam/pkg/config/env"
	"github.com/samber/lo"
	"go.temporal.io/sdk/worker"
)

func main() {
	// TODO: use worker helper for observability
	config := lo.Must(env.New[feed_tasks.Config]())
	graphClient := graphql.NewClient(config.GraphHost, &http.Client{})

	client := lo.Must(client.NewTemporalClient(config.TemporalHost))
	defer client.Close()

	activities := feed_tasks.NewActivities(graphClient, client)

	w := worker.New(client, config.TaskQueue, worker.Options{})
	w.RegisterWorkflow(feed_tasks.GenerateFeedsWorkflow)
	w.RegisterWorkflow(feed_tasks.RefreshFeedsWorkflow)
	w.RegisterActivity(activities)

	err := w.Run(worker.InterruptCh())
	if err != nil {
		slog.Error("unable to start worker", "error", err)
		os.Exit(1) //nolint: gocritic
	}
}
