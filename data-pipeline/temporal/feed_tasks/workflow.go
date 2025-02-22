package feed_tasks

import (
	"time"

	"go.temporal.io/sdk/workflow"
)

func GenerateFeedsWorkflow(ctx workflow.Context, host string, count int /*, userID string*/) error {
	ctx = workflow.WithActivityOptions(ctx, workflow.ActivityOptions{
		StartToCloseTimeout: time.Minute,
	})

	var a *Activities

	err := workflow.ExecuteActivity(ctx, a.GenerateFeeds, host, count /*, userID*/).Get(ctx, nil)
	if err != nil {
		return err
	}

	return nil
}

func RefreshFeedsWorkflow(ctx workflow.Context) error {
	ctx = workflow.WithActivityOptions(ctx, workflow.ActivityOptions{
		StartToCloseTimeout: time.Minute,
	})

	var a *Activities

	err := workflow.ExecuteActivity(ctx, a.RefreshFeeds).Get(ctx, nil)
	if err != nil {
		return err
	}

	return nil
}

func AddFeedWorkflow(ctx workflow.Context, url string, userID string) (string, error) {
	ctx = workflow.WithActivityOptions(ctx, workflow.ActivityOptions{
		StartToCloseTimeout: time.Minute,
	})
	var a *Activities
	var feedID string
	err := workflow.ExecuteActivity(ctx, a.AddFeed, url, userID).Get(ctx, &feedID)
	return feedID, err
}
