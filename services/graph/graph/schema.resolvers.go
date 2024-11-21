package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.55

import (
	"context"
	"log/slog"

	"github.com/ericbutera/amalgam/internal/converters"
	"github.com/ericbutera/amalgam/internal/tasks"
	pb "github.com/ericbutera/amalgam/pkg/feeds/v1"
	"github.com/ericbutera/amalgam/services/graph/graph/model"
	errHelper "github.com/ericbutera/amalgam/services/graph/internal/errors"
	"github.com/samber/lo"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// AddFeed is the resolver for the addFeed field.
func (r *mutationResolver) AddFeed(ctx context.Context, url string, name string) (*model.AddResponse, error) {
	// TODO: middleware to log errors
	resp, err := r.rpcClient.CreateFeed(ctx, &pb.CreateFeedRequest{
		Feed: &pb.CreateFeedRequest_Feed{
			Url:  url,
			Name: name,
		},
	})
	if err != nil {
		return nil, errHelper.HandleGrpcErrors(ctx, err, "failed to create feed")
	}
	return &model.AddResponse{
		ID: resp.Id,
	}, nil
}

// UpdateFeed is the resolver for the updateFeed field.
func (r *mutationResolver) UpdateFeed(ctx context.Context, id string, url *string, name *string) (*model.UpdateResponse, error) {
	_, err := r.rpcClient.UpdateFeed(ctx, &pb.UpdateFeedRequest{
		Feed: &pb.UpdateFeedRequest_Feed{
			Id:   id,
			Url:  lo.FromPtr(url),
			Name: lo.FromPtr(name),
		},
	})
	if err != nil {
		return nil, errHelper.HandleGrpcErrors(ctx, err, "failed to update feed")
	}
	// TODO: converter.ServiceToGraphFeed
	// TODO: revisit returning id (rpc returns empty)
	return &model.UpdateResponse{
		ID: id,
	}, nil
}

// GenerateFeeds is the resolver for the generateFeeds field.
func (r *mutationResolver) GenerateFeeds(ctx context.Context) (*model.GenerateFeedsResponse, error) {
	// TODO: support task type & args
	// func pbTaskToTaskType(task pb.FeedTaskRequest_Task) (tasks.TaskType, error) {
	// 	if task == pb.FeedTaskRequest_TASK_GENERATE_FEEDS {
	// 		return tasks.TaskGenerateFeeds, nil
	// 	}
	// 	return tasks.TaskUnspecified, ErrInvalidTaskType
	// }
	// taskType, err := pbTaskToTaskType(in.GetTask())
	// if err != nil {
	// 	return nil, status.Errorf(codes.InvalidArgument, "invalid task type")
	// }

	task, err := r.tasks.Workflow(ctx, tasks.TaskGenerateFeeds)
	if err != nil {
		slog.ErrorContext(ctx, "failed to start feed task", "error", err)
		return nil, status.Errorf(codes.Internal, "failed to start feed task")
	}
	return &model.GenerateFeedsResponse{
		ID: task.ID,
	}, nil
}

// FetchFeeds is the resolver for the fetchFeeds field.
func (r *mutationResolver) FetchFeeds(ctx context.Context) (*model.FetchFeedsResponse, error) {
	task, err := r.tasks.Workflow(ctx, tasks.TaskFetchFeeds)
	if err != nil {
		slog.ErrorContext(ctx, "failed to start feed task", "error", err)
		return nil, status.Errorf(codes.Internal, "failed to start feed task")
	}
	return &model.FetchFeedsResponse{
		ID: task.ID,
	}, nil
}

// Feeds is the resolver for the feeds field.
func (r *queryResolver) Feeds(ctx context.Context) ([]*model.Feed, error) {
	var feeds []*model.Feed
	res, err := r.rpcClient.ListFeeds(ctx, &pb.ListFeedsRequest{})
	if err != nil {
		return nil, errHelper.HandleGrpcErrors(ctx, err, "failed to list feeds")
	}
	c := converters.New()
	for _, feed := range res.Feeds {
		feeds = append(feeds, c.ProtoToGraphFeed(feed))
	}
	return feeds, nil
}

// Feed is the resolver for the feed field.
func (r *queryResolver) Feed(ctx context.Context, id string) (*model.Feed, error) {
	resp, err := r.rpcClient.GetFeed(ctx, &pb.GetFeedRequest{Id: id})
	if err != nil {
		return nil, errHelper.HandleGrpcErrors(ctx, err, "failed to get feed")
	}
	return converters.New().ProtoToGraphFeed(resp.Feed), nil
}

// Articles is the resolver for the articles field.
func (r *queryResolver) Articles(ctx context.Context, feedID string) ([]*model.Article, error) {
	resp, err := r.rpcClient.ListArticles(ctx, &pb.ListArticlesRequest{FeedId: feedID})
	if err != nil {
		return nil, errHelper.HandleGrpcErrors(ctx, err, "failed to list articles")
	}

	c := converters.New()
	var articles []*model.Article
	for _, article := range resp.Articles {
		articles = append(articles, c.ProtoToGraphArticle(article))
	}
	return articles, nil
}

// Article is the resolver for the article field.
func (r *queryResolver) Article(ctx context.Context, id string) (*model.Article, error) {
	resp, err := r.rpcClient.GetArticle(ctx, &pb.GetArticleRequest{Id: id})
	if err != nil {
		return nil, errHelper.HandleGrpcErrors(ctx, err, "failed to get article")
	}
	article := resp.Article
	return converters.New().ProtoToGraphArticle(article), nil
}

// Mutation returns MutationResolver implementation.
func (r *Resolver) Mutation() MutationResolver { return &mutationResolver{r} }

// Query returns QueryResolver implementation.
func (r *Resolver) Query() QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
