package feed_add

import (
	"time"

	"go.temporal.io/sdk/temporal"
	"go.temporal.io/sdk/workflow"
)

var (
	Timeout     = 5 * time.Minute
	retryPolicy = temporal.RetryPolicy{
		MaximumAttempts: 1,
	}
)

type FeedVerification struct {
	ID         int64
	WorkflowID string
	URL        string
	UserID     string
}

func AddFeedWorkflow(ctx workflow.Context, url string, userID string) error {
	ctx = workflow.WithActivityOptions(ctx, workflow.ActivityOptions{
		StartToCloseTimeout: Timeout,
		RetryPolicy:         &retryPolicy,
	})

	var a *Activities

	workflowID := workflow.GetInfo(ctx).WorkflowExecution.ID

	var verification FeedVerification
	err := workflow.ExecuteActivity(ctx, a.CreateVerifyRecord, FeedVerification{
		URL:        url,
		UserID:     userID,
		WorkflowID: workflowID,
	}).Get(ctx, &verification)
	if err != nil {
		// if feed already exists do not retry
		// if feed response is stop, do not retry
		// https://rachelbythebay.com/w/2024/05/27/feed/
		return err
	}

	var blob string // TODO: write to bucket and pass reference
	err = workflow.ExecuteActivity(ctx, a.Fetch, verification).Get(ctx, &blob)
	if err != nil {
		return err
	}

	err = workflow.ExecuteActivity(ctx, a.CreateFeed, verification).Get(ctx, nil)
	if err != nil {
		return err
	}

	return nil
}
