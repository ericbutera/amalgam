package server

import (
	"context"
	"fmt"

	pb "github.com/ericbutera/amalgam/pkg/rpc/proto"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) ListFeeds(ctx context.Context, in *pb.Empty) (*pb.ListFeedsResponse, error) {
	feeds, err := s.service.Feeds(ctx)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to fetch feeds: %v", err)
	}
	pbFeeds := []*pb.Feed{}
	for _, feed := range feeds {
		pbFeeds = append(pbFeeds, &pb.Feed{
			Id:   uint32(feed.ID),
			Url:  feed.Url,
			Name: feed.Name,
		})
	}
	return &pb.ListFeedsResponse{Feeds: pbFeeds}, nil
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
