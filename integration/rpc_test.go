//go:build integration
// +build integration

package integration_test

import (
	"context"
	"os"
	"testing"

	pb "github.com/ericbutera/amalgam/pkg/feeds/v1"
	"github.com/ericbutera/amalgam/rpc/pkg/client"
	"github.com/samber/lo"
	"github.com/stretchr/testify/require"
)

func getClient(t *testing.T) (pb.FeedServiceClient, client.Closer, error) {
	config := newConfig(t)
	c, closer, err := client.New(config.target, config.useInsecure)
	return c, closer, err
}

type Config struct {
	target      string
	useInsecure bool
}

func newConfig(t *testing.T) *Config {
	// if testing.Short() {
	// 	t.Skip("skipping test in short mode.")
	// }
	c := &Config{
		target:      os.Getenv("RPC_HOST"),
		useInsecure: lo.Ternary(os.Getenv("RPC_INSECURE") == "true", true, false),
	}
	if c.target == "" {
		t.Skip("RPC_ENDPOINT not set")
	}
	return c
}

func TestRpc(t *testing.T) {
	client, closer, err := getClient(t)
	defer func() { require.NoError(t, closer()) }()
	require.NoError(t, err)

	resp, err := client.ListFeeds(context.Background(), &pb.ListFeedsRequest{})
	require.NoError(t, err)
	require.NotNil(t, resp)
}
