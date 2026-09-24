package feed_tasks

import (
	"time"

	"go.temporal.io/sdk/workflow"
)

const (
	LegacyGenerateFeedsWorkflowName = "GenerateFeedsWorkflow"
	GenerateFeedsWorkflowName       = "GenerateFeedsWorkflowV2"
	RefreshFeedsWorkflowName        = "RefreshFeedsWorkflow"
	LegacyAddFeedWorkflowName       = "AddFeedWorkflow"
	AddFeedWorkflowName             = "AddFeedWorkflowV2"
)

type GenerateFeedsInput struct {
	Host  string
	Count int
}

type AddFeedInput struct {
	URL    string
	UserID string
}

func GenerateFeedsWorkflow(ctx workflow.Context, host string, count int) error {
	return generateFeedsWorkflow(ctx, GenerateFeedsInput{Host: host, Count: count})
}

func GenerateFeedsWorkflowV2(ctx workflow.Context, input GenerateFeedsInput) error {
	return generateFeedsWorkflow(ctx, input)
}

func generateFeedsWorkflow(ctx workflow.Context, input GenerateFeedsInput) error {
	// This workflow intentionally enqueues one add-feed job per generated URL;
	// it does not wait for those independent add workflows to complete.
	ctx = workflow.WithActivityOptions(ctx, workflow.ActivityOptions{
		StartToCloseTimeout: time.Minute,
	})

	var a *Activities

	err := workflow.ExecuteActivity(ctx, a.GenerateFeeds, input.Host, input.Count /*, userID*/).Get(ctx, nil)
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
	return addFeedWorkflow(ctx, AddFeedInput{URL: url, UserID: userID})
}

func AddFeedWorkflowV2(ctx workflow.Context, input AddFeedInput) (string, error) {
	return addFeedWorkflow(ctx, input)
}

func addFeedWorkflow(ctx workflow.Context, input AddFeedInput) (string, error) {
	// Compatibility wrapper for executions started before the application
	// dispatcher began starting feed_add.AddFeedWorkflow directly.
	ctx = workflow.WithActivityOptions(ctx, workflow.ActivityOptions{
		StartToCloseTimeout: time.Minute,
	})
	var a *Activities
	var feedID string
	err := workflow.ExecuteActivity(ctx, a.AddFeed, input.URL, input.UserID).Get(ctx, &feedID)
	return feedID, err
}
