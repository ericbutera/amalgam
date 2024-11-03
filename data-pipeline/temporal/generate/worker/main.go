package main

import (
	"log/slog"
	"net/http"
	"os"

	"github.com/Khan/genqlient/graphql"
	"github.com/ericbutera/amalgam/data-pipeline/temporal/generate"
	"github.com/ericbutera/amalgam/data-pipeline/temporal/internal/client"
	"github.com/ericbutera/amalgam/pkg/config"
	"github.com/samber/lo"

	"go.temporal.io/sdk/worker"
)

const GraphHost = "localhost:8082" // TODO: find a cleaner way to run these "local scripts".. i don't want to use .env all over the place

func main() {
	config := lo.Must(config.NewFromEnv[generate.Config]())
	config.GraphHost = GraphHost
	graphClient := graphql.NewClient(config.GraphHost, &http.Client{})
	a := generate.NewActivities(graphClient)

	client := lo.Must(client.NewTemporalClient(config.TemporalHost))
	defer client.Close()

	w := worker.New(client, config.TaskQueue, worker.Options{})
	w.RegisterWorkflow(generate.GenerateFeedsWorkflow)
	w.RegisterActivity(a)

	err := w.Run(worker.InterruptCh())
	if err != nil {
		slog.Error("unable to start worker", "error", err)
		os.Exit(1)
	}
}
