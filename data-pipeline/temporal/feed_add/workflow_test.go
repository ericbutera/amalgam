package feed_add_test

import (
	"testing"

	app "github.com/ericbutera/amalgam/data-pipeline/temporal/feed_add"
	"github.com/ericbutera/amalgam/internal/test/fixtures"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
	"go.temporal.io/sdk/client"
	"go.temporal.io/sdk/testsuite"
	"go.temporal.io/sdk/worker"
)

type FeedAddWorkflowTestSuite struct {
	suite.Suite
	testsuite.WorkflowTestSuite
}

func TestFeedAddWorkflowSuite(t *testing.T) {
	suite.Run(t, new(FeedAddWorkflowTestSuite))
}

func (s *FeedAddWorkflowTestSuite) Test_FeedAddWorkflow() {
	workflowID := "test-workflow-id"

	env := s.NewTestWorkflowEnvironment()
	env.SetWorkerOptions(worker.Options{
		EnableSessionWorker: true,
	})
	env.SetStartWorkflowOptions(client.StartWorkflowOptions{
		ID: workflowID,
	})

	var a *app.Activities

	data := fixtures.NewFeedVerification()
	verification := app.FeedVerification{
		ID:         0,
		URL:        data.URL,
		UserID:     data.UserID,
		WorkflowID: workflowID,
	}

	env.OnActivity(a.CreateVerifyRecord, mock.Anything, verification /*verification.URL, verification.UserID, workflowID*/).Return(&verification, nil)
	env.OnActivity(a.Fetch, mock.Anything, verification).Return("rss_file", nil)
	env.OnActivity(a.CreateFeed, mock.Anything, verification).Return(nil)
	env.RegisterActivity(a)
	env.ExecuteWorkflow(app.AddFeedWorkflow, verification.URL, verification.UserID)

	t := s.T()
	require.NoError(t, env.GetWorkflowError())
	require.True(t, env.IsWorkflowCompleted())
	env.AssertExpectations(t)
}
