package service

import (
	"context"
	"errors"

	"github.com/ericbutera/amalgam/internal/db/pagination"
	svc_model "github.com/ericbutera/amalgam/internal/service/models"
)

var (
	ErrNotFound   = errors.New("not found")
	ErrDuplicate  = errors.New("duplicate entry")
	ErrValidation = errors.New("validation error")
)

// type ListOptions = pagination.ListOptions

type ArticlesByFeedResult struct {
	Articles []svc_model.Article
	Cursor   pagination.Cursor
}

type GetUserFeedsResult struct {
	Feeds []svc_model.UserFeed
}

// domain logic for feeds & articles
type Service interface {
	Feeds(ctx context.Context /*, options *FeedsOptions*/) ([]svc_model.Feed, error) // Fetch all feeds in system (not intended for public use)
	CreateFeed(ctx context.Context, feed *svc_model.Feed) (CreateFeedResult, error)
	UpdateFeed(ctx context.Context, id string, feed *svc_model.Feed) error
	UpdateFeedArticleCount(ctx context.Context, feedID string) error
	GetFeed(ctx context.Context, id string) (*svc_model.Feed, error)
	GetArticlesByFeed(ctx context.Context, feedId string, options pagination.ListOptions) (*ArticlesByFeedResult, error)
	GetArticle(ctx context.Context, id string) (*svc_model.Article, error)
	SaveArticle(ctx context.Context, article *svc_model.Article) (SaveArticleResult, error)
	GetUserFeed(ctx context.Context, userID string, feedID string) (*svc_model.UserFeed, error)
	GetUserFeeds(ctx context.Context, userID string) (*GetUserFeedsResult, error) // Fetch by a specific user
	SaveUserFeed(ctx context.Context, userFeed *svc_model.UserFeed) error         // Associate a feed with a user
	GetUserArticles(ctx context.Context, userID string, articleIDs []string) ([]*svc_model.UserArticle, error)
	SaveUserArticle(ctx context.Context, userArticle *svc_model.UserArticle) error
}
