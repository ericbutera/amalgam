package server_test

import (
	"context"
	"testing"

	"github.com/ericbutera/amalgam/internal/converters"
	"github.com/ericbutera/amalgam/internal/service"
	svcModel "github.com/ericbutera/amalgam/internal/service/models"
	"github.com/ericbutera/amalgam/internal/test"
	"github.com/ericbutera/amalgam/internal/test/fixtures"
	"github.com/ericbutera/amalgam/internal/test/seed"
	"github.com/ericbutera/amalgam/pkg/config/env"
	pb "github.com/ericbutera/amalgam/pkg/feeds/v1"
	helpers "github.com/ericbutera/amalgam/pkg/test"
	"github.com/ericbutera/amalgam/services/rpc/internal/config"
	"github.com/ericbutera/amalgam/services/rpc/internal/server"
	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
	"github.com/samber/lo"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
	"google.golang.org/genproto/googleapis/rpc/errdetails"
	"google.golang.org/grpc"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"
	"gorm.io/gorm"
)

type TestServer struct {
	Server  *server.Server
	Service service.Service
	DB      *gorm.DB
}

func newServer(t *testing.T) *TestServer {
	t.Helper()
	config := lo.Must(env.New[config.Config]())
	db := test.NewDB(t)
	svc := service.NewGorm(db)
	server := lo.Must(server.New(
		server.WithDb(db),
		server.WithService(svc),
		server.WithConfig(config),
	))
	s := grpc.NewServer()
	pb.RegisterFeedServiceServer(s, server)
	return &TestServer{
		Server:  server,
		Service: svc,
		DB:      db,
	}
}

type testMockServer struct {
	svc        *service.MockService
	server     *server.Server
	converters converters.Converter
}

func newMockServer(t *testing.T) *testMockServer {
	t.Helper()

	svc := new(service.MockService)
	s, err := server.New(server.WithService(svc))
	require.NoError(t, err)

	return &testMockServer{
		svc:        svc,
		server:     s,
		converters: converters.New(),
	}
}

func TestCreateFeedValidateError(t *testing.T) {
	t.Parallel()

	ts := newServer(t)
	// TODO: table test to assert all validation errors
	ctx := context.Background()
	_, err := ts.Server.CreateFeed(ctx, &pb.CreateFeedRequest{
		Feed: &pb.CreateFeedRequest_Feed{
			Name: "a",
			Url:  "invalid url",
		},
	})
	require.Error(t, err)
	s := status.Convert(err)
	for _, detail := range s.Details() {
		if br, ok := detail.(*errdetails.BadRequest); ok {
			assert.Len(t, br.GetFieldViolations(), 1)
			violation := br.GetFieldViolations()[0]
			assert.Equal(t, "URL", violation.GetField())
			assert.Contains(t, violation.GetDescription(), "URL")
			return
		}
	}
	assert.Fail(t, "validation error not found")
}

func TestCreateFeed(t *testing.T) {
	t.Parallel()
	ts := newServer(t)
	ctx := context.Background()
	resp, err := ts.Server.CreateFeed(ctx, &pb.CreateFeedRequest{
		Feed: &pb.CreateFeedRequest_Feed{
			Name: "a",
			Url:  "https://example.com",
		},
		User: &pb.User{Id: seed.UserID},
	})
	require.NoError(t, err)
	assert.NotEmpty(t, resp.GetId())
}

func TestUpdateFeed(t *testing.T) {
	t.Parallel()
	ts := newServer(t)
	ctx := context.Background()
	fakes, err := seed.FeedAndArticles(ts.DB, 1)
	require.NoError(t, err)

	attempt := pb.UpdateFeedRequest_Feed{
		Id:   fakes.Feed.ID,
		Name: "new name",
		Url:  "https://example.com/not-allowed",
	}

	_, err = ts.Server.UpdateFeed(ctx, &pb.UpdateFeedRequest{
		Feed: &attempt,
	})
	require.NoError(t, err)

	actual, err := ts.Service.GetFeed(ctx, fakes.Feed.ID)
	require.NoError(t, err)
	assert.Equal(t, attempt.GetName(), actual.Name)
	assert.Equal(t, fakes.Feed.URL, actual.URL, "URL is immutable")
}

func TestSaveArticleValidateError(t *testing.T) {
	t.Parallel()
	// TODO: table test to assert all validation errors
	ctx := context.Background()
	ts := newServer(t)
	_, err := ts.Server.SaveArticle(ctx, &pb.SaveArticleRequest{
		Article: &pb.Article{},
	})
	require.Error(t, err)

	s := status.Convert(err)
	for _, detail := range s.Details() {
		if br, ok := detail.(*errdetails.BadRequest); ok {
			assert.Len(t, br.GetFieldViolations(), 2)
			violations := br.GetFieldViolations()
			assert.Equal(t, "FeedID", violations[0].GetField())
			assert.Equal(t, "URL", violations[1].GetField())
			return
		}
	}
	assert.Fail(t, "validation error not found")
}

