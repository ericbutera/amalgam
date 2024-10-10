package config

import "github.com/spf13/viper"

type Config struct {
	CorsAllowOrigins  []string `mapstructure:"cors_allow_origins"`
	CorsAllowMethods  []string `mapstructure:"cors_allow_methods"`
	CorsAllowHeaders  []string `mapstructure:"cors_allow_headers"`
	CorsExposeHeaders []string `mapstructure:"cors_expose_headers"`
	// TODO:
	// log level
	// db adapter (sqlite, mysql)
}

func NewConfigFromEnv() (*Config, error) {
	viper.AutomaticEnv()
	setDefaults()

	var c Config
	if err := viper.Unmarshal(&c); err != nil {
		return nil, err
	}
	return &c, nil
}

func setDefaults() {
	viper.SetDefault("cors_allow_origins", []string{})
	viper.SetDefault("cors_allow_methods", []string{})
	viper.SetDefault("cors_allow_headers", []string{})
	viper.SetDefault("cors_expose_headers", []string{})
}
