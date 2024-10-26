package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.55

import (
	"context"
	"log/slog"

	"github.com/ericbutera/amalgam/graph/graph/model"
	pb "github.com/ericbutera/amalgam/pkg/feeds/v1"
	"github.com/samber/lo"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// AddFeed is the resolver for the addFeed field.
func (r *mutationResolver) AddFeed(ctx context.Context, url string, name string) (*model.AddResponse, error) {
	// TODO: grpc middleware to log errors
	resp, err := r.rpcClient.CreateFeed(ctx, &pb.CreateFeedRequest{
		Url:  url,
		Name: name,
	})
	if err != nil {
		return nil, status.Error(codes.Internal, "failed to create feed")
	}
	// TODO: converter.ServiceToGraphFeed
	return &model.AddResponse{
		ID: resp.Id,
	}, nil
}

// UpdateFeed is the resolver for the updateFeed field.
func (r *mutationResolver) UpdateFeed(ctx context.Context, id string, url *string, name *string) (*model.UpdateResponse, error) {
	_, err := r.rpcClient.UpdateFeed(ctx, &pb.UpdateFeedRequest{
		Id:   id,
		Url:  lo.FromPtr(url),
		Name: lo.FromPtr(name),
	})
	if err != nil {
		return nil, status.Error(codes.Internal, "failed to update feed")
	}
	// TODO: converter.ServiceToGraphFeed
	// TODO: revisit returning id (rpc returns empty)
	return &model.UpdateResponse{
		ID: id,
	}, nil
}

// Feeds is the resolver for the feeds field.
func (r *queryResolver) Feeds(ctx context.Context) ([]*model.Feed, error) {
	var feeds []*model.Feed
	res, err := r.rpcClient.ListFeeds(ctx, &pb.ListFeedsRequest{})
	if err != nil {
		return nil, status.Error(codes.Internal, "failed to list feeds")
	}
	// TODO: converter.ServiceToGraph
	for _, feed := range res.Feeds {
		feeds = append(feeds, &model.Feed{
			ID:   feed.Id,
			URL:  feed.Url,
			Name: feed.Name,
		})
	}
	return feeds, nil
}

// Feed is the resolver for the feed field.
func (r *queryResolver) Feed(ctx context.Context, id string) (*model.Feed, error) {
	resp, err := r.rpcClient.GetFeed(ctx, &pb.GetFeedRequest{Id: id})
	if err != nil {
		slog.ErrorContext(ctx, "failed to get feed", "error", err) // TODO: use middleware
		return nil, status.Error(codes.Internal, "failed to get feed")
	}
	// TODO: converter.ServiceToGraph
	return &model.Feed{
		ID:   resp.Feed.Id,
		URL:  resp.Feed.Url,
		Name: resp.Feed.Name,
	}, nil
}

// Articles is the resolver for the articles field.
func (r *queryResolver) Articles(ctx context.Context, feedID string) ([]*model.Article, error) {
	resp, err := r.rpcClient.ListArticles(ctx, &pb.ListArticlesRequest{FeedId: feedID})
	if err != nil {
		return nil, status.Error(codes.Internal, "failed to list articles")
	}

	var articles []*model.Article
	for _, article := range resp.Articles {
		// TODO: converter.ServiceToGraph
		articles = append(articles, &model.Article{
			// TODO: limit fields on listing (return preview instead of content)
			ID:          article.Id,
			Title:       article.Title,
			Content:     article.Content,
			FeedID:      article.FeedId,
			URL:         article.Url,
			Preview:     article.Preview,
			GUID:        lo.ToPtr(article.Guid),
			ImageURL:    lo.ToPtr(article.ImageUrl),
			AuthorName:  lo.ToPtr(article.AuthorName),
			AuthorEmail: lo.ToPtr(article.AuthorEmail),
		})
	}
	return articles, nil
}

// Article is the resolver for the article field.
func (r *queryResolver) Article(ctx context.Context, id string) (*model.Article, error) {
	resp, err := r.rpcClient.GetArticle(ctx, &pb.GetArticleRequest{Id: id})
	if err != nil {
		return nil, status.Error(codes.Internal, "failed to get article")
	}
	article := resp.Article
	// TODO: converter.ServiceToGraph
	return &model.Article{
		ID:          article.Id,
		Title:       article.Title,
		Content:     article.Content,
		FeedID:      article.FeedId,
		URL:         article.Url,
		Preview:     article.Preview,
		GUID:        lo.ToPtr(article.Guid),
		ImageURL:    lo.ToPtr(article.ImageUrl),
		AuthorName:  lo.ToPtr(article.AuthorName),
		AuthorEmail: lo.ToPtr(article.AuthorEmail),
	}, nil
}

// Mutation returns MutationResolver implementation.
func (r *Resolver) Mutation() MutationResolver { return &mutationResolver{r} }

// Query returns QueryResolver implementation.
func (r *Resolver) Query() QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
