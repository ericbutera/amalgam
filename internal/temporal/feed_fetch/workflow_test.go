package feed_fetch_test

import (
	"testing"

	app "github.com/ericbutera/amalgam/internal/temporal/feed_fetch"
	"github.com/ericbutera/amalgam/internal/temporal/feeds"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
	"go.temporal.io/sdk/testsuite"
	"go.temporal.io/sdk/worker"
)

type FeedFetchTestSuite struct {
	suite.Suite
	testsuite.WorkflowTestSuite
}

func TestFeedFetchTestSuite(t *testing.T) {
	suite.Run(t, new(FeedFetchTestSuite))
}

func (s *FeedFetchTestSuite) Test_FeedWorkflow() {
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
	env.OnActivity(a.StatsActivity, mock.Anything, feedID).Return(nil)
	env.RegisterActivity(a)
	env.ExecuteWorkflow(app.FeedWorkflowV2, app.FeedInput{FeedID: feedID, URL: feedURL})

	t := s.T()
	require.NoError(t, env.GetWorkflowError())
	assert.True(t, env.IsWorkflowCompleted())
	env.AssertExpectations(t)
}

func (s *FeedFetchTestSuite) Test_FeedWorkflowFailsWhenSavePartiallyFails() {
	env := s.NewTestWorkflowEnvironment()

	var a *app.Activities
	feedID := "213ddff2-e7cc-40cc-87eb-461118d57a58"
	feedURL := "http://faker:8080/feed/e568f1fa-a0e9-4545-bc5b-a167725a75bd"

	env.OnActivity(a.DownloadActivity, mock.Anything, feedID, feedURL).Return("rss_file", nil)
	env.OnActivity(a.ParseActivity, mock.Anything, feedID, "rss_file").Return("articles_file", nil)
	env.OnActivity(a.SaveActivity, mock.Anything, feedID, "articles_file").Return(app.SaveResults{Succeeded: 1, Failed: 1}, nil)
	env.RegisterActivity(a)
	env.ExecuteWorkflow(app.FeedWorkflowV2, app.FeedInput{FeedID: feedID, URL: feedURL})

	t := s.T()
	require.Error(t, env.GetWorkflowError())
	assert.Contains(t, env.GetWorkflowError().Error(), app.PartialSaveErrorType)
	env.AssertExpectations(t)
}

func (s *FeedFetchTestSuite) Test_FetchFeedsWorkflow() {
	urls := []feeds.Feed{{ID: "213ddff2-e7cc-40cc-87eb-461118d57a58", Url: "http://faker:8080/feed/e568f1fa-a0e9-4545-bc5b-a167725a75bd"}}
	env := s.NewTestWorkflowEnvironment()
	env.SetWorkerOptions(worker.Options{
		EnableSessionWorker: true,
	})

	var a *app.Activities

	env.OnActivity(a.GetFeedsActivity, mock.Anything).
		Return(urls, nil)

	env.RegisterActivity(a)
	env.OnWorkflow(app.FeedWorkflowV2, mock.Anything, mock.Anything).Return(nil)
	env.RegisterWorkflow(app.FetchFeedsWorkflow)
	env.ExecuteWorkflow(app.FetchFeedsWorkflow)

	t := s.T()
	require.NoError(t, env.GetWorkflowError())
	assert.True(t, env.IsWorkflowCompleted())
	env.AssertExpectations(t)
}

func (s *FeedFetchTestSuite) Test_FetchFeedsWorkflowCapturesChildFailures() {
	urls := []feeds.Feed{
		{ID: "feed-one", Url: "http://faker:8080/feed/one"},
		{ID: "feed-two", Url: "http://faker:8080/feed/two"},
	}
	env := s.NewTestWorkflowEnvironment()

	var a *app.Activities
	env.OnActivity(a.GetFeedsActivity, mock.Anything).Return(urls, nil)
	env.RegisterActivity(a)
	env.OnWorkflow(app.FeedWorkflowV2, mock.Anything, app.FeedInput{FeedID: "feed-one", URL: "http://faker:8080/feed/one"}).Return(assert.AnError)
	env.OnWorkflow(app.FeedWorkflowV2, mock.Anything, app.FeedInput{FeedID: "feed-two", URL: "http://faker:8080/feed/two"}).Return(nil)
	env.RegisterWorkflow(app.FetchFeedsWorkflow)
	env.ExecuteWorkflow(app.FetchFeedsWorkflow)

	var result app.FetchFeedsResult
	t := s.T()
	require.NoError(t, env.GetWorkflowResult(&result))
	assert.Equal(t, 2, result.Total)
	assert.Equal(t, 1, result.Succeeded)
	assert.Equal(t, 1, result.Failed)
	assert.Equal(t, "WorkflowFailure", result.Failures[0].ErrorType)
	assert.Equal(t, "feed", result.Failures[0].Stage)
	assert.True(t, result.Failures[0].Retryable)
	env.AssertExpectations(t)
}
