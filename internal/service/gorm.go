package service

import (
	"context"
	"errors"
	"fmt"

	"github.com/ericbutera/amalgam/internal/copygen"
	db_model "github.com/ericbutera/amalgam/internal/db/models"
	"github.com/ericbutera/amalgam/internal/sanitize"
	svc_model "github.com/ericbutera/amalgam/internal/service/models"
	"github.com/ericbutera/amalgam/internal/validate"
	"github.com/samber/lo"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

const DefaultLimit = 100

var ErrQueryFailed = errors.New("query failed")

type GormService struct {
	db *gorm.DB
}

func NewGorm(db *gorm.DB) Service {
	return &GormService{db: db}
}

// query returns a new query builder with the given context. required for otel
func (s *GormService) query(ctx context.Context) *gorm.DB {
	return s.db.WithContext(ctx)
}

func (s *GormService) Feeds(ctx context.Context) ([]svc_model.Feed, error) {
	var feeds []svc_model.Feed
	result := s.query(ctx).
		Where("is_active=?", true).
		Limit(DefaultLimit). // TODO: pagination
		Find(&feeds)

	if result.Error != nil {
		return nil, ErrQueryFailed
	}
	return feeds, nil
}

var (
	validateFeedCreate = validate.CustomMessages{
		"URL.required": "The URL is required.",
		"URL.url":      "The URL must be valid.",
	}
	validateArticleSave = validate.CustomMessages{
		"ID.required":       "The ID field is required.",
		"ID.uuid4":          "The ID must be a valid UUID.",
		"FeedID.required":   "The FeedID field is required.",
		"FeedID.uuid4":      "The FeedID must be a valid UUID.",
		"URL.required":      "The URL is required.",
		"URL.url":           "The URL must be valid.",
		"Title.required":    "The title is required.",
		"Title.min":         "The title must be at least 3 characters long.",
		"ImageURL.url":      "The Image URL must be a valid URL if provided.",
		"Preview.required":  "The preview is required.",
		"Preview.min":       "The preview must be at least 5 characters long.",
		"Content.required":  "The content is required.",
		"AuthorEmail.email": "The author email must be a valid email address if provided.",
	}
)

type CreateFeedResult struct {
	ID               string
	ValidationErrors []validate.ValidationError
}

func feedUrlExists(tx *gorm.DB, url string) error {
	res := tx.Find(&svc_model.Feed{}, "url=?", url)
	if res.Error != nil {
		return res.Error
	}
	if res.RowsAffected > 0 {
		return ErrDuplicate
	}
	return nil
}

func (s *GormService) CreateFeed(ctx context.Context, data *svc_model.Feed) (CreateFeedResult, error) {
	res := CreateFeedResult{}

	feed, err := sanitize.Struct(lo.FromPtr(data))
	if err != nil {
		return res, fmt.Errorf("unable to create feed: %w", err)
	}
	res.ValidationErrors = validate.Struct(feed, validateFeedCreate).Errors
	if len(res.ValidationErrors) > 0 {
		return res, ErrValidation
	}

	dbFeed := &db_model.Feed{}
	copygen.ServiceToDbFeed(dbFeed, &feed)

	dbFeed.IsActive = true

	err = s.query(ctx).Transaction(func(tx *gorm.DB) error {
		if err := feedUrlExists(tx, feed.URL); err != nil {
			return err
		}
		if err := tx.Create(&dbFeed).Error; err != nil {
			return err
		}
		res.ID = dbFeed.ID
		return nil
	})

	return res, err
}

func (s *GormService) UpdateFeed(ctx context.Context, id string, feed *svc_model.Feed) error {
	// note: no validation required for update
	dbFeed := &db_model.Feed{}
	copygen.ServiceToDbFeed(dbFeed, feed)
	result := s.query(ctx).
		Model(&dbFeed).
		Where("id=?", id).
		Updates(map[string]any{
			"name": feed.Name,
		})

	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return ErrNotFound
	}
	return nil
}

func (s *GormService) GetFeed(ctx context.Context, id string) (*svc_model.Feed, error) {
	var feed svc_model.Feed
	result := s.query(ctx).First(&feed, "id=?", id)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, ErrNotFound
		}
		return nil, result.Error
	}
	return &feed, nil
}

func (s *GormService) GetArticlesByFeed(ctx context.Context, feedId string) ([]svc_model.Article, error) {
	var articles []svc_model.Article
	result := s.query(ctx).
		Limit(DefaultLimit). // TODO: pagination (cursor)
		Find(&articles, "feed_id=?", feedId)

	if result.Error != nil {
		return nil, result.Error
	}
	return articles, nil
}

func (s *GormService) GetArticle(ctx context.Context, id string) (*svc_model.Article, error) {
	var article svc_model.Article
	result := s.query(ctx).First(&article, "id=?", id)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, ErrNotFound
		}
		return nil, result.Error
	}
	return &article, nil
}

type SaveArticleResult struct {
	ID               string
	ValidationErrors []validate.ValidationError
}

func (s *GormService) SaveArticle(ctx context.Context, data *svc_model.Article) (SaveArticleResult, error) {
	res := SaveArticleResult{}

	article, err := sanitize.Struct(lo.FromPtr(data))
	if err != nil {
		return res, err
	}
	res.ValidationErrors = validate.Struct(article, validateArticleSave).Errors

	if len(res.ValidationErrors) > 0 {
		return res, ErrValidation
	}

	dbArticle := &db_model.Article{}
	copygen.ServiceToDbArticle(dbArticle, &article)

	result := s.query(ctx).
		Model(&db_model.Article{}).
		Omit(clause.Associations).
		Create(&dbArticle)

	if err := result.Error; err != nil {
		return res, err
	}

	res.ID = dbArticle.ID
	return res, err
}
