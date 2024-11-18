package config

import "github.com/spf13/viper"

type Config struct {
	CorsAllowOrigins  []string `mapstructure:"cors_allow_origins"`
	CorsAllowMethods  []string `mapstructure:"cors_allow_methods"`
	CorsAllowHeaders  []string `mapstructure:"cors_allow_headers"`
	CorsExposeHeaders []string `mapstructure:"cors_expose_headers"`
	GraphHost         string   `mapstructure:"graph_host"`
	// TODO:
	// log level
}

func init() { //nolint:gochecknoinits
	viper.SetDefault("cors_allow_origins", []string{})
	viper.SetDefault("cors_allow_methods", []string{})
	viper.SetDefault("cors_allow_headers", []string{})
	viper.SetDefault("cors_expose_headers", []string{})
	viper.SetDefault("graph_host", "")
}
