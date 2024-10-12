package server

import (
	"errors"
	"log/slog"

	"github.com/ericbutera/amalgam/api/internal/config"
	"github.com/ericbutera/amalgam/api/internal/models"
	"github.com/ericbutera/amalgam/api/internal/service"
	"github.com/gin-gonic/gin"
	slogGorm "github.com/orandin/slog-gorm"
	"gorm.io/driver/mysql"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

const (
	MiddlewareName = "api"
)

type server struct {
	config  *config.Config
	router  *gin.Engine
	db      *gorm.DB
	service *service.Service
}

type ServerOption func(*server) error

func New(options ...ServerOption) (*server, error) {
	s := &server{
		router: gin.New(),
	}

	for _, o := range options {
		if err := o(s); err != nil {
			return nil, err
		}
	}

	if s.db == nil {
		return nil, errors.New("database not set")
	}

	s.service = service.New(s.db)
	s.middleware()
	s.metrics()
	s.routes()

	return s, nil
}

func WithConfig(cfg *config.Config) func(*server) error {
	return func(s *server) error {
		s.config = cfg
		return nil
	}
}

func WithSqlite(name string) func(*server) error {
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

func WithMysql(dsn string) func(*server) error {
	gormLogger := slogGorm.New(
		slogGorm.WithTraceAll(), // TODO: only run in debug mode
	)
	return func(s *server) error {
		db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
			Logger: gormLogger,
		})
		if err != nil {
			return err
		}

		// TODO: only migrate in debug mode
		// db.AutoMigrate(&models.Feed{}, &models.Article{}, &models.User{})

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
