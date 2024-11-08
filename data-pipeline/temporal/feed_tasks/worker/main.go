package main

import (
	"log/slog"
	"net/http"
	"os"

	"github.com/Khan/genqlient/graphql"
	"github.com/ericbutera/amalgam/data-pipeline/temporal/feed_tasks"
	"github.com/ericbutera/amalgam/data-pipeline/temporal/internal/client"
	"github.com/ericbutera/amalgam/pkg/config"
	"github.com/samber/lo"
	"go.temporal.io/sdk/worker"
)

func main() {
	config := lo.Must(config.NewFromEnv[feed_tasks.Config]())
	graphClient := graphql.NewClient(config.GraphHost, &http.Client{})
	activities := feed_tasks.NewActivities(graphClient)

	client := lo.Must(client.NewTemporalClient(config.TemporalHost))
	defer client.Close()

	w := worker.New(client, config.TaskQueue, worker.Options{})
	w.RegisterWorkflow(feed_tasks.GenerateFeedsWorkflow)
	w.RegisterActivity(activities)

	err := w.Run(worker.InterruptCh())
	if err != nil {
		slog.Error("unable to start worker", "error", err)
		os.Exit(1)
	}
}
