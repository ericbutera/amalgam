package graph_test

import (
	"context"
	"testing"

	"github.com/ericbutera/amalgam/internal/converters"
	svcModel "github.com/ericbutera/amalgam/internal/service/models"
	"github.com/ericbutera/amalgam/internal/tasks"
	"github.com/ericbutera/amalgam/internal/test/fixtures"
	pb "github.com/ericbutera/amalgam/pkg/feeds/v1"
	helpers "github.com/ericbutera/amalgam/pkg/test"
	"github.com/ericbutera/amalgam/services/graph/graph"
	graphModel "github.com/ericbutera/amalgam/services/graph/graph/model"
	"github.com/google/uuid"
	"github.com/samber/lo"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

type testResolver struct {
	client   *pb.MockFeedServiceClient
	task     *tasks.MockTasks
	resolver *graph.Resolver
}

func newTestResolver() *testResolver {
	client := new(pb.MockFeedServiceClient)
	tasks := new(tasks.MockTasks)
	resolver := graph.NewResolver(client, tasks)
	return &testResolver{
		client:   client,
		task:     tasks,
		resolver: resolver,
	}
}

func newFeed() *svcModel.Feed {
	return fixtures.NewFeed(fixtures.WithFeedID(fixtures.NewID()))
}

func newArticle() *svcModel.Article {
	return fixtures.NewArticle(fixtures.WithArticleID(fixtures.NewID()))
}

func Test_AddFeed(t *testing.T) {
	t.Parallel()
	r := newTestResolver()

	svcFeed := newFeed()

	r.client.EXPECT().
		CreateFeed(mock.Anything, &pb.CreateFeedRequest{
			Feed: &pb.CreateFeedRequest_Feed{
				Url:  svcFeed.URL,
				Name: svcFeed.Name,
			},
		}).
		Return(&pb.CreateFeedResponse{Id: svcFeed.ID}, nil)

	actual, err := r.resolver.Mutation().
		AddFeed(context.Background(), svcFeed.URL, svcFeed.Name)

	require.NoError(t, err)
	helpers.Diff(t, graphModel.AddResponse{ID: svcFeed.ID}, *actual)
}

func Test_UpdateFeed(t *testing.T) {
	t.Parallel()
	r := newTestResolver()

	svcFeed := newFeed()
	r.client.EXPECT().
		UpdateFeed(mock.Anything, &pb.UpdateFeedRequest{
			Feed: &pb.UpdateFeedRequest_Feed{
				Id:   svcFeed.ID,
				Url:  svcFeed.URL,
				Name: svcFeed.Name,
			},
		}).
		Return(nil, nil)

	actual, err := r.resolver.Mutation().
		UpdateFeed(context.Background(), svcFeed.ID, &svcFeed.URL, &svcFeed.Name)

	require.NoError(t, err)
	helpers.Diff(t, graphModel.UpdateResponse{ID: svcFeed.ID}, *actual)
}

func Test_Feeds(t *testing.T) {
	t.Parallel()
	r := newTestResolver()

	svcFeed := newFeed()
	c := converters.New()
	graphFeed := c.ServiceToGraphFeed(svcFeed)
	pbFeed := c.ServiceToProtoFeed(svcFeed)
	expected := []*graphModel.Feed{graphFeed}
	r.client.EXPECT().
		ListFeeds(mock.Anything, &pb.ListFeedsRequest{}).
		Return(&pb.ListFeedsResponse{
			Feeds: []*pb.Feed{pbFeed},
		}, nil)

	actual, err := r.resolver.Query().Feeds(context.Background())
	require.NoError(t, err)
	assert.Len(t, actual, 1)
	helpers.Diff(t, *expected[0], *actual[0])
}

func Test_Feed(t *testing.T) {
	t.Parallel()
	r := newTestResolver()

	svcFeed := newFeed()
	c := converters.New()
	graphFeed := c.ServiceToGraphFeed(svcFeed)
	pbFeed := c.ServiceToProtoFeed(svcFeed)
	expected := graphFeed
	r.client.EXPECT().
		GetFeed(mock.Anything, &pb.GetFeedRequest{Id: svcFeed.ID}).
		Return(&pb.GetFeedResponse{Feed: pbFeed}, nil)

	actual, err := r.resolver.Query().Feed(context.Background(), svcFeed.ID)
	require.NoError(t, err)
	helpers.Diff(t, *expected, *actual)
}

func Test_Articles(t *testing.T) {
	t.Parallel()
	r := newTestResolver()

	feed := newFeed()
	svcArticle := newArticle()

	c := converters.New()
	graphArticle := c.ServiceToGraphArticle(svcArticle)
	rpcArticle := c.ServiceToProtoArticle(svcArticle)
	expected := []*graphModel.Article{graphArticle}

	r.client.EXPECT().
		ListArticles(mock.Anything, &pb.ListArticlesRequest{
			FeedId: feed.ID,
			Options: &pb.ListOptions{
				Cursor: "",
				Limit:  graph.DefaultLimit,
			},
		}).
		Return(&pb.ListArticlesResponse{
			Articles: []*pb.Article{rpcArticle},
		}, nil)

	resp, err := r.resolver.Query().Articles(context.Background(), feed.ID, &graphModel.ListOptions{})
	actual := resp.Articles
	require.NoError(t, err)
	assert.Len(t, actual, 1)
	helpers.Diff(t, *expected[0], *actual[0], "FeedID", "ImageURL")
}

func Test_Articles_Pagination(t *testing.T) {
	t.Parallel()
	r := newTestResolver()

	id := uuid.New().String()
	expectedCursor := "incoming-cursor"
	expectedLimit := 42
	expectedPagination := pb.Pagination{
		Next:     "next",
		Previous: "previous",
	}

	r.client.EXPECT().
		ListArticles(mock.Anything, &pb.ListArticlesRequest{
			FeedId: id,
			Options: &pb.ListOptions{
				Cursor: expectedCursor,
				Limit:  int32(expectedLimit),
			},
		}).
		Return(&pb.ListArticlesResponse{
			Articles:   []*pb.Article{},
			Pagination: &expectedPagination,
		}, nil)

	resp, err := r.resolver.Query().Articles(context.Background(), id, &graphModel.ListOptions{
		Cursor: lo.ToPtr(expectedCursor),
		Limit:  lo.ToPtr(expectedLimit),
	})
	require.NoError(t, err)
	assert.Equal(t, expectedPagination.GetNext(), resp.Pagination.Next)
	assert.Equal(t, expectedPagination.GetPrevious(), resp.Pagination.Previous)
}

func Test_Article(t *testing.T) {
	t.Parallel()
	r := newTestResolver()

	c := converters.New()
	svcArticle := newArticle()
	graphArticle := c.ServiceToGraphArticle(svcArticle)
	rpcArticle := c.ServiceToProtoArticle(svcArticle)
	expected := graphArticle

	r.client.EXPECT().
		GetArticle(mock.Anything, &pb.GetArticleRequest{Id: svcArticle.ID}).
		Return(&pb.GetArticleResponse{
			Article: rpcArticle,
		}, nil)

	actual, err := r.resolver.Query().Article(context.Background(), svcArticle.ID)
	require.NoError(t, err)
	helpers.Diff(t, *expected, *actual, "FeedID", "ImageURL")
}

func TestFeedTask(t *testing.T) {
	t.Parallel()
	ctx := context.Background()
	expectedID := "super-id"

	r := newTestResolver()

	r.task.EXPECT().
		Workflow(mock.Anything, tasks.TaskGenerateFeeds).
		Return(&tasks.TaskResult{ID: expectedID}, nil)

	resp, err := r.resolver.Mutation().GenerateFeeds(ctx)

	require.NoError(t, err)
	assert.Equal(t, expectedID, resp.ID)
}
