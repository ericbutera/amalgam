package config

import (
	"github.com/spf13/viper"
)

type Config struct {
	Port          string `mapstructure:"port"`
	MetricAddress string `mapstructure:"metric_address"`
}

func init() {
	viper.SetDefault("port", "8080")
	viper.SetDefault("metric_address", ":9090")
}
