package server

import (
	"context"

	"github.com/ericbutera/amalgam/internal/copygen"
	models "github.com/ericbutera/amalgam/internal/service/models"
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
		pbFeed := pb.Feed{}
		copygen.ServiceToProtoFeed(&pbFeed, &feed)
		pbFeeds = append(pbFeeds, &pbFeed)
	}
	return &pb.ListFeedsResponse{Feeds: pbFeeds}, nil
}

func (s *Server) CreateFeed(ctx context.Context, in *pb.CreateFeedRequest) (*pb.CreateFeedResponse, error) {
	feed := &models.Feed{}
	copygen.ProtoCreateFeedToService(feed, in.Feed)
	if err := s.service.CreateFeed(ctx, feed); err != nil {
		return nil, status.Errorf(codes.Internal, "failed to create feed: %v", err)
	}
	return &pb.CreateFeedResponse{
		Id: feed.ID,
	}, nil
}

func (s *Server) UpdateFeed(ctx context.Context, in *pb.UpdateFeedRequest) (*pb.UpdateFeedResponse, error) {
	feed := &models.Feed{}
	copygen.ProtoUpdateFeedToService(feed, in.Feed)
	if err := s.service.UpdateFeed(ctx, feed.ID, feed); err != nil {
		return nil, status.Errorf(codes.Internal, "failed to create feed: %v", err)
	}
	return &pb.UpdateFeedResponse{}, nil
}

func (s *Server) GetFeed(ctx context.Context, in *pb.GetFeedRequest) (*pb.GetFeedResponse, error) {
	feed, err := s.service.GetFeed(ctx, in.Id)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to fetch feed: %v", err)
	}
	pbFeed := &pb.Feed{}
	copygen.ServiceToProtoFeed(pbFeed, feed)
	return &pb.GetFeedResponse{
		Feed: pbFeed,
	}, nil
}

func (s *Server) ListArticles(ctx context.Context, in *pb.ListArticlesRequest) (*pb.ListArticlesResponse, error) {
	// TODO: support filters (sorting, pagination)
	// TODO: convert ByFeedId to a filter
	articles, err := s.service.GetArticlesByFeed(ctx, in.FeedId)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to fetch articles: %v", err)
	}
	pbArticles := []*pb.Article{}
	for _, article := range articles {
		pbArticle := pb.Article{}
		copygen.ServiceToProtoArticle(&pbArticle, &article)
		pbArticles = append(pbArticles, &pbArticle)
	}
	return &pb.ListArticlesResponse{Articles: pbArticles}, nil
}

func (s *Server) GetArticle(ctx context.Context, in *pb.GetArticleRequest) (*pb.GetArticleResponse, error) {
	article, err := s.service.GetArticle(ctx, in.Id)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to fetch article: %v", err)
	}
	pbArticle := pb.Article{}
	copygen.ServiceToProtoArticle(&pbArticle, article)
	return &pb.GetArticleResponse{
		Article: &pbArticle,
	}, nil
}
