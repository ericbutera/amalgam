package graph_test

import (
	"context"
	"testing"

	"github.com/ericbutera/amalgam/internal/converters"
	svcModel "github.com/ericbutera/amalgam/internal/service/models"
	"github.com/ericbutera/amalgam/internal/test/fixtures"
	pb "github.com/ericbutera/amalgam/pkg/feeds/v1"
	helpers "github.com/ericbutera/amalgam/pkg/test"
	"github.com/ericbutera/amalgam/services/graph/graph"
	graphModel "github.com/ericbutera/amalgam/services/graph/graph/model"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

func newClient() *pb.MockFeedServiceClient {
	return new(pb.MockFeedServiceClient)
}

func newResolver(client pb.FeedServiceClient) *graph.Resolver {
	return graph.NewResolver(client)
}

func newResolverWithClient() (*pb.MockFeedServiceClient, *graph.Resolver) {
	client := newClient()
	return client, newResolver(client)
}

func newFeed() *svcModel.Feed {
	return fixtures.NewFeed(fixtures.WithFeedID(fixtures.NewID()))
}

func newArticle() *svcModel.Article {
	return fixtures.NewArticle(fixtures.WithArticleID(fixtures.NewID()))
}

func Test_AddFeed(t *testing.T) {
	t.Parallel()
	client, resolver := newResolverWithClient()

	svcFeed := newFeed()

	client.EXPECT().
		CreateFeed(mock.Anything, &pb.CreateFeedRequest{
			Feed: &pb.CreateFeedRequest_Feed{
				Url:  svcFeed.URL,
				Name: svcFeed.Name,
			},
		}).
		Return(&pb.CreateFeedResponse{Id: svcFeed.ID}, nil)

	actual, err := resolver.Mutation().
		AddFeed(context.Background(), svcFeed.URL, svcFeed.Name)

	require.NoError(t, err)
	helpers.Diff(t, graphModel.AddResponse{ID: svcFeed.ID}, *actual)
}

func Test_UpdateFeed(t *testing.T) {
	t.Parallel()
	client, resolver := newResolverWithClient()

	svcFeed := newFeed()
	client.EXPECT().
		UpdateFeed(mock.Anything, &pb.UpdateFeedRequest{
			Feed: &pb.UpdateFeedRequest_Feed{
				Id:   svcFeed.ID,
				Url:  svcFeed.URL,
				Name: svcFeed.Name,
			},
		}).
		Return(nil, nil)

	actual, err := resolver.Mutation().
		UpdateFeed(context.Background(), svcFeed.ID, &svcFeed.URL, &svcFeed.Name)

	require.NoError(t, err)
	helpers.Diff(t, graphModel.UpdateResponse{ID: svcFeed.ID}, *actual)
}

func Test_Feeds(t *testing.T) {
	t.Parallel()
	client, resolver := newResolverWithClient()

	svcFeed := newFeed()
	c := converters.New()
	graphFeed := c.ServiceToGraphFeed(svcFeed)
	pbFeed := c.ServiceToProtoFeed(svcFeed)
	expected := []*graphModel.Feed{graphFeed}
	client.EXPECT().
		ListFeeds(mock.Anything, &pb.ListFeedsRequest{}).
		Return(&pb.ListFeedsResponse{
			Feeds: []*pb.Feed{pbFeed},
		}, nil)

	actual, err := resolver.Query().Feeds(context.Background())
	require.NoError(t, err)
	assert.Len(t, actual, 1)
	helpers.Diff(t, *expected[0], *actual[0])
}

func Test_Feed(t *testing.T) {
	t.Parallel()
	client, resolver := newResolverWithClient()

	svcFeed := newFeed()
	c := converters.New()
	graphFeed := c.ServiceToGraphFeed(svcFeed)
	pbFeed := c.ServiceToProtoFeed(svcFeed)
	expected := graphFeed
	client.EXPECT().
		GetFeed(mock.Anything, &pb.GetFeedRequest{Id: svcFeed.ID}).
		Return(&pb.GetFeedResponse{Feed: pbFeed}, nil)

	actual, err := resolver.Query().Feed(context.Background(), svcFeed.ID)
	require.NoError(t, err)
	helpers.Diff(t, *expected, *actual)
}

func Test_Articles(t *testing.T) {
	t.Parallel()
	client, resolver := newResolverWithClient()

	feed := newFeed()
	svcArticle := newArticle()

	c := converters.New()
	graphArticle := c.ServiceToGraphArticle(svcArticle)
	rpcArticle := c.ServiceToProtoArticle(svcArticle)
	expected := []*graphModel.Article{graphArticle}

	client.EXPECT().
		ListArticles(mock.Anything, &pb.ListArticlesRequest{FeedId: feed.ID}).
		Return(&pb.ListArticlesResponse{
			Articles: []*pb.Article{rpcArticle},
		}, nil)

	actual, err := resolver.Query().Articles(context.Background(), feed.ID)
	require.NoError(t, err)
	assert.Len(t, actual, 1)
	helpers.Diff(t, *expected[0], *actual[0], "FeedID", "ImageURL")
}

func Test_Article(t *testing.T) {
	t.Parallel()
	client, resolver := newResolverWithClient()

	c := converters.New()
	svcArticle := newArticle()
	graphArticle := c.ServiceToGraphArticle(svcArticle)
	rpcArticle := c.ServiceToProtoArticle(svcArticle)
	expected := graphArticle

	client.EXPECT().
		GetArticle(mock.Anything, &pb.GetArticleRequest{Id: svcArticle.ID}).
		Return(&pb.GetArticleResponse{
			Article: rpcArticle,
		}, nil)

	actual, err := resolver.Query().Article(context.Background(), svcArticle.ID)
	require.NoError(t, err)
	helpers.Diff(t, *expected, *actual, "FeedID", "ImageURL")
}
