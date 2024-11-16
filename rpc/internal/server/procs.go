package server

import (
	"context"
	"errors"

	"github.com/ericbutera/amalgam/internal/copygen"
	"github.com/ericbutera/amalgam/internal/service"
	"github.com/ericbutera/amalgam/internal/service/models"
	"github.com/ericbutera/amalgam/internal/validate"
	pb "github.com/ericbutera/amalgam/pkg/feeds/v1"
	"github.com/ericbutera/amalgam/rpc/internal/tasks"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) ListFeeds(ctx context.Context, _ *pb.ListFeedsRequest) (*pb.ListFeedsResponse, error) {
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
	copygen.ProtoCreateFeedToService(feed, in.GetFeed())
	res, err := s.service.CreateFeed(ctx, feed)
	if err != nil {
		if errors.Is(err, service.ErrDuplicate) {
			return nil, status.Errorf(codes.AlreadyExists, "feed already exists")
		} else if errors.Is(err, service.ErrValidation) {
			return nil, validationErr(res.ValidationErrors)
		}
		return nil, status.Errorf(codes.Internal, "failed to create feed: %v", err)
	}
	return &pb.CreateFeedResponse{
		Id: res.ID,
	}, nil
}

func validationErr(errors []validate.ValidationError) error {
	st := status.New(codes.InvalidArgument, "validation error")
	ds, err := st.WithDetails(&pb.ValidationErrors{
		Errors: validationErrToPb(errors),
	})
	if err != nil {
		return err
	}
	return ds.Err()
}

func (s *Server) UpdateFeed(ctx context.Context, in *pb.UpdateFeedRequest) (*pb.UpdateFeedResponse, error) {
	feed := &models.Feed{}
	copygen.ProtoUpdateFeedToService(feed, in.GetFeed())
	if err := s.service.UpdateFeed(ctx, feed.ID, feed); err != nil {
		return nil, status.Errorf(codes.Internal, "failed to create feed: %v", err)
	}
	return &pb.UpdateFeedResponse{}, nil
}

func (s *Server) GetFeed(ctx context.Context, in *pb.GetFeedRequest) (*pb.GetFeedResponse, error) {
	feed, err := s.service.GetFeed(ctx, in.GetId())
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
	feedId := in.GetFeedId()
	if feedId == "" {
		return nil, status.Errorf(codes.InvalidArgument, "feed_id is required")
	}
	articles, err := s.service.GetArticlesByFeed(ctx, feedId)
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
	article, err := s.service.GetArticle(ctx, in.GetId())
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to fetch article: %v", err)
	}
	pbArticle := pb.Article{}
	copygen.ServiceToProtoArticle(&pbArticle, article)
	return &pb.GetArticleResponse{
		Article: &pbArticle,
	}, nil
}

func (s *Server) SaveArticle(ctx context.Context, in *pb.SaveArticleRequest) (*pb.SaveArticleResponse, error) {
	article := &models.Article{}

	copygen.ProtoToServiceArticle(article, in.GetArticle())
	article.FeedID = in.GetArticle().GetFeedId() // TODO: bug with ProtoToServiceArticle

	res, err := s.service.SaveArticle(ctx, article)
	if err != nil {
		if errors.Is(err, service.ErrValidation) {
			return nil, validationErr(res.ValidationErrors)
		}
		return nil, status.Errorf(codes.Internal, "failed to save article: %v", err)
	}
	return &pb.SaveArticleResponse{
		Id: res.ID,
	}, nil
}

func validationErrToPb(errs []validate.ValidationError) []*pb.ValidationError {
	protoErrs := []*pb.ValidationError{}
	for _, err := range errs {
		protoErrs = append(protoErrs, &pb.ValidationError{
			Field:      err.Field,
			Tag:        err.Tag,
			RawMessage: err.RawMessage,
			Message:    err.FriendlyMessage,
		})
	}
	return protoErrs
}

func pbTaskToTaskType(task pb.FeedTaskRequest_Task) (tasks.TaskType, error) {
	if task == pb.FeedTaskRequest_TASK_GENERATE_FEEDS {
		return tasks.TaskGenerateFeeds, nil
	}
	return tasks.TaskUnspecified, errors.New("invalid task type")
}

func (*Server) FeedTask(ctx context.Context, in *pb.FeedTaskRequest) (*pb.FeedTaskResponse, error) {
	taskType, err := pbTaskToTaskType(in.GetTask())
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "invalid task type")
	}

	task, err := tasks.New(ctx, taskType) // TODO: dependency injection
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to start feed task: %s", err)
	}
	return &pb.FeedTaskResponse{
		Id: task.ID,
	}, nil
}
