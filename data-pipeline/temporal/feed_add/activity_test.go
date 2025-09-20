package feed_add_test

import (
	"context"
	"testing"

	app "github.com/ericbutera/amalgam/data-pipeline/temporal/feed_add"
	"github.com/ericbutera/amalgam/data-pipeline/temporal/internal/bucket"
	"github.com/ericbutera/amalgam/internal/http/fetch"
	pb "github.com/ericbutera/amalgam/pkg/feeds/v1"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

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

	actual, err := s.activities.CreateVerifyRecord(context.Background(), expected)

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

	_, err := s.activities.Fetch(context.Background(), data)

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
		Return(&pb.CreateFeedResponse{}, nil)

	_, err := s.activities.CreateFeed(context.Background(), data)

	require.NoError(t, err)
}
