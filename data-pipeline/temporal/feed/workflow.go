package app

import (
	"errors"
	"time"

	"github.com/ericbutera/amalgam/data-pipeline/temporal/feed/internal/config"
	"github.com/ericbutera/amalgam/data-pipeline/temporal/feed/internal/feeds"
	cfg "github.com/ericbutera/amalgam/pkg/config"
	"github.com/samber/lo"
	"go.temporal.io/sdk/temporal"
	"go.temporal.io/sdk/workflow"
)

const FetchFeedWorkflowID = "fetch-feed"
const Concurrency = 100

var (
	retryPolicy = temporal.RetryPolicy{
		MaximumAttempts: 1,
	}
)

func FeedWorkflow(ctx workflow.Context, feedId string, url string) error {
	ctx = workflow.WithActivityOptions(ctx, workflow.ActivityOptions{
		StartToCloseTimeout: time.Minute,
		RetryPolicy:         &retryPolicy,
	})

	var a *Activities

	// fetch feeds into an activity, save historical feed list to bucket
	var rssFile string
	err := workflow.ExecuteActivity(ctx, a.DownloadActivity, feedId, url).Get(ctx, &rssFile)
	if err != nil {
		return err
	}
	var articlesFile string
	err = workflow.ExecuteActivity(ctx, a.ParseActivity, feedId, rssFile).Get(ctx, &articlesFile)
	if err != nil {
		return err
	}
	err = workflow.ExecuteActivity(ctx, a.SaveActivity, feedId, articlesFile).Get(ctx, nil)
	if err != nil {
		return err
	}
	return nil
}

func FetchFeedsWorkflow(ctx workflow.Context) error {
	ctx = workflow.WithActivityOptions(ctx, workflow.ActivityOptions{
		StartToCloseTimeout: 1 * time.Minute,
		RetryPolicy:         &retryPolicy,
	})

	config := lo.Must(cfg.NewFromEnv[config.Config]())
	feeds := lo.Must(feeds.NewFeeds(config))
	defer feeds.Close()
	urls := lo.Must(feeds.GetFeeds())

	var errs []error
	semaphore := workflow.NewSemaphore(ctx, Concurrency)

	for _, feed := range urls {
		if err := semaphore.Acquire(ctx, 1); err != nil {
			return err
		}
		workflow.Go(ctx, func(ctx workflow.Context) {
			defer semaphore.Release(1)
			gCtx := workflow.WithChildOptions(ctx, workflow.ChildWorkflowOptions{
				WorkflowID:  FetchFeedWorkflowID,
				RetryPolicy: &retryPolicy,
			})

			var result string
			err := workflow.ExecuteChildWorkflow(gCtx, FeedWorkflow, feed.ID, feed.Url).Get(gCtx, &result)
			if err != nil {
				errs = append(errs, err)
			}
		})
	}

	_ = workflow.Await(ctx, func() bool {
		// wait for capacity to be fully released
		return semaphore.TryAcquire(ctx, Concurrency)
	})

	if len(errs) > 0 {
		return errors.Join(errs...)
	}

	return nil
}
