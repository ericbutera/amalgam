package db

import (
	"errors"

	"github.com/ericbutera/amalgam/internal/db/models"
	"github.com/ericbutera/amalgam/internal/test/seed"
	"github.com/ericbutera/amalgam/pkg/config/env"
	slog "github.com/orandin/slog-gorm"
	"gorm.io/driver/mysql"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/plugin/opentelemetry/tracing"
)

var ErrInvalidAdapter = errors.New("invalid adapter")

type Adapters string

const (
	MysqlAdapter  Adapters = "mysql"
	SqliteAdapter Adapters = "sqlite"
)

var (
	ErrDsnNotSet  = errors.New("dsn not set")
	ErrNameNotSet = errors.New("name not set")
)

type Config struct {
	DbAdapter    Adapters `env:"DB_ADAPTER"     envDefault:"sqlite"`
	DbMysqlDsn   string   `env:"DB_MYSQL_DSN"   example:"user:pass@tcp(127.0.0.1:3306)/dbname?charset=utf8mb4&parseTime=True&loc=Local"`
	DbSqliteName string   `env:"DB_SQLITE_NAME" example:"file::memory:?cache=shared"`
}

// Convenience function to create a connection using Config.
func NewFromEnv() (*gorm.DB, error) {
	config, err := env.New[Config]()
	if err != nil {
		return nil, err
	}

	return NewFromConfig(config)
}

func NewFromConfig(config *Config) (*gorm.DB, error) {
	switch config.DbAdapter {
	case MysqlAdapter:
		return NewMysql(config.DbMysqlDsn, WithMiddleware(), WithTraceAll())
	case SqliteAdapter:
		return NewSqlite(
			config.DbSqliteName,
			WithAutoMigrate(),
			WithSeedData(),
			WithMiddleware(),
			WithTraceAll(),
		)
	default:
		return nil, ErrInvalidAdapter
	}
}

func setOpts(db *gorm.DB, opts ...Options) error {
	for _, opt := range opts {
		err := opt(db)
		if err != nil {
			return err
		}
	}

	return nil
}

func newDb(d gorm.Dialector, opts ...Options) (*gorm.DB, error) {
	db, err := gorm.Open(d)
	if err != nil {
		return nil, err
	}

	if err := setOpts(db, opts...); err != nil {
		return nil, err
	}

	return db, nil
}

func NewMysql(dsn string, opts ...Options) (*gorm.DB, error) {
	if dsn == "" {
		return nil, ErrDsnNotSet
	}

	return newDb(mysql.Open(dsn), opts...)
}

// Create a new sqlite database connection
// Runs migrations (sqlite is for local dev only)
func NewSqlite(name string, opts ...Options) (*gorm.DB, error) {
	if name == "" {
		return nil, ErrNameNotSet
	}

	return newDb(sqlite.Open(name), opts...)
}

type Options func(*gorm.DB) error

func WithTraceAll() Options {
	return func(db *gorm.DB) error {
		// Note: this should really be a two part setting of set logger and then logger opt of trace all
		// for now it is combined as one convenience function
		db.Logger = slog.New(slog.WithTraceAll())
		return nil
	}
}

func WithAutoMigrate() Options {
	return func(db *gorm.DB) error {
		return db.AutoMigrate(
			&models.Feed{},
			&models.Article{},
			&models.User{},
			&models.UserFeeds{},
			&models.UserArticles{},
			&models.FeedVerification{},
			&models.FetchHistory{},
		)
	}
}

// create starter data
// note: this is intended to be used with sqlite in memory as there isn't any cleanup
func WithSeedData() Options {
	return func(db *gorm.DB) error {
		return db.Transaction(func(tx *gorm.DB) error {
			if _, err := seed.FeedAndArticles(tx, 1); err != nil {
				return err
			}

			return nil
		})
	}
}

func WithMiddleware() Options {
	return middleware
}

func middleware(db *gorm.DB) error {
	return db.Use(tracing.NewPlugin())
}
