package server

import (
	"log/slog"

	"github.com/ericbutera/amalgam/api/internal/models"
	"github.com/gin-gonic/gin"
	slogGorm "github.com/orandin/slog-gorm"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

const (
	MiddlewareName = "api"
)

type server struct {
	// Gin Router
	router *gin.Engine
	db     *gorm.DB
}

type ServerOption func(*server) error

func New(options ...ServerOption) (*server, error) {
	// TODO: ensure gin uses signal.NotifyContext ctx
	// TODO: ensure gin uses slog as logger (for otel exporter)

	s := &server{router: gin.Default()}

	for _, o := range options {
		if err := o(s); err != nil {
			return nil, err
		}
	}

	s.middleware()
	s.metrics()
	s.routes()

	return s, nil
}

func WithDatabase(name string) func(*server) error {
	if name == "" {
		name = "test.db"
	}
	gormLogger := slogGorm.New(
		slogGorm.WithTraceAll(), // TODO: only run in debug mode
	)
	return func(s *server) error {
		db, err := gorm.Open(sqlite.Open(name), &gorm.Config{
			Logger: gormLogger,
		})
		if err != nil {
			return err
		}

		// TODO: only migrate in debug mode
		db.AutoMigrate(&models.Feed{}, &models.Article{}, &models.User{})
		seed(db)

		s.db = db
		return nil
	}
}

// TODO populate with fixtures
// TODO: only run in debug mode
func seed(db *gorm.DB) {
	db.Transaction(func(tx *gorm.DB) error {
		var feed models.Feed
		result := tx.First(&feed, 1)

		if result.RowsAffected > 0 {
			slog.Debug("database already seeded")
			return nil
		}

		slog.Info("seeding database")
		feed = models.Feed{
			Url:  "https://example.com/",
			Name: "Example",
		}
		if err := tx.Create(&feed).Error; err != nil {
			return err
		}
		if err := tx.Create(&models.Article{
			FeedID:  feed.ID,
			Url:     "https://example.com/article",
			Title:   "Example Article",
			Content: "This is an example article",
		}).Error; err != nil {
			return err
		}
		return nil
	})
}

func (s *server) Run() error {
	return s.router.Run()
}

/*
Routes:
GET /health
POST /register
POST /login
POST /logout
GET /feeds
POST /feed
	- adds a feed source
	- generic to all users
	- create a user_feed record if exists
GET /feed/:id
GET /feed/:id/articles
GET /article/:id

*/
