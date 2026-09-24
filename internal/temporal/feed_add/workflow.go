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

const (
	LegacyWorkflowName = "AddFeedWorkflow"
	WorkflowName       = "AddFeedWorkflowV2"
)

func isApplicationErrorType(err error, errorType string) bool {
	var applicationErr *temporal.ApplicationError
	return errors.As(err, &applicationErr) && applicationErr.Type() == errorType
}

func subscribeUserToFeed(ctx workflow.Context, a *Activities, url string, userID string) (string, error) {
	var feedID string
	if err := workflow.ExecuteActivity(ctx, a.SubscribeUserToUrl, url, userID).Get(ctx, &feedID); err != nil {
		return feedID, err
	}
	if feedID == "" {
		return feedID, temporal.NewNonRetryableApplicationError("subscription returned no feed ID", "MissingFeedID", nil)
	}
	return feedID, nil
}

type FeedVerification struct {
	ID         int64
	WorkflowID string
	URL        string
	Name       string
	UserID     string
}

type AddFeedInput struct {
	URL    string
	Name   string
	UserID string
}

// AddFeedWorkflow preserves the original positional-input workflow contract
// for executions that were started before AddFeedInput was introduced.
func AddFeedWorkflow(ctx workflow.Context, url string, userID string) (string, error) {
	return addFeedWorkflow(ctx, AddFeedInput{URL: url, UserID: userID})
}

func AddFeedWorkflowV2(ctx workflow.Context, input AddFeedInput) (string, error) {
	return addFeedWorkflow(ctx, input)
}

func addFeedWorkflow(ctx workflow.Context, input AddFeedInput) (string, error) {
	// 1. determine if feed exists
	//    - if it does, subscribe user to feed and return feed ID
	// 2. create verification record
	// 3. fetch & validate feed
	// 4. create feed record
	// 5. subscribe user to feed

	ctx = workflow.WithActivityOptions(ctx, workflow.ActivityOptions{
		StartToCloseTimeout: Timeout,
		RetryPolicy:         &retryPolicy,
	})

	var err error
	var a *Activities
	var feedID string

	workflowID := workflow.GetInfo(ctx).WorkflowExecution.ID

	// 1. determine if feed exists by attempting to subscribe the user.
	//    If the feed exists, the RPC will succeed and return the feed ID.
	err = workflow.ExecuteActivity(ctx, a.SubscribeUserToUrl, input.URL, input.UserID).Get(ctx, &feedID)
	if err == nil && feedID != "" {
		return feedID, nil
	}
	if err != nil && !isApplicationErrorType(err, FeedNotFoundErrorType) {
		return feedID, err
	}

	var verification FeedVerification
	err = workflow.ExecuteActivity(ctx, a.CreateVerifyRecord, FeedVerification{
		URL:        input.URL,
		Name:       input.Name,
		UserID:     input.UserID,
		WorkflowID: workflowID,
	}).Get(ctx, &verification)
	if err != nil {
		if isApplicationErrorType(err, DuplicateFeedErrorType) {
			return subscribeUserToFeed(ctx, a, input.URL, input.UserID)
		}
		return feedID, err
	}

	var blob string // TODO: write to bucket and pass reference

	err = workflow.ExecuteActivity(ctx, a.Fetch, verification).Get(ctx, &blob)
	if err != nil {
		return feedID, err
	}

	err = workflow.ExecuteActivity(ctx, a.CreateFeed, verification).Get(ctx, &feedID)
	if err != nil {
		if isApplicationErrorType(err, DuplicateFeedErrorType) {
			return subscribeUserToFeed(ctx, a, input.URL, input.UserID)
		}
		return feedID, err
	}

	return subscribeUserToFeed(ctx, a, input.URL, input.UserID)
}
