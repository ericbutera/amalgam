package tasks

import (
	"context"
	"errors"

	"github.com/ericbutera/amalgam/data-pipeline/temporal/feed_tasks"
)

var ErrInvalidTaskType = errors.New("invalid task type")

type TaskType string

const (
	TaskUnspecified   TaskType = ""
	TaskGenerateFeeds TaskType = "generate_feeds"
	TaskFetchFeeds    TaskType = "fetch_feeds"
	TaskAddFeed       TaskType = "add_feed"
)

type Tasks interface {
	// Args will be passed as parameters to the workflow.
	Workflow(ctx context.Context, task TaskType, args []any) (*TaskResult, error)
}

type TaskResult struct {
	ID string
}

func taskTypeToWorkflow(taskType TaskType) (any, error) {
	switch taskType { //nolint:exhaustive
	case TaskGenerateFeeds:
		return feed_tasks.GenerateFeedsWorkflow, nil
	case TaskFetchFeeds:
		return feed_tasks.RefreshFeedsWorkflow, nil
	case TaskAddFeed:
		return feed_tasks.AddFeedWorkflow, nil
	}
	return nil, ErrInvalidTaskType
}
