//go:build integration
// +build integration

package integration_test

// TODO: revisit test location
// TODO: create a database cleanup function (truncate feed/article tables)
// TODO: convert to test suite

import (
	"context"
	"os"
	"testing"

	pb "github.com/ericbutera/amalgam/pkg/feeds/v1"
	"github.com/ericbutera/amalgam/services/rpc/pkg/client"
	"github.com/samber/lo"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"google.golang.org/genproto/googleapis/rpc/errdetails"
	"google.golang.org/grpc/status"
)

const (
	TestUserID = "e97f8e74-1183-4280-a48d-dd592e013ee1"
	TestFeedID = "e97f8e74-1183-4280-a48d-dd592e013ee1"
)

func TestRpcListFeeds(t *testing.T) {
	t.Parallel()
	// TODO: seed data, assert specific result exists
	client, closer := getRpcClient(t)
	defer func() { require.NoError(t, closer()) }()

	resp, err := client.ListFeeds(context.Background(), &pb.ListFeedsRequest{})
	require.NoError(t, err)
	require.NotNil(t, resp)
}

func TestRpcGetFeed(t *testing.T) {
	t.Parallel()
	client, closer := getRpcClient(t)
	defer func() { require.NoError(t, closer()) }()

	_, err := client.GetFeed(context.Background(), &pb.GetFeedRequest{Id: TestFeedID})
	require.Error(t, err)
}

func TestRpcGetUserFeed(t *testing.T) {
	t.Parallel()
	client, closer := getRpcClient(t)
	defer func() { require.NoError(t, closer()) }()

	_, err := client.GetUserFeed(context.Background(), &pb.GetUserFeedRequest{
		UserId: TestUserID,
		FeedId: TestFeedID,
	})
	require.Error(t, err)
}

func TestRpcAddFeed_InvalidURL(t *testing.T) {
	t.Parallel()
	client, closer := getRpcClient(t)
	defer func() { require.NoError(t, closer()) }()

	_, err := client.CreateFeed(context.Background(), &pb.CreateFeedRequest{
		Feed: &pb.CreateFeedRequest_Feed{
			Url:  "invalid-url",
			Name: "name",
		},
		User: &pb.User{Id: TestUserID},
	})
	// URI" does not contain "rpc error: code = InvalidArgument desc = validation error:\n - feed.url: value must be a valid URI [string.uri]
	require.Error(t, err)
	assert.Contains(t, err.Error(), "validation failed")

	s := status.Convert(err)
	for _, detail := range s.Details() {
		if br, ok := detail.(*errdetails.BadRequest); ok {
			assert.Len(t, br.GetFieldViolations(), 1)
			violations := br.GetFieldViolations()
			assert.Equal(t, "feed.url", violations[0].GetField()) // error comes from protovalidate
			return
		}
	}
	assert.Fail(t, "validation error not found")
}

func getRpcClient(t *testing.T) (pb.FeedServiceClient, client.Closer) {
	t.Helper()
	target := os.Getenv("RPC_HOST")
	useInsecure := lo.Ternary(os.Getenv("RPC_INSECURE") == "true", true, false)
	if target == "" {
		t.Skip("skipping test; RPC_HOST not set")
		return nil, nil
	}

	c, closer, err := client.New(target, useInsecure)
	require.NoError(t, err)
	return c, closer
}
