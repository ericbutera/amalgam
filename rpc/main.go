package main

import (
	"context"
	"fmt"
	"log"
	"net"

	pb "github.com/ericbutera/amalgam/pkg/rpc/proto"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

const Port = "50055"

type server struct {
	pb.UnimplementedFeedServiceServer
}

// Implement the ListFeeds method
func (s *server) ListFeeds(ctx context.Context, in *pb.Empty) (*pb.ListFeedsResponse, error) {
	// TODO call database
	feeds := []*pb.Feed{
		{Id: 1, Url: "http://example.com", Name: "Example Feed"},
		{Id: 2, Url: "http://another.com", Name: "Another Feed"},
	}
	return &pb.ListFeedsResponse{Feeds: feeds}, nil
}

// Implement the CreateFeed method
func (s *server) CreateFeed(ctx context.Context, in *pb.CreateFeedRequest) (*pb.Feed, error) {
	feed := &pb.Feed{
		Id:   3, // Generate a new ID
		Url:  in.GetUrl(),
		Name: in.GetName(),
	}
	fmt.Printf("Created feed: %v\n", feed)
	return feed, nil
}

// Implement the UpdateFeed method
func (s *server) UpdateFeed(ctx context.Context, in *pb.UpdateFeedRequest) (*pb.Feed, error) {
	feed := &pb.Feed{
		Id:   in.GetId(),
		Url:  in.GetUrl(),
		Name: in.GetName(),
	}
	fmt.Printf("Updated feed: %v\n", feed)
	return feed, nil
}

// Implement the ListArticlesByFeed method
func (s *server) ListArticlesByFeed(ctx context.Context, in *pb.ListArticlesByFeedRequest) (*pb.ListArticlesByFeedResponse, error) {
	articles := []*pb.Article{
		{Id: 1, Title: "Article 1", Content: "Content of article 1", FeedId: in.GetFeedId()},
		{Id: 2, Title: "Article 2", Content: "Content of article 2", FeedId: in.GetFeedId()},
	}
	return &pb.ListArticlesByFeedResponse{Articles: articles}, nil
}

// Implement the GetArticleById method
func (s *server) GetArticleById(ctx context.Context, in *pb.GetArticleByIdRequest) (*pb.Article, error) {
	article := &pb.Article{
		Id:      in.GetId(),
		Title:   "Example Article",
		Content: "Example content",
		FeedId:  1,
	}
	return article, nil
}

func main() {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%s", Port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterFeedServiceServer(s, &server{})
	reflection.Register(s)
	fmt.Printf("gRPC server is running on port %s...\n", Port)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
