package server

import (
	"context"
	"fmt"

	"github.com/ericbutera/amalgam/internal/db/models"
	pb "github.com/ericbutera/amalgam/pkg/feeds/v1"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) ListFeeds(ctx context.Context, in *pb.ListFeedsRequest) (*pb.ListFeedsResponse, error) {
	feeds, err := s.service.Feeds(ctx)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to fetch feeds: %v", err)
	}
	pbFeeds := []*pb.Feed{}
	for _, feed := range feeds {
		pbFeeds = append(pbFeeds, &pb.Feed{
			Id:   feed.ID,
			Url:  feed.Url,
			Name: feed.Name,
		})
	}
	return &pb.ListFeedsResponse{Feeds: pbFeeds}, nil
}

func (s *Server) CreateFeed(ctx context.Context, in *pb.CreateFeedRequest) (*pb.CreateFeedResponse, error) {
	feed := &models.Feed{
		Url:  in.Url,
		Name: in.Name,
	}
	if err := s.service.CreateFeed(ctx, feed); err != nil {
		return nil, status.Errorf(codes.Internal, "failed to create feed: %v", err)
	}
	return &pb.CreateFeedResponse{
		Id: feed.ID,
	}, nil
}

func (s *Server) UpdateFeed(ctx context.Context, in *pb.UpdateFeedRequest) (*pb.UpdateFeedResponse, error) {
	feed := &models.Feed{
		Name: in.Name,
	}
	if err := s.service.UpdateFeed(ctx, in.Id, feed); err != nil {
		return nil, status.Errorf(codes.Internal, "failed to create feed: %v", err)
	}
	return &pb.UpdateFeedResponse{}, nil
}

func (s *Server) ListArticles(ctx context.Context, in *pb.ListArticlesRequest) (*pb.ListArticlesResponse, error) {
	// articles := []*pb.Article{
	// 	{Id: 1, Title: "Article 1", Content: "Content of article 1", FeedId: in.GetFeedId()},
	// 	{Id: 2, Title: "Article 2", Content: "Content of article 2", FeedId: in.GetFeedId()},
	// }
	// return &pb.ListArticlesResponse{Articles: articles}, nil
	return nil, fmt.Errorf("not implemented")
}

func (s *Server) GetArticle(ctx context.Context, in *pb.GetArticleRequest) (*pb.GetArticleResponse, error) {
	// article := &pb.Article{
	// 	Id:      in.GetId(),
	// 	Title:   "Example Article",
	// 	Content: "Example content",
	// 	FeedId:  1,
	// }
	// return article, nil
	return nil, fmt.Errorf("not implemented")
}
