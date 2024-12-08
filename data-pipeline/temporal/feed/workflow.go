package app

import (
	"errors"
	"log/slog"
	"time"

	"github.com/ericbutera/amalgam/data-pipeline/temporal/internal/feeds"
	"go.temporal.io/sdk/temporal"
	"go.temporal.io/sdk/workflow"
)

const (
	MaxConcurrentFeeds int = 10
	Timeout                = 5 * time.Minute
)

var (
	ErrTimeout = errors.New("workflow timeout")
	ErrProcess = errors.New("error processing feeds")
)

var retryPolicy = temporal.RetryPolicy{
	MaximumAttempts: 1,
}

func FeedWorkflow(ctx workflow.Context, feedId string, url string) error {
	ctx = workflow.WithActivityOptions(ctx, workflow.ActivityOptions{
		StartToCloseTimeout: time.Minute,
		RetryPolicy:         &retryPolicy,
	})

	var a *Activities

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
	err = workflow.ExecuteActivity(ctx, a.StatsActivity, feedId).Get(ctx, nil)
	if err != nil {
		return err
	}
	return nil
}

func FetchFeedsWorkflow(ctx workflow.Context) error {
	ctx = workflow.WithActivityOptions(ctx, workflow.ActivityOptions{
		StartToCloseTimeout: Timeout,
		RetryPolicy:         &retryPolicy,
	})

	var a *Activities
	var urls []feeds.Feed
	if err := workflow.ExecuteActivity(ctx, a.GetFeedsActivity).Get(ctx, &urls); err != nil {
		return err
	}
	semaphore := workflow.NewSemaphore(ctx, int64(MaxConcurrentFeeds))

	var hasError bool
	done := 0
	for i, feed := range urls {
		// TODO: heartbeat url offset (resume on error)
		err := semaphore.Acquire(ctx, 1)
		if err != nil {
			return err
		}
		workflow.Go(ctx, func(ctx workflow.Context) {
			defer semaphore.Release(1)
			future := workflow.ExecuteChildWorkflow(ctx, FeedWorkflow, feed.ID, feed.Url)
			if err := future.Get(ctx, nil); err != nil {
				slog.Error("Failed to process feed", "i", i, "feed_id", feed.ID, "error", err)
				hasError = true
			} else {
				slog.Info("Processed feed", "i", i, "feed_id", feed.ID)
			}
			done++
		})
	}

	ok, err := workflow.AwaitWithTimeout(ctx, Timeout, func() bool {
		return hasError || len(urls) == done
	})
	if err != nil {
		return err
	}
	if !ok {
		return ErrTimeout
	}
	if hasError {
		return ErrProcess
	}
	return nil
}
