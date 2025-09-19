package graph_test

import (
	"context"
	"testing"
	"time"

	"github.com/ericbutera/amalgam/internal/converters"
	svcModel "github.com/ericbutera/amalgam/internal/service/models"
	"github.com/ericbutera/amalgam/internal/tasks"
	"github.com/ericbutera/amalgam/internal/test/fixtures"
	"github.com/ericbutera/amalgam/internal/test/seed"
	pb "github.com/ericbutera/amalgam/pkg/feeds/v1"
	helpers "github.com/ericbutera/amalgam/pkg/test"
	"github.com/ericbutera/amalgam/services/graph/graph"
	graphModel "github.com/ericbutera/amalgam/services/graph/graph/model"
	"github.com/ericbutera/amalgam/services/graph/internal/middleware"
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

func newAuthCtx() context.Context {
	return middleware.WithUserID(context.Background(), seed.UserID) // TODO: a better way to do this would be DI auth provider in the resolver
}

func newFeed() *svcModel.Feed {
	return fixtures.NewFeed(fixtures.WithFeedID(fixtures.NewID()))
}

func newArticle() *svcModel.Article {
	return fixtures.NewArticle(fixtures.WithArticleID(fixtures.NewID()))
}

func newUserFeed() *svcModel.UserFeed {
	now := time.Now().UTC()
	id := uuid.New().String()

	return &svcModel.UserFeed{
		FeedID:        id,
		Name:          "Feed Name",
		URL:           "https://faker:8080/feeds/" + id,
		CreatedAt:     now,
		ViewedAt:      now,
		UnreadStartAt: now,
		UnreadCount:   1,
	}
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
			User: &pb.User{
				Id: seed.UserID,
			},
		}).
		Return(&pb.CreateFeedResponse{Id: svcFeed.ID}, nil)

	actual, err := r.resolver.Mutation().
		AddFeed(newAuthCtx(), svcFeed.URL, svcFeed.Name)

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
		UpdateFeed(newAuthCtx(), svcFeed.ID, &svcFeed.URL, &svcFeed.Name)

	require.NoError(t, err)
	helpers.Diff(t, graphModel.UpdateResponse{ID: svcFeed.ID}, *actual)
}

func Test_Feeds(t *testing.T) {
	t.Parallel()

	r := newTestResolver()

	userFeed := newUserFeed()
	c := converters.New()
	graphFeed := c.ServiceToGraphFeed(userFeed)
	pbFeed := c.ServiceToProtoUserFeed(userFeed)
	expected := []*graphModel.Feed{graphFeed}

	r.client.EXPECT().
		ListUserFeeds(mock.Anything, &pb.ListUserFeedsRequest{
			User: &pb.User{Id: seed.UserID},
		}).
		Return(&pb.ListUserFeedsResponse{
			Feeds: []*pb.UserFeed{pbFeed},
		}, nil)

	actual, err := r.resolver.Query().Feeds(newAuthCtx())
	require.NoError(t, err)
	assert.Len(t, actual.Feeds, 1)
	helpers.Diff(t, *expected[0], *actual.Feeds[0])
}

func Test_Feed(t *testing.T) {
	t.Parallel()

	r := newTestResolver()

	feed := newUserFeed()
	c := converters.New()
	graphFeed := c.ServiceToGraphFeed(feed)
	pbFeed := c.ServiceToProtoUserFeed(feed)
	expected := graphFeed

	r.client.EXPECT().
		GetUserFeed(mock.Anything, &pb.GetUserFeedRequest{FeedId: feed.FeedID, UserId: seed.UserID}).
		Return(&pb.GetUserFeedResponse{Feed: pbFeed}, nil)

	actual, err := r.resolver.Query().Feed(newAuthCtx(), feed.FeedID)
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
			FeedId:  feed.ID,
			Options: &pb.ListOptions{Limit: 0},
		}).
		Return(&pb.ListArticlesResponse{
			Articles: []*pb.Article{rpcArticle},
		}, nil)

	r.client.EXPECT().
		GetUserArticles(mock.Anything, &pb.GetUserArticlesRequest{
			User:       &pb.User{Id: seed.UserID},
			ArticleIds: []string{svcArticle.ID},
		}).
		Return(&pb.GetUserArticlesResponse{
			Articles: map[string]*pb.GetUserArticlesResponse_UserArticle{
				svcArticle.ID: nil,
			},
		}, nil)

	resp, err := r.resolver.Query().Articles(newAuthCtx(), feed.ID, &graphModel.ListOptions{})
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
	expectedPagination := pb.Cursor{
		Next:     "next",
		Previous: "previous",
	}

	r.client.EXPECT().
		ListArticles(mock.Anything, &pb.ListArticlesRequest{
			FeedId: id,
			Options: &pb.ListOptions{
				Cursor: &pb.Cursor{Next: expectedCursor},
				Limit:  int32(expectedLimit),
			},
		}).
		Return(&pb.ListArticlesResponse{
			Articles: []*pb.Article{},
			Cursor:   &expectedPagination,
		}, nil)

	resp, err := r.resolver.Query().Articles(newAuthCtx(), id, &graphModel.ListOptions{
		Cursor: &graphModel.ListCursor{
			Next: lo.ToPtr(expectedCursor),
		},
		Limit: lo.ToPtr(expectedLimit),
	})
	require.NoError(t, err)
	assert.Equal(t, expectedPagination.GetNext(), resp.Cursor.Next)
	assert.Equal(t, expectedPagination.GetPrevious(), resp.Cursor.Previous)
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

	actual, err := r.resolver.Query().Article(newAuthCtx(), svcArticle.ID)
	require.NoError(t, err)
	helpers.Diff(t, *expected, *actual, "FeedID", "ImageURL")
}

func TestFeedTasks(t *testing.T) {
	t.Parallel()

	cases := []struct {
		name      string
		taskType  tasks.TaskType
		graphType graphModel.TaskType
	}{
		{"generate feeds", tasks.TaskGenerateFeeds, graphModel.TaskTypeGenerateFeeds},
		{"fetch feeds", tasks.TaskFetchFeeds, graphModel.TaskTypeRefreshFeeds},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			expectedID := "super-id"

			r := newTestResolver()

			r.task.EXPECT().
				Workflow(mock.Anything, tc.taskType).
				Return(&tasks.TaskResult{ID: expectedID}, nil)

			resp, err := r.resolver.Mutation().
				FeedTask(newAuthCtx(), tc.graphType)

			require.NoError(t, err)
			assert.Equal(t, expectedID, resp.TaskID)
		})
	}
}
