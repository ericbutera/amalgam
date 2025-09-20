package feed_tasks_test

import (
	"testing"

	"github.com/ericbutera/amalgam/data-pipeline/temporal/feed_tasks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
	"go.temporal.io/sdk/testsuite"
	"go.temporal.io/sdk/worker"
)

type UnitTestSuite struct {
	suite.Suite
	testsuite.WorkflowTestSuite
}

func TestUnitTestSuite(t *testing.T) {
	suite.Run(t, new(UnitTestSuite))
}

func (s *UnitTestSuite) Test_GenerateFeedsWorkflow() {
	env := s.NewTestWorkflowEnvironment()
	env.SetWorkerOptions(worker.Options{EnableSessionWorker: true})

	var a *feed_tasks.Activities

	host := "faker:8080"
	count := 10

	env.OnActivity(a.GenerateFeeds, mock.Anything, host, count).Return(nil)
	env.RegisterActivity(a)
	env.ExecuteWorkflow(feed_tasks.GenerateFeedsWorkflow, host, count)

	t := s.T()
	assert.True(t, env.IsWorkflowCompleted())
	require.NoError(t, env.GetWorkflowError())

	env.AssertExpectations(t)
}

func (s *UnitTestSuite) Test_RefreshFeedsWorkflow() {
	env := s.NewTestWorkflowEnvironment()
	env.SetWorkerOptions(worker.Options{EnableSessionWorker: true})

	var a *feed_tasks.Activities

	env.OnActivity(a.RefreshFeeds, mock.Anything).Return(nil)
	env.RegisterActivity(a)
	env.ExecuteWorkflow(feed_tasks.RefreshFeedsWorkflow)

	t := s.T()
	assert.True(t, env.IsWorkflowCompleted())
	require.NoError(t, env.GetWorkflowError())

	env.AssertExpectations(t)
}

func (s *UnitTestSuite) Test_AddFeedWorkflow() {
	env := s.NewTestWorkflowEnvironment()
	env.SetWorkerOptions(worker.Options{EnableSessionWorker: true})
	var a *feed_tasks.Activities

	url := "http://localhost/feed.xml"
	userID := "test-user-id"

	env.OnActivity(a.AddFeed, mock.Anything, url, userID).Return("feed-id-123", nil)
	env.RegisterActivity(a)
	env.ExecuteWorkflow(feed_tasks.AddFeedWorkflow, url, userID)

	t := s.T()
	assert.True(t, env.IsWorkflowCompleted())
	require.NoError(t, env.GetWorkflowError())

	env.AssertExpectations(t)
}
