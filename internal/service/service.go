package service

import (
	"context"
	"errors"

	//converter "github.com/ericbutera/amalgam/internal/copygen"
	"github.com/ericbutera/amalgam/internal/db/models"
	"gorm.io/gorm"
)

var (
	ErrNotFound  = errors.New("not found")
	ErrDuplicate = errors.New("duplicate entry")
)

type Feed struct {
	ID   string `json:"id"`
	Name string `json:"name"`
	Url  string `json:"url"`
}

type Article struct {
	ID          string `json:"id"`
	FeedID      string `json:"feedId"`
	Url         string `json:"url"`
	Title       string `json:"title"`
	ImageUrl    string `json:"imageUrl"`
	Preview     string `json:"preview"`
	Content     string `json:"content"`
	Guid        string `json:"guid"`
	AuthorName  string `json:"authorName"`
	AuthorEmail string `json:"authorEmail"`
}

type Service interface {
	Feeds(ctx context.Context) ([]Feed, error)
	CreateFeed(ctx context.Context, feed *Feed) error
	UpdateFeed(ctx context.Context, id string, feed *Feed) error
	GetFeed(ctx context.Context, id string) (*Feed, error)
	GetArticlesByFeed(ctx context.Context, feedId string) ([]Article, error)
	GetArticle(ctx context.Context, id string) (*Article, error)
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

func (s *GormService) Feeds(ctx context.Context) ([]Feed, error) {
	//var feeds []models.Feed
	var feeds []Feed
	result := s.query(ctx).
		Find(&feeds).
		Limit(100) // TODO: pagination

	if result.Error != nil {
		return nil, errors.New("failed to fetch feeds")
	}
	return feeds, nil
}

func (s *GormService) CreateFeed(ctx context.Context, feed *Feed) error {
	// TODO: normalize URL to prevent duplicates
	// TODO: create user_feed if not exists
	//var dbFeed models.Feed
	dbFeed := models.Feed{
		Name: feed.Name,
		Url:  feed.Url,
	}
	// converter.ServiceToDbFeed(dbFeed, feed)
	return s.query(ctx).Transaction(func(tx *gorm.DB) error {
		res := tx.Find(&Feed{}, "url=?", feed.Url)
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

func (s *GormService) UpdateFeed(ctx context.Context, id string, feed *Feed) error {
	result := s.query(ctx).
		Model(&feed).
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

func (s *GormService) GetFeed(ctx context.Context, id string) (*Feed, error) {
	//var feed models.Feed
	var feed Feed
	result := s.query(ctx).First(&feed, "id=?", id)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, ErrNotFound
		}
		return nil, result.Error
	}
	return &feed, nil
}

func (s *GormService) GetArticlesByFeed(ctx context.Context, feedId string) ([]Article, error) {
	var articles []Article //[]models.Article
	result := s.query(ctx).
		Find(&articles, "feed_id=?", feedId).
		Limit(100) // TODO: pagination (cursor)

	if result.Error != nil {
		return nil, result.Error
	}
	return articles, nil
}

func (s *GormService) GetArticle(ctx context.Context, id string) (*Article, error) {
	//var article models.Article
	var article Article
	result := s.query(ctx).First(&article, "id=?", id)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, ErrNotFound
		}
		return nil, result.Error
	}
	return &article, nil
}
