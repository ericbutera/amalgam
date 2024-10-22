package db

import (
	"errors"

	"github.com/ericbutera/amalgam/pkg/config"
	"github.com/spf13/viper"
	"gorm.io/gorm"
)

type Adapters string

const MysqlAdapter Adapters = "mysql"
const SqliteAdapter Adapters = "sqlite"

type Config struct {
	DbAdapter    Adapters `mapstructure:"db_adapter"`
	DbMysqlDsn   string   `mapstructure:"db_mysql_dsn"`
	DbSqliteName string   `mapstructure:"db_sqlite_name"`
}

func init() {
	viper.SetDefault("db_adapter", SqliteAdapter)
	viper.SetDefault("db_sqlite_name", "test.sqlite")
	viper.SetDefault("db_mysql_dsn", "user:pass@tcp(127.0.0.1:3306)/dbname?charset=utf8mb4&parseTime=True&loc=Local")
}

// Convenience function to create a connection using Config.
func NewFromEnv() (*gorm.DB, error) {
	cfg, err := config.NewFromEnv[Config]()
	if err != nil {
		return nil, err
	}
	return NewFromConfig(cfg)
}

func NewFromConfig(cfg *Config) (*gorm.DB, error) {
	switch cfg.DbAdapter {
	case MysqlAdapter:
		return NewMysql(cfg.DbMysqlDsn)
	case SqliteAdapter:
		return NewSqlite(cfg.DbSqliteName)
	}
	return nil, errors.New("db adapter not supported")
}
