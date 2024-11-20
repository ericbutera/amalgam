package server_test

import (
	"context"
	"testing"

	"github.com/ericbutera/amalgam/internal/db"
	"github.com/ericbutera/amalgam/internal/service"
	"github.com/ericbutera/amalgam/pkg/config/env"
	pb "github.com/ericbutera/amalgam/pkg/feeds/v1"
	"github.com/ericbutera/amalgam/services/rpc/internal/config"
	"github.com/ericbutera/amalgam/services/rpc/internal/server"
	"github.com/ericbutera/amalgam/services/rpc/internal/tasks"
	"github.com/samber/lo"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc"
	"google.golang.org/grpc/status"
)

func newServer(opts ...server.Option) *server.Server {
	db := lo.Must(db.NewSqlite("file::memory:", db.WithAutoMigrate()))
	config := lo.Must(env.New[config.Config]())

	defs := []server.Option{
		server.WithTasks(nil),
		server.WithService(service.NewGorm(db)),
		server.WithConfig(config),
	}
	opts = append(defs, opts...)

	server := lo.Must(server.New(
		opts...,
	))
	s := grpc.NewServer()
	pb.RegisterFeedServiceServer(s, server)
	return server
}

func TestCreateFeedValidateError(t *testing.T) {
	t.Parallel()
	// TODO: table test to assert all validation errors
	ctx := context.Background()
	_, err := newServer().CreateFeed(ctx, &pb.CreateFeedRequest{
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
	ctx := context.Background()
	resp, err := newServer().CreateFeed(ctx, &pb.CreateFeedRequest{
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
	_, err := newServer().SaveArticle(ctx, &pb.SaveArticleRequest{
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
	resp, err := newServer().SaveArticle(ctx, &pb.SaveArticleRequest{
		Article: &pb.Article{
			FeedId: "0e597e90-ece5-463e-8608-ff687bf286da",
			Url:    "https://example.com",
		},
	})
	require.NoError(t, err)
	assert.NotEmpty(t, resp.GetId())
}

func TestFeedTask(t *testing.T) {
	t.Parallel()
	ctx := context.Background()
	expectedID := "super-id"

	mockTasks := new(tasks.MockTasks)
	mockTasks.EXPECT().
		Workflow(mock.Anything, tasks.TaskGenerateFeeds).
		Return(&tasks.TaskResult{ID: expectedID}, nil)

	resp, err := newServer(server.WithTasks(mockTasks)).
		FeedTask(ctx, &pb.FeedTaskRequest{
			Task: pb.FeedTaskRequest_TASK_GENERATE_FEEDS,
		})
	require.NoError(t, err)
	assert.Equal(t, expectedID, resp.GetId())
}
