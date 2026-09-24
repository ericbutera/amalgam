package tasks

import (
	"context"
	"errors"

	"github.com/ericbutera/amalgam/internal/temporal/feed_add"
	"github.com/ericbutera/amalgam/internal/temporal/feed_fetch"
	"github.com/ericbutera/amalgam/internal/temporal/feed_tasks"
	"github.com/ericbutera/amalgam/pkg/config/env"
	"go.temporal.io/api/enums/v1"
	sdk "go.temporal.io/sdk/client"
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

func (t *Temporal) Enqueue(ctx context.Context, request WorkflowRequest) (*TaskResult, error) {
	workflow, taskQueue, input, workflowType, err := t.workflowSpec(request)
	if err != nil {
		return nil, err
	}

	opts := sdk.StartWorkflowOptions{
		TaskQueue: taskQueue,
	}
	if request.WorkflowID != "" {
		opts.ID = request.WorkflowID
	}

	// ExecuteWorkflow is the enqueue operation. Do not call Get here: callers
	// receive the workflow handle and can choose when to retrieve the result.
	var we sdk.WorkflowRun
	if input == nil {
		we, err = t.client.ExecuteWorkflow(ctx, opts, workflow)
	} else {
		we, err = t.client.ExecuteWorkflow(ctx, opts, workflow, input)
	}
	if err != nil {
		return nil, err
	}

	return &TaskResult{
		ID:           we.GetID(),
		RunID:        we.GetRunID(),
		WorkflowType: workflowType,
	}, nil
}

func (t *Temporal) Status(ctx context.Context, taskID string) (*TaskStatusResult, error) {
	description, err := t.client.DescribeWorkflowExecution(ctx, taskID, "")
	if err != nil {
		return nil, err
	}

	info := description.GetWorkflowExecutionInfo()
	if info == nil {
		return nil, errors.New("workflow execution description is empty")
	}

	return &TaskStatusResult{
		ID:     taskID,
		RunID:  info.GetExecution().GetRunId(),
		Status: workflowStatus(info.GetStatus()),
	}, nil
}

func (t *Temporal) GetResult(ctx context.Context, taskID string) (*TaskWorkflowResult, error) {
	description, err := t.client.DescribeWorkflowExecution(ctx, taskID, "")
	if err != nil {
		return nil, err
	}

	info := description.GetWorkflowExecutionInfo()
	if info == nil || info.GetExecution() == nil {
		return nil, errors.New("workflow execution description is empty")
	}

	workflowType := info.GetType().GetName()
	result := &TaskWorkflowResult{
		ID:           taskID,
		RunID:        info.GetExecution().GetRunId(),
		WorkflowType: workflowType,
	}
	run := t.client.GetWorkflow(ctx, taskID, result.RunID)

	switch workflowType {
	case feed_add.WorkflowName, feed_add.LegacyWorkflowName:
		if err := run.Get(ctx, &result.AddFeedID); err != nil {
			return nil, err
		}
	case feed_fetch.WorkflowName:
		result.FetchFeeds = &feed_fetch.FetchFeedsResult{}
		if err := run.Get(ctx, result.FetchFeeds); err != nil {
			return nil, err
		}
	case feed_tasks.GenerateFeedsWorkflowName, feed_tasks.LegacyGenerateFeedsWorkflowName:
		if err := run.Get(ctx, nil); err != nil {
			return nil, err
		}
	default:
		return nil, errors.New("unsupported workflow result type")
	}

	return result, nil
}

func workflowStatus(status enums.WorkflowExecutionStatus) string {
	switch status {
	case enums.WORKFLOW_EXECUTION_STATUS_RUNNING:
		return "RUNNING"
	case enums.WORKFLOW_EXECUTION_STATUS_COMPLETED:
		return "COMPLETED"
	case enums.WORKFLOW_EXECUTION_STATUS_FAILED:
		return "FAILED"
	case enums.WORKFLOW_EXECUTION_STATUS_CANCELED:
		return "CANCELED"
	case enums.WORKFLOW_EXECUTION_STATUS_TERMINATED:
		return "TERMINATED"
	case enums.WORKFLOW_EXECUTION_STATUS_CONTINUED_AS_NEW:
		return "CONTINUED_AS_NEW"
	case enums.WORKFLOW_EXECUTION_STATUS_TIMED_OUT:
		return "TIMED_OUT"
	default:
		return "UNKNOWN"
	}
}

func (t *Temporal) workflowSpec(request WorkflowRequest) (any, string, any, string, error) {
	switch request.Task { //nolint:exhaustive
	case TaskGenerateFeeds:
		return feed_tasks.GenerateFeedsWorkflowV2,
			t.config.TaskQueue,
			feed_tasks.GenerateFeedsInput{Host: t.config.FakeHost, Count: t.config.GenerateCount},
			feed_tasks.GenerateFeedsWorkflowName,
			nil
	case TaskFetchFeeds:
		return feed_fetch.FetchFeedsWorkflow,
			t.config.FeedFetchQueue,
			nil,
			feed_fetch.WorkflowName,
			nil
	case TaskAddFeed:
		if request.URL == "" || request.UserID == "" {
			return nil, "", nil, "", errors.New("add feed requires URL and user ID")
		}

		return feed_add.AddFeedWorkflowV2,
			t.config.FeedAddQueue,
			feed_add.AddFeedInput{URL: request.URL, Name: request.Name, UserID: request.UserID},
			feed_add.WorkflowName,
			nil
	default:
		return nil, "", nil, "", ErrInvalidTaskType
	}
}
