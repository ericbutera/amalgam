package service

import (
	"context"
	"errors"

	"github.com/ericbutera/amalgam/internal/db/models"
	"gorm.io/gorm"
)

var (
	ErrNotFound = errors.New("not found")
)

// TODO: move validation from api into service
// this will include changing go tags binding -> validation

// domain logic for feeds & articles
type Service struct {
	db *gorm.DB
}

func New(db *gorm.DB) *Service {
	return &Service{db: db}
}

// query returns a new query builder with the given context. required for otel
func (s *Service) query(ctx context.Context) *gorm.DB {
	return s.db.WithContext(ctx)
}

func (s *Service) Feeds(ctx context.Context) ([]models.Feed, error) {
	var feeds []models.Feed
	result := s.query(ctx).
		Find(&feeds).
		Limit(100) // TODO: pagination
	if result.Error != nil {
		return nil, errors.New("failed to fetch feeds")
	}
	return feeds, nil
}

func (s *Service) CreateFeed(ctx context.Context, feed *models.Feed) error {
	// TODO: normalize URL to prevent duplicates
	// TODO: create user_feed if not exists
	return s.query(ctx).Transaction(func(tx *gorm.DB) error {
		res := tx.Find(&models.Feed{}, "url = ?", feed.Url)
		if res.Error != nil {
			return res.Error
		}
		if res.RowsAffected > 0 {
			return errors.New("feed already exists")
		}
		if err := tx.Create(feed).Error; err != nil {
			return err
		}
		return nil
	})
}

func (s *Service) UpdateFeed(ctx context.Context, id string, feed *models.Feed) error {
	result := s.query(ctx).Model(&feed).Where("id=?", id).Updates(map[string]interface{}{
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

// TODO: make this update user's feed (user_feed)
//UpdateUserFeed() error {

func (s *Service) GetFeed(ctx context.Context, id string) (*models.Feed, error) {
	var feed models.Feed
	result := s.query(ctx).First(&feed, id)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, ErrNotFound
		}
		return nil, result.Error
	}
	return &feed, nil
}

func (s *Service) GetArticlesByFeed(ctx context.Context, id string) ([]models.Article, error) {
	var articles []models.Article
	result := s.query(ctx).
		Find(&articles, "feed_id = ?", id).
		Limit(100) // TODO: pagination (cursor)
	if result.Error != nil {
		return nil, result.Error
	}
	return articles, nil
}

func (s *Service) GetArticle(ctx context.Context, id string) (*models.Article, error) {
	var article models.Article
	result := s.query(ctx).First(&article, id)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, ErrNotFound
		}
		return nil, result.Error
	}
	return &article, nil
}
