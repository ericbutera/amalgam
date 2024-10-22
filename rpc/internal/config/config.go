package config

import (
	"github.com/spf13/viper"
)

type Config struct {
	Port string `mapstructure:"port"`
}

func init() {
	viper.SetDefault("port", "8080")
}
