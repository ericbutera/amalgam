package app_test

import (
	"testing"

	app "github.com/ericbutera/amalgam/data-pipeline/temporal/feed"
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

func (s *UnitTestSuite) Test_FeedWorkflow() {
	env := s.NewTestWorkflowEnvironment()
	env.SetWorkerOptions(worker.Options{
		EnableSessionWorker: true, // Important for a worker to participate in the session
	})
	var a *app.Activities

	feedID := "213ddff2-e7cc-40cc-87eb-461118d57a58"
	feedURL := "http://faker:8080/feed/e568f1fa-a0e9-4545-bc5b-a167725a75bd"

	env.OnActivity(a.DownloadActivity, mock.Anything, feedID, feedURL).Return("rss_file", nil)
	env.OnActivity(a.ParseActivity, mock.Anything, feedID, "rss_file").Return("articles_file", nil)
	env.OnActivity(a.SaveActivity, mock.Anything, feedID, "articles_file").Return(app.SaveResults{}, nil)
	env.RegisterActivity(a)
	env.ExecuteWorkflow(app.FeedWorkflow, feedID, feedURL)

	t := s.T()
	require.NoError(t, env.GetWorkflowError())
	assert.True(t, env.IsWorkflowCompleted())

	env.AssertExpectations(t)
}