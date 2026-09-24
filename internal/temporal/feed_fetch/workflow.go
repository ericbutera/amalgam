package feed_fetch

import (
	"errors"
	"fmt"
	"time"

	"github.com/ericbutera/amalgam/internal/temporal/feeds"
	"go.temporal.io/api/enums/v1"
	"go.temporal.io/sdk/temporal"
	"go.temporal.io/sdk/workflow"
)

const (
	MaxConcurrentFeeds     int = 10
	Timeout                    = 5 * time.Minute
	WorkflowName               = "FetchFeedsWorkflow"
	LegacyFeedWorkflowName     = "FeedWorkflow"
	FeedWorkflowName           = "FeedWorkflowV2"
)

var (
	ErrTimeout  = errors.New("workflow timeout")
	ErrProcess  = errors.New("error processing feeds")
	retryPolicy = temporal.RetryPolicy{
		MaximumAttempts: 1,
	}
)

type FeedFailure struct {
	FeedID    string
	Stage     string
	ErrorType string
	Retryable bool
}

type FeedInput struct {
	FeedID string
	URL    string
}

type FetchFeedsResult struct {
	Total     int
	Succeeded int
	Failed    int
	Failures  []FeedFailure
}

func applicationErrorInfo(err error) (string, bool) {
	var applicationErr *temporal.ApplicationError
	if errors.As(err, &applicationErr) {
		if applicationErr.Type() != "" {
			return applicationErr.Type(), !applicationErr.NonRetryable()
		}
	}

	return "WorkflowFailure", true
}

func applicationErrorType(err error) string {
	errorType, _ := applicationErrorInfo(err)
	return errorType
}

// FeedWorkflow preserves the original positional-input child workflow contract
// for executions that were started before FeedInput was introduced.
func FeedWorkflow(ctx workflow.Context, feedID string, url string) error {
	return feedWorkflow(ctx, FeedInput{FeedID: feedID, URL: url})
}

func FeedWorkflowV2(ctx workflow.Context, input FeedInput) error {
	return feedWorkflow(ctx, input)
}

func feedWorkflow(ctx workflow.Context, input FeedInput) error {
	ctx = workflow.WithActivityOptions(ctx, workflow.ActivityOptions{
		StartToCloseTimeout: time.Minute,
		RetryPolicy:         &retryPolicy,
	})

	var a *Activities

	var rssFile string

	err := workflow.ExecuteActivity(ctx, a.DownloadActivity, input.FeedID, input.URL).Get(ctx, &rssFile)
	if err != nil {
		if applicationErrorType(err) == ContentNotChangedErrorType {
			workflow.GetLogger(ctx).Debug("content not changed", "feed_id", input.FeedID)
			return nil
		}

		return err
	}

	var articlesFile string

	err = workflow.ExecuteActivity(ctx, a.ParseActivity, input.FeedID, rssFile).Get(ctx, &articlesFile)
	if err != nil {
		return err
	}

	var saveResults SaveResults
	err = workflow.ExecuteActivity(ctx, a.SaveActivity, input.FeedID, articlesFile).Get(ctx, &saveResults)
	if err != nil {
		return err
	}
	if saveResults.Failed > 0 {
		return temporal.NewNonRetryableApplicationError(
			"some articles could not be saved",
			PartialSaveErrorType,
			nil,
			saveResults,
		)
	}

	err = workflow.ExecuteActivity(ctx, a.StatsActivity, input.FeedID).Get(ctx, nil)
	if err != nil {
		return err
	}

	return nil
}

func FetchFeedsWorkflow(ctx workflow.Context) (*FetchFeedsResult, error) {
	ctx = workflow.WithActivityOptions(ctx, workflow.ActivityOptions{
		StartToCloseTimeout: Timeout,
		RetryPolicy:         &retryPolicy,
	})
	// The five-minute timeout is the current batch SLA. Feed pagination and
	// history growth remain separate follow-up work.

	var (
		a    *Activities
		urls []feeds.Feed
	)

	if err := workflow.ExecuteActivity(ctx, a.GetFeedsActivity).Get(ctx, &urls); err != nil {
		return nil, err
	}

	semaphore := workflow.NewSemaphore(ctx, int64(MaxConcurrentFeeds))
	result := &FetchFeedsResult{Total: len(urls), Failures: make([]FeedFailure, 0)}
	done := 0
	childVersion := workflow.GetVersion(ctx, "feed-child-input", workflow.DefaultVersion, 1)
	info := workflow.GetInfo(ctx)
	childContext := workflow.WithChildOptions(ctx, workflow.ChildWorkflowOptions{
		TaskQueue:                info.TaskQueueName,
		ParentClosePolicy:        enums.PARENT_CLOSE_POLICY_REQUEST_CANCEL,
		WorkflowIDReusePolicy:    enums.WORKFLOW_ID_REUSE_POLICY_REJECT_DUPLICATE,
		WorkflowExecutionTimeout: Timeout,
	})

	for i, feed := range urls {
		index := i
		feed := feed
		err := semaphore.Acquire(ctx, 1)
		if err != nil {
			return nil, err
		}

		workflow.Go(ctx, func(ctx workflow.Context) {
			defer semaphore.Release(1)

			childOptions := workflow.GetChildWorkflowOptions(childContext)
			childOptions.WorkflowID = fmt.Sprintf("%s:feed:%d:%s", info.WorkflowExecution.RunID, index, feed.ID)
			childCtx := workflow.WithChildOptions(ctx, childOptions)
			var future workflow.ChildWorkflowFuture
			if childVersion == workflow.DefaultVersion {
				future = workflow.ExecuteChildWorkflow(childCtx, FeedWorkflow, feed.ID, feed.Url)
			} else {
				future = workflow.ExecuteChildWorkflow(childCtx, FeedWorkflowV2, FeedInput{FeedID: feed.ID, URL: feed.Url})
			}

			err := future.Get(ctx, nil)
			if err != nil {
				workflow.GetLogger(ctx).Error("failed to process feed", "i", index, "feed_id", feed.ID, "error", err)
				errorType, retryable := applicationErrorInfo(err)
				result.Failed++
				result.Failures = append(result.Failures, FeedFailure{
					FeedID:    feed.ID,
					Stage:     "feed",
					ErrorType: errorType,
					Retryable: retryable,
				})
			} else {
				workflow.GetLogger(ctx).Info("processed feed", "i", index, "feed_id", feed.ID)
				result.Succeeded++
			}

			done++
		})
	}

	ok, err := workflow.AwaitWithTimeout(ctx, Timeout, func() bool {
		return len(urls) == done
	})
	if err != nil {
		return nil, err
	}

	if !ok {
		return nil, ErrTimeout
	}

	return result, nil
}
