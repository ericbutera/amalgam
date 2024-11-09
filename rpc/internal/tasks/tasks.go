package tasks

import (
	"context"
	"errors"

	"github.com/ericbutera/amalgam/data-pipeline/temporal/feed_tasks"
	"github.com/ericbutera/amalgam/pkg/config/env"
	"github.com/samber/lo"
	sdk "go.temporal.io/sdk/client"
	"go.temporal.io/sdk/temporal"
)

type TaskType string

const (
	TaskUnspecified   TaskType = ""
	TaskGenerateFeeds TaskType = "generate_feeds"
)

type TaskResult struct {
	ID string
}

func New(ctx context.Context, task TaskType) (*TaskResult, error) {
	workflow, err := taskTypeToWorkflow(task)
	if err != nil {
		return nil, err
	}

	// TODO: Dependency injection
	// TODO: handle connection (defer & reconnect)
	config := lo.Must(env.New[feed_tasks.Config]())
	client := lo.Must(sdk.Dial(sdk.Options{
		HostPort: config.TemporalHost,
	}))

	defer client.Close()

	opts := sdk.StartWorkflowOptions{
		TaskQueue: config.TaskQueue,
		RetryPolicy: &temporal.RetryPolicy{
			MaximumAttempts: 1,
		},
	}
	we, err := client.ExecuteWorkflow(ctx, opts, workflow, config.FakeHost, config.GenerateCount)
	if err != nil {
		return nil, err
	}
	return &TaskResult{ID: we.GetID()}, nil
}

func taskTypeToWorkflow(taskType TaskType) (any, error) {
	switch taskType {
	case TaskGenerateFeeds:
		return feed_tasks.GenerateFeedsWorkflow, nil
	default:
		return nil, errors.New("unknown task type")
	}
}
