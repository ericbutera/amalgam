package app

import (
	"time"

	"github.com/ericbutera/amalgam/data-pipeline/temporal/feed/internal/feeds"
	"go.temporal.io/sdk/workflow"
)

func FeedWorkflow(ctx workflow.Context, feeds []feeds.Feed) error {
	ctx = workflow.WithActivityOptions(ctx, workflow.ActivityOptions{
		StartToCloseTimeout: time.Minute,
	})

	var a *Activities

	for _, feed := range feeds {
		// TODO: concurrency
		// https://github.com/temporalio/samples-go/blob/main/goroutine/goroutine_workflow.go
		// workflow.Go(ctx, func(gCtx workflow.Context) {...}
		// workflow.Await(ctx, func() bool { return err != nil || len(results) == parallelism })
		var rssFile string
		err := workflow.ExecuteActivity(ctx, a.DownloadActivity, feed.ID, feed.Url).Get(ctx, &rssFile)
		if err != nil {
			return err
		}

		var articlesFile string
		err = workflow.ExecuteActivity(ctx, a.ParseActivity, feed.ID, rssFile).Get(ctx, &articlesFile)
		if err != nil {
			return err
		}

		err = workflow.ExecuteActivity(ctx, a.SaveActivity, feed.ID, articlesFile).Get(ctx, nil)
		if err != nil {
			return err
		}
	}

	return nil
}
