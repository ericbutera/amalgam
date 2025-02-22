package tasks

import (
	"context"
	"errors"

	"github.com/ericbutera/amalgam/data-pipeline/temporal/feed_tasks"
	"github.com/ericbutera/amalgam/pkg/config/env"
	"go.temporal.io/api/enums/v1"
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

func (t *Temporal) Close() {
	t.client.Close()
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
	var result any
	_ = we.Get(ctx, result)
	return &TaskResult{
		ID:     we.GetID(),
		RunID:  we.GetRunID(),
		Result: result,
	}, nil
}

func (t *Temporal) Status(ctx context.Context, taskID string) (*TaskStatusResult, error) {
	data := ""
	history := t.client.GetWorkflowHistory(ctx, taskID, "", false, enums.HISTORY_EVENT_FILTER_TYPE_CLOSE_EVENT)
	for {
		if !history.HasNext() {
			break
		}
		event, err := history.Next()
		if err != nil {
			break
		}
		data = event.EventType.String()
	}

	return &TaskStatusResult{
		Status: data,
	}, nil
}
