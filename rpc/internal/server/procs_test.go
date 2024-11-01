package server

import (
	"context"
	"testing"

	"github.com/ericbutera/amalgam/internal/db"
	"github.com/ericbutera/amalgam/internal/service"
	pb "github.com/ericbutera/amalgam/pkg/feeds/v1"
	"github.com/samber/lo"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc"
	"google.golang.org/grpc/status"
)

func newServer() *Server {
	db := lo.Must(db.NewSqlite("file::memory:", db.WithAutoMigrate()))
	server := Server{
		service: service.NewGorm(db),
	}
	s := grpc.NewServer()
	pb.RegisterFeedServiceServer(s, &server)
	return &server
}

func TestCreateFeedValidateError(t *testing.T) {
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
			assert.Len(t, v.Errors, 1)
			assert.Equal(t, "The URL must be valid.", v.Errors[0].Message)
			return
		}
	}
	assert.Fail(t, "validation error not found")
}

func TestCreateFeed(t *testing.T) {
	ctx := context.Background()
	resp, err := newServer().CreateFeed(ctx, &pb.CreateFeedRequest{
		Feed: &pb.CreateFeedRequest_Feed{
			Name: "a",
			Url:  "https://example.com",
		},
	})
	assert.NoError(t, err)
	assert.NotEmpty(t, resp.Id)
}

func TestSaveArticleValidateError(t *testing.T) {
	// TODO: table test to assert all validation errors
	ctx := context.Background()
	_, err := newServer().SaveArticle(ctx, &pb.SaveArticleRequest{
		Article: &pb.Article{},
	})
	require.Error(t, err)
	s := status.Convert(err)
	for _, detail := range s.Details() {
		if v, ok := detail.(*pb.ValidationErrors); ok {
			assert.Len(t, v.Errors, 2)
			assert.Equal(t, "The FeedID field is required.", v.Errors[0].Message)
			assert.Equal(t, "The URL is required.", v.Errors[1].Message)
			return
		}
	}
	assert.Fail(t, "validation error not found")
}

func TestSaveArticleFeed(t *testing.T) {
	ctx := context.Background()
	resp, err := newServer().SaveArticle(ctx, &pb.SaveArticleRequest{
		Article: &pb.Article{
			FeedId: "0e597e90-ece5-463e-8608-ff687bf286da",
			Url:    "https://example.com",
		},
	})
	assert.NoError(t, err)
	assert.NotEmpty(t, resp.Id)
}
