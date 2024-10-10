package service

import (
	"errors"

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

func (s *Service) Feeds() ([]models.Feed, error) {
	var feeds []models.Feed
	result := s.db.Find(&feeds).Limit(100)
	if result.Error != nil {
		return nil, errors.New("failed to fetch feeds")
	}
	return feeds, nil
}

func (s *Service) GetFeed(id string) (*models.Feed, error) {
	var feed models.Feed
	result := s.db.First(&feed, id)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, ErrNotFound
		}
		return nil, errors.New("failed to fetch feed")
	}
	return &feed, nil
}

func (s *Service) GetArticlesByFeed(feedId string) ([]models.Article, error) {
	var articles []models.Article
	result := s.db.Find(&articles, "feed_id = ?", feedId).Limit(100)
	if result.Error != nil {
		return nil, errors.New("failed to fetch articles")
	}
	return articles, nil
}
