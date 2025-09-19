package main

import (
	"context"
	"log/slog"
	"os"

	app "github.com/ericbutera/amalgam/data-pipeline/temporal/feed_fetch"
	workerHelper "github.com/ericbutera/amalgam/data-pipeline/temporal/internal/worker"
	"github.com/samber/lo"
	"go.temporal.io/sdk/worker"
)

func main() {
	ctx := context.Background()

	a := app.NewActivitiesFromEnv()
	defer a.Closers()

	w, closers := lo.Must2(workerHelper.New(ctx))
	defer closers()

	w.RegisterWorkflow(app.FetchFeedsWorkflow)
	w.RegisterWorkflow(app.FeedWorkflow)
	w.RegisterActivity(a)

	err := w.Run(worker.InterruptCh())
	if err != nil {
		slog.Error("unable to start worker", "error", err)
		os.Exit(1) //nolint: gocritic
	}
}
