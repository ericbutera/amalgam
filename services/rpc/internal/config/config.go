package config

import (
	"github.com/spf13/viper"
)

type Config struct {
	OtelEnable       bool     `mapstructure:"otel_enable"`
	IgnoredSpanNames []string `mapstructure:"ignored_span_names"`
	Port             string   `mapstructure:"port"`           // grpc server port
	MetricAddress    string   `mapstructure:"metric_address"` // metric server address
}

func init() { //nolint:gochecknoinits
	viper.SetDefault("otel_enable", false)
	viper.SetDefault("ignored_span_names", []string{"grpc.health.v1.Health/Check"})
	viper.SetDefault("port", "8080")
	viper.SetDefault("metric_address", ":9090")
}
