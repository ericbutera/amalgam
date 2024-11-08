package app

import (
	"testing"

	"github.com/stretchr/testify/mock"
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
	var a *Activities

	feed_id := "213ddff2-e7cc-40cc-87eb-461118d57a58"
	feed_url := "http://faker:8080/feed/e568f1fa-a0e9-4545-bc5b-a167725a75bd"

	env.OnActivity(a.DownloadActivity, mock.Anything, feed_id, feed_url).Return("rss_file", nil)
	env.OnActivity(a.ParseActivity, mock.Anything, feed_id, "rss_file").Return("articles_file", nil)
	env.OnActivity(a.SaveActivity, mock.Anything, feed_id, "articles_file").Return(nil)

	env.RegisterActivity(a)

	env.ExecuteWorkflow(FeedWorkflow, feed_id, feed_url)

	s.True(env.IsWorkflowCompleted())
	s.NoError(env.GetWorkflowError())

	env.AssertExpectations(s.T())
}
