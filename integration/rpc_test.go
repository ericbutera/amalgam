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
	"github.com/stretchr/testify/require"
)

func TestRpcListFeeds(t *testing.T) {
	t.Parallel()
	// TODO: seed data, assert specific result exists
	client, closer, err := getRpcClient(t)
	defer func() { require.NoError(t, closer()) }()
	require.NoError(t, err)

	resp, err := client.ListFeeds(context.Background(), &pb.ListFeedsRequest{})
	require.NoError(t, err)
	require.NotNil(t, resp)
}

func TestRpcGetFeed(t *testing.T) {
	t.Parallel()
	client, closer, err := getRpcClient(t)
	defer func() { require.NoError(t, closer()) }()
	require.NoError(t, err)

	_, err = client.GetFeed(context.Background(), &pb.GetFeedRequest{Id: "e97f8e74-1183-4280-a48d-dd592e013ee1"})
	require.Error(t, err)
}

func TestRpcGetUserFeed(t *testing.T) {
	t.Parallel()
	client, closer, err := getRpcClient(t)
	defer func() { require.NoError(t, closer()) }()
	require.NoError(t, err)

	_, err = client.GetUserFeed(context.Background(), &pb.GetUserFeedRequest{
		UserId: "e97f8e74-1183-4280-a48d-dd592e013ee1",
		FeedId: "e97f8e74-1183-4280-a48d-dd592e013ee1",
	})
	require.Error(t, err)
}

func getRpcClient(t *testing.T) (pb.FeedServiceClient, client.Closer, error) {
	target := os.Getenv("RPC_HOST")
	useInsecure := lo.Ternary(os.Getenv("RPC_INSECURE") == "true", true, false)
	require.NotEmpty(t, target, "RPC_HOST not set")

	c, closer, err := client.New(target, useInsecure)
	return c, closer, err
}

type Config struct {
	target      string
	useInsecure bool
}
