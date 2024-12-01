package app

import (
	"time"

	"github.com/ericbutera/amalgam/data-pipeline/temporal/feed/internal/config"
	"github.com/ericbutera/amalgam/data-pipeline/temporal/internal/feeds"
	"github.com/ericbutera/amalgam/pkg/config/env"
	"github.com/samber/lo"
	"go.temporal.io/sdk/temporal"
	"go.temporal.io/sdk/workflow"
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
		StartToCloseTimeout: 1 * time.Minute,
		RetryPolicy:         &retryPolicy,
	})

	// TODO: fetch feeds into an activity, save historical feed list to bucket
	config := lo.Must(env.New[config.Config]())
	feeds := lo.Must(feeds.NewFeeds(config.RpcHost, config.RpcInsecure))
	defer feeds.Close()
	urls := lo.Must(feeds.GetFeeds())

	for _, feed := range urls {
		err := workflow.ExecuteChildWorkflow(ctx, FeedWorkflow, feed.ID, feed.Url).
			Get(ctx, nil)
		if err != nil {
			return err
		}
	}

	return nil
}
