package server

import (
	"context"
	"errors"

	"github.com/ericbutera/amalgam/internal/service"
	"github.com/ericbutera/amalgam/internal/service/models"
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
		pbFeeds = append(pbFeeds, s.converters.ServiceToProtoFeed(&feed))
	}
	return &pb.ListFeedsResponse{Feeds: pbFeeds}, nil
}

func (s *Server) ListUserFeeds(ctx context.Context, in *pb.ListUserFeedsRequest) (*pb.ListUserFeedsResponse, error) {
	userID := in.GetUser().GetId()
	res, err := s.service.GetUserFeeds(ctx, userID)
	if err != nil {
		return nil, serviceToProtoErr(err, nil)
	}
	feeds := []*pb.UserFeed{}
	for _, feed := range res.Feeds {
		feeds = append(feeds, s.converters.ServiceToProtoUserFeed(&feed))
	}
	return &pb.ListUserFeedsResponse{Feeds: feeds}, nil
}

func (s *Server) GetUserArticles(ctx context.Context, in *pb.GetUserArticlesRequest) (*pb.GetUserArticlesResponse, error) {
	userID := in.GetUser().GetId()
	articleIDs := in.GetArticleIds()
	res, err := s.service.GetUserArticles(ctx, userID, articleIDs)
	if err != nil {
		return nil, serviceToProtoErr(err, nil)
	}

	articles := map[string]*pb.GetUserArticlesResponse_UserArticle{}
	for _, article := range res {
		articles[article.ArticleID] = s.converters.ServiceToProtoUserArticle(article)
	}

	return &pb.GetUserArticlesResponse{Articles: articles}, nil
}

func (s *Server) CreateFeed(ctx context.Context, in *pb.CreateFeedRequest) (*pb.CreateFeedResponse, error) {
	feed := s.converters.ProtoCreateFeedToService(in.GetFeed())

	// 1. create feed (if not exists)
	res, err := s.service.CreateFeed(ctx, feed)
	if err != nil {
		return nil, serviceToProtoErr(err, res.ValidationErrors)
	}

	// 2. associate feed with user
	if in.GetUser() != nil && in.GetUser().GetId() != "" {
		uf := &models.UserFeed{
			UserID: in.GetUser().GetId(),
			FeedID: res.ID,
		}
		if err := s.service.SaveUserFeed(ctx, uf); err != nil {
			return nil, serviceToProtoErr(err, nil)
		}
	}

	return &pb.CreateFeedResponse{
		Id: res.ID,
	}, nil
}

func (s *Server) UpdateFeed(ctx context.Context, in *pb.UpdateFeedRequest) (*pb.UpdateFeedResponse, error) {
	feed := s.converters.ProtoUpdateFeedToService(in.GetFeed())
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
	return &pb.GetFeedResponse{
		Feed: s.converters.ServiceToProtoFeed(feed),
	}, nil
}

func (s *Server) GetUserFeed(ctx context.Context, in *pb.GetUserFeedRequest) (*pb.GetUserFeedResponse, error) {
	feed, err := s.service.GetUserFeed(ctx, in.GetUserId(), in.GetFeedId())
	if err != nil {
		return nil, serviceToProtoErr(err, nil)
	}
	return &pb.GetUserFeedResponse{
		Feed: s.converters.ServiceToProtoUserFeed(feed),
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
	articles := []*pb.Article{}
	for _, article := range result.Articles {
		articles = append(articles, s.converters.ServiceToProtoArticle(&article))
	}
	return &pb.ListArticlesResponse{
		Articles: articles,
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
	return &pb.GetArticleResponse{
		Article: s.converters.ServiceToProtoArticle(article),
	}, nil
}

func (s *Server) SaveArticle(ctx context.Context, in *pb.SaveArticleRequest) (*pb.SaveArticleResponse, error) {
	article := s.converters.ProtoToServiceArticle(in.GetArticle())
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

func (s *Server) UpdateStats(ctx context.Context, in *pb.UpdateStatsRequest) (*pb.UpdateStatsResponse, error) {
	var err error

	if in.GetStat() == pb.UpdateStatsRequest_STAT_FEED_ARTICLE_COUNT {
		feedID := in.GetFeedId()
		if feedID == "" {
			return nil, status.Error(codes.InvalidArgument, "feed_id is required")
		}
		err = s.service.UpdateFeedArticleCount(ctx, feedID)
	}

	if err != nil {
		return nil, serviceToProtoErr(err, nil)
	}

	return &pb.UpdateStatsResponse{}, nil
}

func (s *Server) Ready(_ context.Context, _ *pb.ReadyRequest) (*pb.ReadyResponse, error) {
	// TODO: without this the graph and rpc services are "ready" before the database is actually ready
	res := s.db.Exec("SELECT 1 FROM feeds LIMIT 1")
	if res.Error != nil {
		return nil, status.Error(codes.Internal, "database not ready")
	}
	return &pb.ReadyResponse{}, nil
}

func (s *Server) MarkArticleAsRead(ctx context.Context, in *pb.MarkArticleAsReadRequest) (*pb.MarkArticleAsReadResponse, error) {
	err := s.service.SaveUserArticle(ctx, &models.UserArticle{
		UserID:    in.GetUser().GetId(),
		ArticleID: in.GetArticleId(),
	})
	if err != nil {
		return nil, serviceToProtoErr(err, nil)
	}
	return &pb.MarkArticleAsReadResponse{}, nil
}
