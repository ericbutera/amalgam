//go:build integration
// +build integration

package integration_test

// Graph is the entrypoint for the entire system. It is the best place to build out coverage.

import (
	"context"
	"net/http"
	"os"
	"testing"

	"github.com/Khan/genqlient/graphql"
	graphClient "github.com/ericbutera/amalgam/pkg/clients/graphql"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// TODO: ability to seed specific data for testing
// TODO: truncate between tests

func TestGraphListFeeds(t *testing.T) {
	t.Parallel()
	res, err := graphClient.ListFeeds(context.Background(), getGraphQLClient(t))
	require.NoError(t, err)
	assert.NotEmpty(t, res.Feeds) // populated by seed data
}

func TestGraphGetFeedMissingID(t *testing.T) {
	t.Parallel()
	_, err := graphClient.GetFeed(context.Background(), getGraphQLClient(t), "uid")
	require.Error(t, err)
}

func TestGraphGetFeedNotFound(t *testing.T) {
	t.Parallel()
	_, err := graphClient.GetFeed(context.Background(), getGraphQLClient(t), "e97f8e74-1183-4280-a48d-dd592e013ee1")
	require.Error(t, err)
}

func TestGraphAddFeed_InvalidURL(t *testing.T) {
	t.Parallel()
	_, err := graphClient.AddFeed(context.Background(), getGraphQLClient(t), "invalid-url", "name")
	// "input: value must be a valid URI"
	// "input: validation error"
	require.Error(t, err)
}

func TestListArticlesMissingID(t *testing.T) {
	t.Parallel()
	_, err := graphClient.ListArticles(context.Background(), getGraphQLClient(t), "uid")
	require.Error(t, err)
}

func TestGraphGetArticleMissingID(t *testing.T) {
	t.Parallel()
	_, err := graphClient.GetArticle(context.Background(), getGraphQLClient(t), "")
	require.Error(t, err)
}

func getGraphQLClient(t *testing.T) graphql.Client {
	t.Helper()
	target := os.Getenv("GRAPH_HOST")
	if target == "" {
		t.Skip("skipping test; GRAPH_HOST is not set")
	}

	httpClient := http.Client{}
	return graphql.NewClient(
		target,
		&httpClient,
	)
}
