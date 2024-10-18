package service

import (
	"context"
	"errors"
	"strconv"

	"github.com/ericbutera/amalgam/api/internal/models"
	"gorm.io/gorm"
)

var (
	ErrNotFound = errors.New("not found")
)

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
	result := s.query(ctx).Create(feed)
	if result.Error != nil {
		return errors.New("failed to create feed")
	}
	return nil
}

func parseUint(s string) (uint, error) {
	number, err := strconv.ParseUint(s, 10, 64)
	if err != nil {
		return 0, err
	}
	return uint(number), nil
}

func (s *Service) UpdateFeed(ctx context.Context, id string, feed *models.Feed) error {
	// normalize URL to prevent duplicates
	// create feed if not exists (feeds are global)
	// create user_feed if not exists
	uid, err := parseUint(id)
	if err != nil {
		return errors.New("invalid feed id")
	}
	feed.Base.ID = uid
	result := s.query(ctx).Save(feed)
	if result.Error != nil {
		return errors.New("failed to update feed")
	}
	return nil
}

func (s *Service) GetFeed(ctx context.Context, id string) (*models.Feed, error) {
	var feed models.Feed
	result := s.query(ctx).First(&feed, id)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, ErrNotFound
		}
		return nil, errors.New("failed to fetch feed")
	}
	return &feed, nil
}

func (s *Service) GetArticlesByFeed(ctx context.Context, id string) ([]models.Article, error) {
	var articles []models.Article
	result := s.query(ctx).
		Find(&articles, "feed_id = ?", id).
		Limit(100) // TODO: pagination (cursor)
	if result.Error != nil {
		return nil, errors.New("failed to fetch articles")
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
		return nil, errors.New("failed to fetch article")
	}
	return &article, nil
}
