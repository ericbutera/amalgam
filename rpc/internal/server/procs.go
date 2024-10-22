package server

import (
	"context"
	"fmt"

	pb "github.com/ericbutera/amalgam/pkg/rpc/proto"
)

func (s *Server) ListFeeds(ctx context.Context, in *pb.Empty) (*pb.ListFeedsResponse, error) {
	// TODO: use database
	feeds := []*pb.Feed{
		{Id: 1, Url: "http://example.com", Name: "Example Feed"},
		{Id: 2, Url: "http://another.com", Name: "Another Feed"},
	}
	return &pb.ListFeedsResponse{Feeds: feeds}, nil
}

func (s *Server) CreateFeed(ctx context.Context, in *pb.CreateFeedRequest) (*pb.Feed, error) {
	// TODO: use database
	feed := &pb.Feed{
		Id:   3, // Generate a new ID
		Url:  in.GetUrl(),
		Name: in.GetName(),
	}
	fmt.Printf("Created feed: %v\n", feed)
	return feed, nil
}

// Implement the UpdateFeed method
func (s *Server) UpdateFeed(ctx context.Context, in *pb.UpdateFeedRequest) (*pb.Feed, error) {
	// TODO: use database
	feed := &pb.Feed{
		Id:   in.GetId(),
		Url:  in.GetUrl(),
		Name: in.GetName(),
	}
	fmt.Printf("Updated feed: %v\n", feed)
	return feed, nil
}

func (s *Server) ListArticlesByFeed(ctx context.Context, in *pb.ListArticlesByFeedRequest) (*pb.ListArticlesByFeedResponse, error) {
	// TODO: use database
	articles := []*pb.Article{
		{Id: 1, Title: "Article 1", Content: "Content of article 1", FeedId: in.GetFeedId()},
		{Id: 2, Title: "Article 2", Content: "Content of article 2", FeedId: in.GetFeedId()},
	}
	return &pb.ListArticlesByFeedResponse{Articles: articles}, nil
}

func (s *Server) GetArticleById(ctx context.Context, in *pb.GetArticleByIdRequest) (*pb.Article, error) {
	// TODO: use database
	article := &pb.Article{
		Id:      in.GetId(),
		Title:   "Example Article",
		Content: "Example content",
		FeedId:  1,
	}
	return article, nil
}
