package service

import (
	"context"
	"errors"

	svc_model "github.com/ericbutera/amalgam/internal/service/models"
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
	Articles   []svc_model.Article
	Pagination PaginationResult
}

// domain logic for feeds & articles
type Service interface {
	Feeds(ctx context.Context) ([]svc_model.Feed, error)
	CreateFeed(ctx context.Context, feed *svc_model.Feed) (CreateFeedResult, error)
	UpdateFeed(ctx context.Context, id string, feed *svc_model.Feed) error
	GetFeed(ctx context.Context, id string) (*svc_model.Feed, error)
	GetArticlesByFeed(ctx context.Context, feedId string, options ListOptions) (*ArticlesByFeedResult, error)
	GetArticle(ctx context.Context, id string) (*svc_model.Article, error)
	SaveArticle(ctx context.Context, article *svc_model.Article) (SaveArticleResult, error)
}
