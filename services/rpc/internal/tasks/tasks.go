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
)

type Tasks interface {
	Workflow(ctx context.Context, task TaskType) (*TaskResult, error)
}

type TaskResult struct {
	ID string
}

func taskTypeToWorkflow(taskType TaskType) (any, error) {
	if taskType == TaskGenerateFeeds {
		return feed_tasks.GenerateFeedsWorkflow, nil
	}
	return nil, ErrInvalidTaskType
}
