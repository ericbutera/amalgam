package config

import "github.com/spf13/viper"

type Config struct {
	CorsAllowOrigins  []string `mapstructure:"cors_allow_origins"`
	CorsAllowMethods  []string `mapstructure:"cors_allow_methods"`
	CorsAllowHeaders  []string `mapstructure:"cors_allow_headers"`
	CorsExposeHeaders []string `mapstructure:"cors_expose_headers"`
	DbAdapter         string   `mapstructure:"db_adapter"`
	DbMysqlDsn        string   `mapstructure:"db_mysql_dsn"`
	DbSqliteName      string   `mapstructure:"db_sqlite_name"`
	// TODO:
	// log level
}

func init() {
	viper.SetDefault("cors_allow_origins", []string{})
	viper.SetDefault("cors_allow_methods", []string{})
	viper.SetDefault("cors_allow_headers", []string{})
	viper.SetDefault("cors_expose_headers", []string{})
	viper.SetDefault("db_adapter", "sqlite")
	viper.SetDefault("db_sqlite_name", "test.db")
	viper.SetDefault("db_mysql_dsn", "user:pass@tcp(127.0.0.1:3306)/dbname?charset=utf8mb4&parseTime=True&loc=Local")
}
