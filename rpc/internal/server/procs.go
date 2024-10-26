package server

import (
	"context"

	"github.com/ericbutera/amalgam/internal/service"
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
		// TODO: converter.ServiceToProto
		pbFeeds = append(pbFeeds, &pb.Feed{
			Id:   feed.ID,
			Url:  feed.Url,
			Name: feed.Name,
		})
	}
	return &pb.ListFeedsResponse{Feeds: pbFeeds}, nil
}

func (s *Server) CreateFeed(ctx context.Context, in *pb.CreateFeedRequest) (*pb.CreateFeedResponse, error) {
	// TODO: converter.ServiceToProto
	feed := &service.Feed{
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
	feed := &service.Feed{
		Name: in.Name,
	}
	if err := s.service.UpdateFeed(ctx, in.Id, feed); err != nil {
		return nil, status.Errorf(codes.Internal, "failed to create feed: %v", err)
	}
	return &pb.UpdateFeedResponse{}, nil
}

func (s *Server) GetFeed(ctx context.Context, in *pb.GetFeedRequest) (*pb.GetFeedResponse, error) {
	feed, err := s.service.GetFeed(ctx, in.Id)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to fetch feed: %v", err)
	}
	return &pb.GetFeedResponse{
		// TODO: converter.ServiceToProto
		Feed: &pb.Feed{
			Id:   feed.ID,
			Url:  feed.Url,
			Name: feed.Name,
		},
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
		// TODO: converter.ServiceToProto
		// TODO: limit list fields
		pbArticles = append(pbArticles, &pb.Article{
			Id:          article.ID,
			Title:       article.Title,
			Content:     article.Content,
			FeedId:      article.FeedID,
			Preview:     article.Preview,
			Url:         article.Url,
			ImageUrl:    article.ImageUrl,
			Guid:        article.Guid,
			AuthorName:  article.AuthorName,
			AuthorEmail: article.AuthorEmail,
		})
	}
	return &pb.ListArticlesResponse{Articles: pbArticles}, nil
}

func (s *Server) GetArticle(ctx context.Context, in *pb.GetArticleRequest) (*pb.GetArticleResponse, error) {
	article, err := s.service.GetArticle(ctx, in.Id)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to fetch article: %v", err)
	}
	return &pb.GetArticleResponse{
		// TODO converter.ServiceToProto
		Article: &pb.Article{
			Id:          article.ID,
			Title:       article.Title,
			Content:     article.Content,
			FeedId:      article.FeedID,
			Url:         article.Url,
			ImageUrl:    article.ImageUrl,
			Preview:     article.Preview,
			Guid:        article.Guid,
			AuthorName:  article.AuthorName,
			AuthorEmail: article.AuthorEmail,
		},
	}, nil
}
