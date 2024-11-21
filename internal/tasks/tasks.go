package tasks

import (
	"context"
	"errors"

	feeds "github.com/ericbutera/amalgam/data-pipeline/temporal/feed"
	"github.com/ericbutera/amalgam/data-pipeline/temporal/feed_tasks"
)

var ErrInvalidTaskType = errors.New("invalid task type")

type TaskType string

const (
	TaskUnspecified   TaskType = ""
	TaskGenerateFeeds TaskType = "generate_feeds"
	TaskFetchFeeds    TaskType = "fetch_feeds"
)

type Tasks interface {
	Workflow(ctx context.Context, task TaskType) (*TaskResult, error)
}

type TaskResult struct {
	ID string
}

func taskTypeToWorkflow(taskType TaskType) (any, error) {
	switch taskType { //nolint:exhaustive
	case TaskGenerateFeeds:
		return feed_tasks.GenerateFeedsWorkflow, nil
	case TaskFetchFeeds:
		return feeds.FetchFeedsWorkflow, nil
	}
	return nil, ErrInvalidTaskType
}
