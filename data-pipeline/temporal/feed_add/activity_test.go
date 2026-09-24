package feed_add_test

import (
	"testing"

	app "github.com/ericbutera/amalgam/data-pipeline/temporal/feed_add"
	"github.com/ericbutera/amalgam/data-pipeline/temporal/internal/bucket"
	"github.com/ericbutera/amalgam/internal/http/fetch"
	pb "github.com/ericbutera/amalgam/pkg/feeds/v1"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
	"go.temporal.io/sdk/temporal"
	"go.temporal.io/sdk/testsuite"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func newActivityEnv(activities interface{}) *testsuite.TestActivityEnvironment {
	var suite testsuite.WorkflowTestSuite
	env := suite.NewTestActivityEnvironment()
	env.RegisterActivity(activities)
	return env
}

type activitySetup struct {
	fetcher    *fetch.MockFetch
	bucket     *bucket.MockBucket
	activities *app.Activities
	rpc        *pb.MockFeedServiceClient
}

func setupActivities(t *testing.T) *activitySetup {
	fetcher := fetch.NewMockFetch(t)
	bucketClient := bucket.NewMockBucket(t)
	rpcClient := new(pb.MockFeedServiceClient)

	return &activitySetup{
		fetcher:    fetcher,
		bucket:     bucketClient,
		rpc:        rpcClient,
		activities: app.NewActivities(fetcher, bucketClient, rpcClient),
	}
}

func NewFeedVerification() app.FeedVerification {
	return app.FeedVerification{
		URL:        "http://localhost/feed.xml",
		UserID:     "test-user-id",
		WorkflowID: "test-workflow-id",
	}
}

func TestCreateVerifyRecord(t *testing.T) {
	t.Parallel()
	s := setupActivities(t)
	expected := NewFeedVerification()
	pbVerification := &pb.FeedVerification{
		Url:        expected.URL,
		UserId:     expected.UserID,
		WorkflowId: expected.WorkflowID,
	}

	s.rpc.EXPECT().
		CreateFeedVerification(mock.Anything, &pb.CreateFeedVerificationRequest{
			Verification: pbVerification,
		}).
		Return(&pb.CreateFeedVerificationResponse{
			Verification: pbVerification,
		}, nil)

	env := newActivityEnv(s.activities)

	var actual *app.FeedVerification
	future, actErr := env.ExecuteActivity(s.activities.CreateVerifyRecord, expected)
	require.NoError(t, actErr)
	err := future.Get(&actual)

	require.NoError(t, err)
	require.Equal(t, expected.URL, actual.URL)
	require.Equal(t, expected.UserID, actual.UserID)
	require.Equal(t, expected.WorkflowID, actual.WorkflowID)
}

func TestFetch(t *testing.T) {
	t.Parallel()
	s := setupActivities(t)
	data := NewFeedVerification()
	s.fetcher.EXPECT().
		Url(mock.Anything, data.URL, mock.Anything, mock.Anything). // Assert URL param
		Return(nil)
	env := newActivityEnv(s.activities)

	future, actErr := env.ExecuteActivity(s.activities.Fetch, data)
	require.NoError(t, actErr)
	var res string
	err := future.Get(&res)

	require.NoError(t, err)
}

func TestCreateFeed(t *testing.T) {
	t.Parallel()
	s := setupActivities(t)
	data := NewFeedVerification()
	s.rpc.EXPECT().
		CreateFeed(mock.Anything, &pb.CreateFeedRequest{
			Feed: &pb.CreateFeedRequest_Feed{
				Url: data.URL,
			},
			User: &pb.User{Id: data.UserID},
		}).
		Return(&pb.CreateFeedResponse{
			Id: "test-feed-id",
		}, nil)

	env := newActivityEnv(s.activities)
	future, actErr := env.ExecuteActivity(s.activities.CreateFeed, data)
	require.NoError(t, actErr)
	var resp string
	err := future.Get(&resp)

	require.NoError(t, err)
	require.Equal(t, "test-feed-id", resp)
}

func TestSubscribeUserToUrl(t *testing.T) {
	t.Parallel()
	s := setupActivities(t)
	data := NewFeedVerification()
	s.rpc.EXPECT().
		SubscribeUserToUrl(mock.Anything, &pb.SubscribeUserToUrlRequest{
			Url:  data.URL,
			User: &pb.User{Id: data.UserID},
		}).
		Return(&pb.SubscribeUserToUrlResponse{
			FeedId: "test-feed-id",
		}, nil)

	env := newActivityEnv(s.activities)
	future, actErr := env.ExecuteActivity(s.activities.SubscribeUserToUrl, data.URL, data.UserID)
	require.NoError(t, actErr)
	var resp string
	err := future.Get(&resp)

	require.NoError(t, err)
	require.Equal(t, "test-feed-id", resp)
}

func TestSubscribeUserToUrl_FeedNotFound(t *testing.T) {
	t.Parallel()
	s := setupActivities(t)
	data := NewFeedVerification()
	s.rpc.EXPECT().
		SubscribeUserToUrl(mock.Anything, mock.Anything).
		Return(nil, status.Error(codes.NotFound, "not found"))

	env := newActivityEnv(s.activities)
	_, actErr := env.ExecuteActivity(s.activities.SubscribeUserToUrl, data.URL, data.UserID)
	require.Error(t, actErr)
	require.True(t, temporal.IsApplicationError(actErr))
	require.Contains(t, actErr.Error(), "type: "+app.FeedNotFoundErrorType)
}
