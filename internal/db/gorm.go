package db

import (
	"log/slog"

	"github.com/ericbutera/amalgam/internal/db/models"
	"github.com/go-errors/errors"
	slogGorm "github.com/orandin/slog-gorm"
	"gorm.io/driver/mysql"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/plugin/opentelemetry/tracing"
)

func newLogger() logger.Interface {
	return slogGorm.New(
		slogGorm.WithTraceAll(), // TODO: only run in debug mode
	)
}

func newConfig() *gorm.Config {
	return &gorm.Config{
		Logger: newLogger(),
	}
}

func NewMysql(dsn string) (*gorm.DB, error) {
	if dsn == "" {
		return nil, errors.New("mysql: dsn not set")
	}
	db, err := gorm.Open(mysql.Open(dsn), newConfig())
	if err != nil {
		return nil, err
	}
	if err := middleware(db); err != nil {
		return nil, err
	}
	return db, nil
}

// Create a new sqlite database connection
// Runs migrations (sqlite is for local dev only)
func NewSqlite(name string) (*gorm.DB, error) {
	if name == "" {
		return nil, errors.New("sqlite: name not set")
	}
	db, err := gorm.Open(sqlite.Open(name), newConfig())
	if err != nil {
		return nil, err
	}
	if err := middleware(db); err != nil {
		return nil, err
	}
	if err := seedSqlite(db); err != nil {
		return nil, err
	}
	return db, nil
}

// TODO populate with fixtures
// TODO: only run in debug mode
func seedSqlite(db *gorm.DB) error {
	if err := db.AutoMigrate(&models.Feed{}, &models.Article{}, &models.User{}); err != nil {
		return err
	}
	return db.Transaction(func(tx *gorm.DB) error {
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

func middleware(db *gorm.DB) error {
	return db.Use(tracing.NewPlugin())
}
