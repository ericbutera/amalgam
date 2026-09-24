package tasks

import (
	"context"
	"errors"

	"github.com/ericbutera/amalgam/internal/temporal/feed_fetch"
)

var ErrInvalidTaskType = errors.New("invalid task type")

type TaskType string

const (
	TaskUnspecified   TaskType = ""
	TaskGenerateFeeds TaskType = "generate_feeds"
	TaskFetchFeeds    TaskType = "fetch_feeds"
	TaskAddFeed       TaskType = "add_feed"
)

// WorkflowRequest is the application-facing enqueue contract. It deliberately
// contains business inputs instead of exposing Temporal's variadic argument
// list to callers.
type WorkflowRequest struct {
	Task       TaskType
	URL        string
	Name       string
	UserID     string
	WorkflowID string
}

type Tasks interface {
	// Enqueue starts durable background work and returns its Temporal handle.
	Enqueue(ctx context.Context, request WorkflowRequest) (*TaskResult, error)
	Status(ctx context.Context, taskID string) (*TaskStatusResult, error)
	GetResult(ctx context.Context, taskID string) (*TaskWorkflowResult, error)
}

type TaskResult struct {
	ID           string
	RunID        string
	WorkflowType string
}

type TaskStatusResult struct {
	ID     string
	RunID  string
	Status string
}

// TaskWorkflowResult is the internal typed union for workflow payloads. The
// workflow type determines which payload field is populated.
type TaskWorkflowResult struct {
	ID           string
	RunID        string
	WorkflowType string
	AddFeedID    string
	FetchFeeds   *feed_fetch.FetchFeedsResult
}
