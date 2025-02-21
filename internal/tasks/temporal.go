package tasks

import (
	"context"
	"errors"

	"github.com/ericbutera/amalgam/data-pipeline/temporal/feed_tasks"
	"github.com/ericbutera/amalgam/pkg/config/env"
	sdk "go.temporal.io/sdk/client"
	"go.temporal.io/sdk/temporal"
)

var ErrHostRequired = errors.New("temporal host is required")

type Temporal struct {
	config feed_tasks.Config
	client sdk.Client
}

func NewTemporal(config *feed_tasks.Config, client *sdk.Client) (*Temporal, error) {
	return &Temporal{
		config: *config,
		client: *client,
	}, nil
}

func (t *Temporal) Close() {
	t.client.Close()
}

func NewTemporalFromEnv() (*Temporal, error) {
	config, err := env.New[feed_tasks.Config]()
	if err != nil {
		return nil, err
	}
	if config.TemporalHost == "" {
		return nil, ErrHostRequired
	}
	client, err := sdk.Dial(sdk.Options{
		HostPort: config.TemporalHost,
	})
	if err != nil {
		return nil, err
	}
	return NewTemporal(config, &client)
}

func (t *Temporal) Workflow(ctx context.Context, task TaskType, args []any) (*TaskResult, error) {
	workflow, err := taskTypeToWorkflow(task)
	if err != nil {
		return nil, err
	}

	if task == TaskGenerateFeeds {
		args = []any{
			t.config.FakeHost,
			t.config.GenerateCount,
		}
	}

	opts := sdk.StartWorkflowOptions{
		TaskQueue: t.config.TaskQueue,
		RetryPolicy: &temporal.RetryPolicy{
			MaximumAttempts: 1, // TODO: configurable
		},
	}
	we, err := t.client.ExecuteWorkflow(
		ctx,
		opts,
		workflow,
		args...,
	)
	if err != nil {
		return nil, err
	}
	return &TaskResult{
		ID:    we.GetID(),
		RunID: we.GetRunID(),
	}, nil
}

// func (t *Temporal) JobStatus(ctx context.Context, workflowID string, runID string) (*TaskResult, error) {
// 	run := t.client.GetWorkflow(ctx, workflowID, runID)
// 	var feedID string
// 	err := run.Get(ctx, feedID)
// 	if err != nil {
// 		return nil, err
// 	}
// }
