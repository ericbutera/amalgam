package tasks

import (
	"context"
	"testing"

	"github.com/ericbutera/amalgam/internal/temporal/feed_add"
	"github.com/ericbutera/amalgam/internal/temporal/feed_fetch"
	"github.com/ericbutera/amalgam/internal/temporal/feed_tasks"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
	commonpb "go.temporal.io/api/common/v1"
	"go.temporal.io/api/enums/v1"
	workflowpb "go.temporal.io/api/workflow/v1"
	workflowservice "go.temporal.io/api/workflowservice/v1"
	sdk "go.temporal.io/sdk/client"
	sdkmocks "go.temporal.io/sdk/mocks"
)

func TestTemporalEnqueueStartsOwningAddWorkflow(t *testing.T) {
	client := &sdkmocks.Client{}
	run := &sdkmocks.WorkflowRun{}
	run.On("GetID").Return("add-workflow")
	run.On("GetRunID").Return("add-run")

	client.On(
		"ExecuteWorkflow",
		mock.Anything,
		mock.MatchedBy(func(options sdk.StartWorkflowOptions) bool {
			return options.TaskQueue == "feed-add" && options.ID == ""
		}),
		mock.AnythingOfType("func(internal.Context, feed_add.AddFeedInput) (string, error)"),
		feed_add.AddFeedInput{URL: "https://example.com/feed.xml", Name: "Example feed", UserID: "user-1"},
	).Return(run, nil)

	config := &feed_tasks.Config{FeedAddQueue: "feed-add"}
	var sdkClient sdk.Client = client
	temporal, err := NewTemporal(config, &sdkClient)
	require.NoError(t, err)

	result, err := temporal.Enqueue(context.Background(), WorkflowRequest{
		Task:   TaskAddFeed,
		URL:    "https://example.com/feed.xml",
		Name:   "Example feed",
		UserID: "user-1",
	})
	require.NoError(t, err)
	require.Equal(t, &TaskResult{
		ID:           "add-workflow",
		RunID:        "add-run",
		WorkflowType: feed_add.WorkflowName,
	}, result)
	client.AssertExpectations(t)
	run.AssertExpectations(t)
}

func TestTemporalStatusUsesDescribeAndStableStatus(t *testing.T) {
	client := &sdkmocks.Client{}
	client.On("DescribeWorkflowExecution", mock.Anything, "workflow-1", "").Return(
		&workflowservice.DescribeWorkflowExecutionResponse{
			WorkflowExecutionInfo: &workflowpb.WorkflowExecutionInfo{
				Execution: &commonpb.WorkflowExecution{WorkflowId: "workflow-1", RunId: "run-1"},
				Status:    enums.WORKFLOW_EXECUTION_STATUS_RUNNING,
			},
		},
		nil,
	)

	config := &feed_tasks.Config{}
	var sdkClient sdk.Client = client
	temporal, err := NewTemporal(config, &sdkClient)
	require.NoError(t, err)

	result, err := temporal.Status(context.Background(), "workflow-1")
	require.NoError(t, err)
	require.Equal(t, &TaskStatusResult{
		ID:     "workflow-1",
		RunID:  "run-1",
		Status: "RUNNING",
	}, result)
	client.AssertExpectations(t)
}

func TestTemporalGetResultDecodesFetchResult(t *testing.T) {
	client := &sdkmocks.Client{}
	run := &sdkmocks.WorkflowRun{}
	run.On("Get", mock.Anything, mock.AnythingOfType("*feed_fetch.FetchFeedsResult")).Run(func(args mock.Arguments) {
		result := args.Get(1).(*feed_fetch.FetchFeedsResult)
		*result = feed_fetch.FetchFeedsResult{Total: 2, Succeeded: 1, Failed: 1}
	}).Return(nil)

	client.On("DescribeWorkflowExecution", mock.Anything, "fetch-workflow", "").Return(
		&workflowservice.DescribeWorkflowExecutionResponse{
			WorkflowExecutionInfo: &workflowpb.WorkflowExecutionInfo{
				Execution: &commonpb.WorkflowExecution{WorkflowId: "fetch-workflow", RunId: "fetch-run"},
				Type:      &commonpb.WorkflowType{Name: feed_fetch.WorkflowName},
				Status:    enums.WORKFLOW_EXECUTION_STATUS_COMPLETED,
			},
		},
		nil,
	)
	client.On("GetWorkflow", mock.Anything, "fetch-workflow", "fetch-run").Return(run)

	config := &feed_tasks.Config{}
	var sdkClient sdk.Client = client
	temporal, err := NewTemporal(config, &sdkClient)
	require.NoError(t, err)

	result, err := temporal.GetResult(context.Background(), "fetch-workflow")
	require.NoError(t, err)
	require.Equal(t, &TaskWorkflowResult{
		ID:           "fetch-workflow",
		RunID:        "fetch-run",
		WorkflowType: feed_fetch.WorkflowName,
		FetchFeeds:   &feed_fetch.FetchFeedsResult{Total: 2, Succeeded: 1, Failed: 1},
	}, result)
	client.AssertExpectations(t)
	run.AssertExpectations(t)
}

func TestWorkflowStatus(t *testing.T) {
	tests := []struct {
		status enums.WorkflowExecutionStatus
		want   string
	}{
		{enums.WORKFLOW_EXECUTION_STATUS_RUNNING, "RUNNING"},
		{enums.WORKFLOW_EXECUTION_STATUS_COMPLETED, "COMPLETED"},
		{enums.WORKFLOW_EXECUTION_STATUS_FAILED, "FAILED"},
		{enums.WORKFLOW_EXECUTION_STATUS_CANCELED, "CANCELED"},
		{enums.WORKFLOW_EXECUTION_STATUS_TIMED_OUT, "TIMED_OUT"},
		{enums.WORKFLOW_EXECUTION_STATUS_TERMINATED, "TERMINATED"},
		{enums.WORKFLOW_EXECUTION_STATUS_UNSPECIFIED, "UNKNOWN"},
	}

	for _, test := range tests {
		t.Run(test.want, func(t *testing.T) {
			require.Equal(t, test.want, workflowStatus(test.status))
		})
	}
}
