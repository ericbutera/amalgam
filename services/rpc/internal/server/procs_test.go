package server_test

import (
	"context"
	"testing"

	"github.com/ericbutera/amalgam/internal/converters"
	"github.com/ericbutera/amalgam/internal/service"
	"github.com/ericbutera/amalgam/internal/test"
	"github.com/ericbutera/amalgam/internal/test/seed"
	"github.com/ericbutera/amalgam/pkg/config/env"
	pb "github.com/ericbutera/amalgam/pkg/feeds/v1"
	helpers "github.com/ericbutera/amalgam/pkg/test"
	"github.com/ericbutera/amalgam/services/rpc/internal/config"
	"github.com/ericbutera/amalgam/services/rpc/internal/server"
	"github.com/samber/lo"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc"
	"google.golang.org/grpc/status"
	"gorm.io/gorm"
)

type TestServer struct {
	Server  *server.Server
	Service service.Service
	DB      *gorm.DB
}

func newServer(t *testing.T) *TestServer {
	config := lo.Must(env.New[config.Config]())
	db := test.NewDB(t) // db := lo.Must(db.NewSqlite("file::memory:", db.WithAutoMigrate()))
	svc := service.NewGorm(db)
	server := lo.Must(server.New(
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

func TestCreateFeedValidateError(t *testing.T) {
	t.Parallel()

	ts := newServer(t)
	// TODO: table test to assert all validation errors
	ctx := context.Background()
	_, err := ts.Server.CreateFeed(ctx, &pb.CreateFeedRequest{
		Feed: &pb.CreateFeedRequest_Feed{
			Name: "a",
			Url:  "b",
		},
	})
	require.Error(t, err)
	s := status.Convert(err)
	for _, detail := range s.Details() {
		if v, ok := detail.(*pb.ValidationErrors); ok {
			errors := v.GetErrors()
			assert.Len(t, errors, 1)
			assert.Contains(t, errors[0].GetField(), "URL")
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
	})
	require.NoError(t, err)
	assert.NotEmpty(t, resp.GetId())
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
		if v, ok := detail.(*pb.ValidationErrors); ok {
			assert.Len(t, v.GetErrors(), 2)
			assert.Equal(t, "The FeedID field is required.", v.GetErrors()[0].GetMessage())
			assert.Equal(t, "The URL is required.", v.GetErrors()[1].GetMessage())
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

/*
TODO:
func TestListArticles_Pagination(t *testing.T) {
	t.Parallel()

	ts := newServer(t)
	fakes, err := seed.Feed(ts.DB, 11)
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
*/
