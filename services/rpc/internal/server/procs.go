package server

import (
	"context"
	"errors"

	"github.com/ericbutera/amalgam/internal/converters"
	"github.com/ericbutera/amalgam/internal/service"
	"github.com/ericbutera/amalgam/internal/validate"
	pb "github.com/ericbutera/amalgam/pkg/feeds/v1"
	"github.com/samber/lo"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

var ErrInvalidTaskType = errors.New("invalid task type")

func serviceToProtoErr(err error, validationErrs []validate.ValidationError) error {
	switch {
	case errors.Is(err, service.ErrNotFound):
		return status.Error(codes.NotFound, "not found")
	case errors.Is(err, service.ErrDuplicate):
		return status.Error(codes.AlreadyExists, "already exists")
	case errors.Is(err, service.ErrValidation):
		return validationErr(validationErrs)
	}
	return status.Errorf(codes.Internal, "failed to perform action: %v", err)
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

func (s *Server) ListFeeds(ctx context.Context, _ *pb.ListFeedsRequest) (*pb.ListFeedsResponse, error) {
	feeds, err := s.service.Feeds(ctx)
	if err != nil {
		return nil, serviceToProtoErr(err, nil)
	}
	pbFeeds := []*pb.Feed{}
	for _, feed := range feeds {
		pbFeed := converters.New().ServiceToProtoFeed(&feed)
		pbFeeds = append(pbFeeds, pbFeed)
	}
	return &pb.ListFeedsResponse{Feeds: pbFeeds}, nil
}

func (s *Server) CreateFeed(ctx context.Context, in *pb.CreateFeedRequest) (*pb.CreateFeedResponse, error) {
	feed := converters.New().ProtoCreateFeedToService(in.GetFeed())
	res, err := s.service.CreateFeed(ctx, feed)
	if err != nil {
		return nil, serviceToProtoErr(err, res.ValidationErrors)
	}
	return &pb.CreateFeedResponse{
		Id: res.ID,
	}, nil
}

func (s *Server) UpdateFeed(ctx context.Context, in *pb.UpdateFeedRequest) (*pb.UpdateFeedResponse, error) {
	feed := converters.New().ProtoUpdateFeedToService(in.GetFeed())
	if err := s.service.UpdateFeed(ctx, feed.ID, feed); err != nil {
		return nil, serviceToProtoErr(err, nil)
	}
	return &pb.UpdateFeedResponse{}, nil
}

func (s *Server) GetFeed(ctx context.Context, in *pb.GetFeedRequest) (*pb.GetFeedResponse, error) {
	feed, err := s.service.GetFeed(ctx, in.GetId())
	if err != nil {
		return nil, serviceToProtoErr(err, nil)
	}

	pbFeed := converters.New().ServiceToProtoFeed(feed)
	return &pb.GetFeedResponse{
		Feed: pbFeed,
	}, nil
}

func (s *Server) ListArticles(ctx context.Context, in *pb.ListArticlesRequest) (*pb.ListArticlesResponse, error) {
	options := in.GetOptions()
	result, err := s.service.GetArticlesByFeed(ctx, in.GetFeedId(), service.ListOptions{
		Cursor: options.GetCursor(),
		Limit:  int(options.GetLimit()),
	})
	if err != nil {
		return nil, serviceToProtoErr(err, nil)
	}
	c := converters.New()
	pbArticles := []*pb.Article{}
	for _, article := range result.Articles {
		pbArticles = append(pbArticles, c.ServiceToProtoArticle(&article))
	}
	return &pb.ListArticlesResponse{
		Articles: pbArticles,
		Pagination: &pb.Pagination{
			Previous: lo.FromPtr(result.Cursor.Before),
			Next:     lo.FromPtr(result.Cursor.After),
		},
	}, nil
}

func (s *Server) GetArticle(ctx context.Context, in *pb.GetArticleRequest) (*pb.GetArticleResponse, error) {
	article, err := s.service.GetArticle(ctx, in.GetId())
	if err != nil {
		return nil, serviceToProtoErr(err, nil)
	}
	pbArticle := converters.New().ServiceToProtoArticle(article)
	return &pb.GetArticleResponse{
		Article: pbArticle,
	}, nil
}

func (s *Server) SaveArticle(ctx context.Context, in *pb.SaveArticleRequest) (*pb.SaveArticleResponse, error) {
	article := converters.New().ProtoToServiceArticle(in.GetArticle())

	res, err := s.service.SaveArticle(ctx, article)
	if err != nil {
		return nil, serviceToProtoErr(err, res.ValidationErrors)
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

func (*Server) FeedTask(_ context.Context, _ *pb.FeedTaskRequest) (*pb.FeedTaskResponse, error) { //nolint
	return nil, status.Error(codes.Unimplemented, "feed task has been moved to graphql")
}