func TestSaveArticleFeed(t *testing.T) {
	t.Parallel()
	ctx := context.Background()
	ts := newServer(t)
	resp, err := ts.Server.SaveArticle(ctx, &pb.SaveArticleRequest{
		Article: &pb.Article{
			FeedId: "0e597e90-ece5-463e-8608-ff687bf286da",
			Url:    "https://example.com",
		},
	})
	require.NoError(t, err)
	assert.NotEmpty(t, resp.GetId())
}

func TestListArticles(t *testing.T) {
	t.Parallel()

	ts := newServer(t)
	fakes, err := seed.FeedAndArticles(ts.DB, 1)
	require.NoError(t, err)

	ctx := context.Background()
	resp, err := ts.Server.ListArticles(ctx, &pb.ListArticlesRequest{
		FeedId: fakes.Feed.ID,
	})
	articles := resp.GetArticles()
	require.NoError(t, err)
	assert.Len(t, articles, 1)
	c := converters.New()
	helpers.Diff(t, *fakes.Articles[0], *c.ProtoToServiceArticle(articles[0]))
}

func TestListArticles_Pagination(t *testing.T) {
	// note: this test only checks for pagination, not the actual content
	t.Parallel()

	ts := newServer(t)
	fakes, err := seed.FeedAndArticles(ts.DB, 2)
	require.NoError(t, err)
	ctx := context.Background()

	// page 1
	resp, err := ts.Server.ListArticles(ctx, &pb.ListArticlesRequest{
		FeedId: fakes.Feed.ID,
		Options: &pb.ListOptions{
			Limit: 1,
			// Test default cursor
		},
	})
	cursor := resp.GetCursor().GetNext()
	require.NoError(t, err)
	assert.Len(t, resp.GetArticles(), 1)
	assert.NotEmpty(t, cursor)

	// page 2
	resp, err = ts.Server.ListArticles(ctx, &pb.ListArticlesRequest{
		FeedId: fakes.Feed.ID,
		Options: &pb.ListOptions{
			Limit: 1,
			Cursor: &pb.Cursor{
				Next: resp.GetCursor().GetNext(),
			},
		},
	})
	require.NoError(t, err)
	assert.Len(t, resp.GetArticles(), 1)
}

func TestListFeeds(t *testing.T) {
	t.Parallel()
	ts := newServer(t)
	ctx := context.Background()

	fakes, err := seed.FeedAndArticles(ts.DB, 1)
	require.NoError(t, err)

	resp, err := ts.Server.ListFeeds(ctx, &pb.ListFeedsRequest{})
	require.NoError(t, err)
	assert.Len(t, resp.GetFeeds(), 1)
	assert.Equal(t, fakes.Feed.ID, resp.GetFeeds()[0].GetId())
}

func TestListUserFeeds(t *testing.T) {
	t.Parallel()
	ts := newServer(t)
	ctx := context.Background()

	fakes, err := seed.FeedAndArticles(ts.DB, 1)
	require.NoError(t, err)

	resp, err := ts.Server.ListUserFeeds(ctx, &pb.ListUserFeedsRequest{
		User: &pb.User{Id: seed.UserID},
	})
	require.NoError(t, err)
	assert.Len(t, resp.GetFeeds(), 1)
	assert.Equal(t, fakes.Feed.ID, resp.GetFeeds()[0].GetFeedId())
}

func TestGetFeed(t *testing.T) {
	t.Parallel()
	ts := newServer(t)
	fakes, err := seed.FeedAndArticles(ts.DB, 1)
	require.NoError(t, err)

	ctx := context.Background()
	resp, err := ts.Server.GetFeed(ctx, &pb.GetFeedRequest{
		Id: fakes.Feed.ID,
	})
	require.NoError(t, err)
	c := converters.New()
	helpers.Diff(t, *fakes.Feed, *c.ProtoToServiceFeed(resp.GetFeed()), "IsActive")
}

