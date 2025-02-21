package feed_add

import (
	"errors"
	"time"

	"go.temporal.io/sdk/temporal"
	"go.temporal.io/sdk/workflow"
)

var (
	Timeout     = 5 * time.Minute
	retryPolicy = temporal.RetryPolicy{
		MaximumAttempts: 1,
	}

	ErrDuplicateFeed = errors.New("duplicate feed")
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

	var err error
	var a *Activities
	var feedID string

	workflowID := workflow.GetInfo(ctx).WorkflowExecution.ID

	err = workflow.ExecuteActivity(ctx, a.SubscribeUserToUrl, url, userID).Get(ctx, &feedID)
	if err != nil {
		return err
	}
	if feedID != "" {
		return nil // feed exists and user is associated, exit!
	}

	var verification FeedVerification
	err = workflow.ExecuteActivity(ctx, a.CreateVerifyRecord, FeedVerification{
		URL:        url,
		UserID:     userID,
		WorkflowID: workflowID,
	}).Get(ctx, &verification)
	if err != nil {
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
