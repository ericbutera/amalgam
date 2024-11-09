package db

import (
	"errors"

	"github.com/ericbutera/amalgam/internal/copygen"
	"github.com/ericbutera/amalgam/internal/db/models"
	"github.com/ericbutera/amalgam/internal/test/fixtures"
	"github.com/ericbutera/amalgam/pkg/config/env"
	slog "github.com/orandin/slog-gorm"
	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/plugin/opentelemetry/tracing"
)

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
	DbAdapter    Adapters `mapstructure:"db_adapter"`
	DbMysqlDsn   string   `mapstructure:"db_mysql_dsn"`
	DbSqliteName string   `mapstructure:"db_sqlite_name"`
}

func init() {
	viper.SetDefault("db_adapter", SqliteAdapter)
	viper.SetDefault("db_sqlite_name", "file::memory:?cache=shared")
	viper.SetDefault("db_mysql_dsn", "user:pass@tcp(127.0.0.1:3306)/dbname?charset=utf8mb4&parseTime=True&loc=Local")
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
	}
	return nil, errors.New("db adapter not supported")
}

func setOpts(db *gorm.DB, opts ...DbOptions) error {
	for _, opt := range opts {
		if err := opt(db); err != nil {
			return err
		}
	}
	return nil
}

func newDb(d gorm.Dialector, opts ...DbOptions) (*gorm.DB, error) {
	db, err := gorm.Open(d)
	if err != nil {
		return nil, err
	}
	if err := setOpts(db, opts...); err != nil {
		return nil, err
	}
	return db, nil
}

func NewMysql(dsn string, opts ...DbOptions) (*gorm.DB, error) {
	if dsn == "" {
		return nil, ErrDsnNotSet
	}
	return newDb(mysql.Open(dsn), opts...)
}

// Create a new sqlite database connection
// Runs migrations (sqlite is for local dev only)
func NewSqlite(name string, opts ...DbOptions) (*gorm.DB, error) {
	if name == "" {
		return nil, ErrNameNotSet
	}
	return newDb(sqlite.Open(name), opts...)
}

type DbOptions func(*gorm.DB) error

func WithTraceAll() DbOptions {
	return func(db *gorm.DB) error {
		// Note: this should really be a two part setting of set logger and then logger opt of trace all
		// for now it is combined as one convenience function
		db.Logger = slog.New(slog.WithTraceAll())
		return nil
	}
}

func WithAutoMigrate() DbOptions {
	return func(db *gorm.DB) error {
		return db.AutoMigrate(&models.Feed{}, &models.Article{}, &models.User{})
	}
}

// create starter data
// note: this is intended to be used with sqlite in memory as there isn't any cleanup
func WithSeedData() DbOptions {
	return func(db *gorm.DB) error {
		return db.Transaction(func(tx *gorm.DB) error {
			feed := models.Feed{}
			copygen.ServiceToDbFeed(&feed, fixtures.NewFeed())

			article := models.Article{}
			copygen.ServiceToDbArticle(&article, fixtures.NewArticle())
			if err := tx.Create(&feed).Error; err != nil {
				return err
			}
			if err := tx.Create(&article).Error; err != nil {
				return err
			}
			return nil
		})
	}
}

func WithMiddleware() DbOptions {
	return func(db *gorm.DB) error {
		return middleware(db)
	}
}

func middleware(db *gorm.DB) error {
	return db.Use(tracing.NewPlugin())
}
