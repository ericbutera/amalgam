package feed_tasks

import (
	"time"

	"go.temporal.io/sdk/workflow"
)

func GenerateFeedsWorkflow(ctx workflow.Context, host string, count int) error {
	ctx = workflow.WithActivityOptions(ctx, workflow.ActivityOptions{
		StartToCloseTimeout: time.Minute,
	})

	var a *Activities

	err := workflow.ExecuteActivity(ctx, a.GenerateFeeds, host, count).Get(ctx, nil)
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
