package service

import (
	"context"
	"errors"

	svc_model "github.com/ericbutera/amalgam/internal/service/models"
	"github.com/pilagod/gorm-cursor-paginator/v2/paginator"
)

var (
	ErrNotFound   = errors.New("not found")
	ErrDuplicate  = errors.New("duplicate entry")
	ErrValidation = errors.New("validation error")
)

type ListOptions struct {
	Cursor string // Encoded cursor for current page
	Limit  int    // Limit for pagination
	// Filters   map[string]interface{} // Dynamic filters (key-value pairs)
}

type PaginationResult struct {
	NextCursor string // Encoded cursor for the next page
}

type ArticlesByFeedResult struct {
	Articles []svc_model.Article
	Cursor   paginator.Cursor // TODO: make custom type (don't leak internal impl)
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
	GetUserFeed(ctx context.Context, userID string, feedID string) (*svc_model.UserFeed, error)
	GetArticlesByFeed(ctx context.Context, feedId string, options ListOptions) (*ArticlesByFeedResult, error)
	GetArticle(ctx context.Context, id string) (*svc_model.Article, error)
	SaveArticle(ctx context.Context, article *svc_model.Article) (SaveArticleResult, error)
	GetUserFeeds(ctx context.Context, userID string) (*GetUserFeedsResult, error) // Fetch by a specific user
	SaveUserFeed(ctx context.Context, userFeed *svc_model.UserFeed) error         // Associate a feed with a user
}
