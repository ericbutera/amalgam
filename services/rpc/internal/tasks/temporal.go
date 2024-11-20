package tasks

import (
	"context"

	"github.com/ericbutera/amalgam/data-pipeline/temporal/feed_tasks"
	"github.com/ericbutera/amalgam/pkg/config/env"
	sdk "go.temporal.io/sdk/client"
	"go.temporal.io/sdk/temporal"
)

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

func NewTemporalWithDefaults() (*Temporal, error) {
	config, err := env.New[feed_tasks.Config]()
	if err != nil {
		return nil, err
	}
	client, err := sdk.Dial(sdk.Options{
		HostPort: config.TemporalHost,
	})
	if err != nil {
		return nil, err
	}
	return NewTemporal(config, &client)
}

func (t *Temporal) Workflow(ctx context.Context, task TaskType) (*TaskResult, error) {
	workflow, err := taskTypeToWorkflow(task)
	if err != nil {
		return nil, err
	}
	opts := sdk.StartWorkflowOptions{
		TaskQueue: t.config.TaskQueue,
		RetryPolicy: &temporal.RetryPolicy{
			MaximumAttempts: 1,
		},
	}
	we, err := t.client.ExecuteWorkflow(
		ctx,
		opts,
		workflow,
		t.config.FakeHost,
		t.config.GenerateCount,
	)
	if err != nil {
		return nil, err
	}
	return &TaskResult{ID: we.GetID()}, nil
}