func TestGetUserFeed(t *testing.T) {
	t.Parallel()
	ts := newServer(t)
	articleCount := 1
	fakes, err := seed.FeedAndArticles(ts.DB, articleCount)
	require.NoError(t, err)

	expected := &pb.UserFeed{
		FeedId:        fakes.Feed.ID,
		Url:           fakes.Feed.URL,
		Name:          fakes.Feed.Name,
		UnreadCount:   int32(articleCount),
		CreatedAt:     timestamppb.New(fakes.UserFeed.CreatedAt),
		ViewedAt:      timestamppb.New(fakes.UserFeed.ViewedAt),
		UnreadStartAt: timestamppb.New(fakes.UserFeed.UnreadStartAt),
	}

	ctx := context.Background()
	resp, err := ts.Server.GetUserFeed(ctx, &pb.GetUserFeedRequest{
		UserId: seed.UserID,
		FeedId: fakes.Feed.ID,
	})
	require.NoError(t, err)

	assert.Empty(t, cmp.Diff(
		expected,
		resp.GetFeed(),
		cmpopts.IgnoreUnexported(pb.UserFeed{}, timestamppb.Timestamp{}),
	))
}

func TestGetArticle(t *testing.T) {
	t.Parallel()
	ts := newServer(t)
	fakes, err := seed.FeedAndArticles(ts.DB, 1)
	require.NoError(t, err)

	ctx := context.Background()
	resp, err := ts.Server.GetArticle(ctx, &pb.GetArticleRequest{
		Id: fakes.Articles[0].ID,
	})
	require.NoError(t, err)
	c := converters.New()
	helpers.Diff(t, *fakes.Articles[0], *c.ProtoToServiceArticle(resp.GetArticle()))
}

func TestFeedTasks(t *testing.T) {
	t.Parallel()
	ts := newServer(t)
	ctx := context.Background()
	_, err := ts.Server.FeedTask(ctx, &pb.FeedTaskRequest{}) //nolint: staticcheck
	require.Error(t, err)
}

func TestUpdateStats(t *testing.T) {
	t.Parallel()

	feedID := "0e597e90-ece5-463e-8608-ff687bf286da"

	s := newMockServer(t)
	s.svc.EXPECT().UpdateFeedArticleCount(mock.Anything, feedID).Return(nil)

	resp, err := s.server.UpdateStats(context.Background(), &pb.UpdateStatsRequest{
		Stat:   pb.UpdateStatsRequest_STAT_FEED_ARTICLE_COUNT,
		FeedId: feedID,
	})
	require.NoError(t, err)
	assert.NotNil(t, resp)
}

func TestReady(t *testing.T) {
	t.Parallel()
	ts := newServer(t)
	ctx := context.Background()
	_, err := ts.Server.Ready(ctx, &pb.ReadyRequest{})
	require.NoError(t, err)
}

func TestMarkArticleAsRead(t *testing.T) {
	t.Parallel()
	s := newMockServer(t)

	feedID := "2e597e90-ece5-463e-8608-ff687bf286da"
	articleID := "3e597e90-ece5-463e-8608-ff687bf286da"

	s.svc.EXPECT().
		SaveUserArticle(mock.Anything, &svcModel.UserArticle{
			UserID:    seed.UserID,
			ArticleID: articleID,
		}).
		Return(nil)
	s.svc.EXPECT().
		GetArticle(mock.Anything, articleID).
		Return(&svcModel.Article{FeedID: feedID}, nil)
	s.svc.EXPECT().
		UpdateFeedArticleCount(mock.Anything, feedID).
		Return(nil)

	_, err := s.server.MarkArticleAsRead(context.Background(), &pb.MarkArticleAsReadRequest{
		User:      &pb.User{Id: seed.UserID},
		ArticleId: articleID,
	})
	require.NoError(t, err)
}

func TestCreateFeedVerification(t *testing.T) {
	t.Parallel()
	s := newMockServer(t)

	data := fixtures.NewFeedVerification()
	expected := s.converters.ServiceToProtoFeedVerification(data)

	s.svc.EXPECT().
		CreateFeedVerification(mock.Anything, data).
		Return(data, nil)

	resp, err := s.server.CreateFeedVerification(context.Background(), &pb.CreateFeedVerificationRequest{
		Verification: expected,
	})
	require.NoError(t, err)
	helpers.DiffProto(t, expected, resp.GetVerification())
}

func TestCreateFetchHistory(t *testing.T) {
	t.Parallel()
	s := newMockServer(t)

	data := fixtures.NewFetchHistory()
	expected := s.converters.ServiceToProtoFetchHistory(data)

	s.svc.EXPECT().
		CreateFetchHistory(mock.Anything, data).
		Return(data, nil)

	resp, err := s.server.CreateFetchHistory(context.Background(), &pb.CreateFetchHistoryRequest{
		History: expected,
	})
	require.NoError(t, err)
	helpers.DiffProto(t, expected, resp.GetHistory())
}
