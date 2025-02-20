package main

import (
	"context"

	app "github.com/ericbutera/amalgam/data-pipeline/temporal/feed_fetch"
	workerHelper "github.com/ericbutera/amalgam/data-pipeline/temporal/internal/worker"
	"github.com/samber/lo"
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
}
