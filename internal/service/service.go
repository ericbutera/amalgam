package service

import (
	"context"
	"errors"

	//converter "github.com/ericbutera/amalgam/internal/copygen"

	"github.com/ericbutera/amalgam/internal/copygen"
	db_model "github.com/ericbutera/amalgam/internal/db/models"
	svc_model "github.com/ericbutera/amalgam/internal/service/models"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

var (
	ErrNotFound  = errors.New("not found")
	ErrDuplicate = errors.New("duplicate entry")
)

type Service interface {
	Feeds(ctx context.Context) ([]svc_model.Feed, error)
	CreateFeed(ctx context.Context, feed *svc_model.Feed) error
	UpdateFeed(ctx context.Context, id string, feed *svc_model.Feed) error
	GetFeed(ctx context.Context, id string) (*svc_model.Feed, error)
	GetArticlesByFeed(ctx context.Context, feedId string) ([]svc_model.Article, error)
	GetArticle(ctx context.Context, id string) (*svc_model.Article, error)
	SaveArticle(ctx context.Context, article *svc_model.Article) error
}

// TODO: move validation from api into service
// this will include changing go tags binding -> validation

// domain logic for feeds & articles
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
		Find(&feeds).
		Limit(100) // TODO: pagination

	if result.Error != nil {
		return nil, errors.New("failed to fetch feeds")
	}
	return feeds, nil
}

func (s *GormService) CreateFeed(ctx context.Context, feed *svc_model.Feed) error {
	// TODO: normalize URL to prevent duplicates
	// TODO: validation
	dbFeed := &db_model.Feed{}
	copygen.ServiceToDbFeed(dbFeed, feed)
	return s.query(ctx).Transaction(func(tx *gorm.DB) error {
		res := tx.Find(&svc_model.Feed{}, "url=?", feed.Url)
		if res.Error != nil {
			return res.Error
		}
		if res.RowsAffected > 0 {
			return ErrDuplicate
		}
		if err := tx.Create(&dbFeed).Error; err != nil {
			return err
		}
		feed.ID = dbFeed.ID
		return nil
	})
}

func (s *GormService) UpdateFeed(ctx context.Context, id string, feed *svc_model.Feed) error {
	// TODO: validation
	dbFeed := &db_model.Feed{}
	copygen.ServiceToDbFeed(dbFeed, feed)
	result := s.query(ctx).
		Model(&dbFeed).
		Where("id=?", id).
		Updates(map[string]interface{}{
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
		Find(&articles, "feed_id=?", feedId).
		Limit(100) // TODO: pagination (cursor)

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

func (s *GormService) SaveArticle(ctx context.Context, article *svc_model.Article) error {
	// TODO: validation
	if article.FeedID == "" {
		return errors.New("missing feed ID")
	}
	dbArticle := &db_model.Article{}
	copygen.ServiceToDbArticle(dbArticle, article)
	result := s.query(ctx).
		Model(&db_model.Article{}).
		Omit(clause.Associations).
		Create(&dbArticle)
	if err := result.Error; err != nil {
		return err
	}
	article.ID = dbArticle.ID
	return nil
}
