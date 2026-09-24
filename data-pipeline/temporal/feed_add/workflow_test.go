package feed_add_test

import (
	"testing"

	app "github.com/ericbutera/amalgam/data-pipeline/temporal/feed_add"
	"github.com/ericbutera/amalgam/internal/test/fixtures"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
	"go.temporal.io/sdk/client"
	"go.temporal.io/sdk/temporal"
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

	// Subscribe attempted first; simulate not found by returning error so
	// workflow proceeds to verification flow.
	env.OnActivity(a.SubscribeUserToUrl, mock.Anything, verification.URL, verification.UserID).
		Return("", temporal.NewNonRetryableApplicationError("feed not found", app.FeedNotFoundErrorType, nil)).Once()
	env.OnActivity(a.CreateVerifyRecord, mock.Anything, verification).Return(&verification, nil)
	env.OnActivity(a.Fetch, mock.Anything, verification).Return("rss_file", nil)
	env.OnActivity(a.CreateFeed, mock.Anything, verification).Return("test-feed-id", nil)
	env.OnActivity(a.SubscribeUserToUrl, mock.Anything, verification.URL, verification.UserID).Return("test-feed-id", nil).Once()
	env.RegisterActivity(a)
	env.ExecuteWorkflow(app.AddFeedWorkflow, verification.URL, verification.UserID)

	t := s.T()
	require.NoError(t, env.GetWorkflowError())
	require.True(t, env.IsWorkflowCompleted())
	var feedID string
	require.NoError(t, env.GetWorkflowResult(&feedID))
	require.Equal(t, "test-feed-id", feedID)
	env.AssertExpectations(t)
}

func (s *FeedAddWorkflowTestSuite) Test_FeedAddWorkflow_ExistingFeed() {
	t := s.T()
	env := s.NewTestWorkflowEnvironment()
	env.SetStartWorkflowOptions(client.StartWorkflowOptions{ID: "existing-feed-workflow-id"})

	var a *app.Activities
	env.OnActivity(a.SubscribeUserToUrl, mock.Anything, "https://example.com/feed.xml", "test-user-id").Return("existing-feed-id", nil).Once()
	env.RegisterActivity(a)
	env.ExecuteWorkflow(app.AddFeedWorkflow, "https://example.com/feed.xml", "test-user-id")

	require.NoError(t, env.GetWorkflowError())
	var feedID string
	require.NoError(t, env.GetWorkflowResult(&feedID))
	require.Equal(t, "existing-feed-id", feedID)
	env.AssertExpectations(t)
}
