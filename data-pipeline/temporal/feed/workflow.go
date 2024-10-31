package app

import (
	"time"

	"github.com/ericbutera/amalgam/data-pipeline/temporal/feed/internal/config"
	"github.com/ericbutera/amalgam/data-pipeline/temporal/feed/internal/feeds"
	cfg "github.com/ericbutera/amalgam/pkg/config"
	"github.com/samber/lo"
	"go.temporal.io/sdk/workflow"
)

func FeedWorkflow(ctx workflow.Context) error {
	ctx = workflow.WithActivityOptions(ctx, workflow.ActivityOptions{
		StartToCloseTimeout: time.Minute,
	})

	var a *Activities

	// TODO: split into parent child workflows
	// - move fetch feeds into an activity, save historical feed list to bucket
	// - start child "fetch feed" workflows for each feed
	config := lo.Must(cfg.NewFromEnv[config.Config]())
	feeds := lo.Must(feeds.NewFeeds(config))
	defer feeds.Close()
	urls := lo.Must(feeds.GetFeeds())

	for _, feed := range urls {
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
