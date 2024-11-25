package config

import "github.com/spf13/viper"

const DefaultComplexityLimit = 15

type Config struct {
	OtelEnable        bool     `mapstructure:"otel_enable"`
	IgnoredSpanNames  []string `mapstructure:"ignored_span_names"`
	Port              string   `mapstructure:"graph_port"`
	ComplexityLimit   int      `mapstructure:"graph_complexity_limit"`
	RpcHost           string   `mapstructure:"rpc_host"`
	RpcInsecure       bool     `mapstructure:"rpc_insecure"`
	CorsAllowOrigins  []string `mapstructure:"cors_allow_origins"`
	CorsAllowMethods  []string `mapstructure:"cors_allow_methods"`
	CorsAllowHeaders  []string `mapstructure:"cors_allow_headers"`
	CorsExposeHeaders []string `mapstructure:"cors_expose_headers"`
}

func init() { //nolint:gochecknoinits
	viper.SetDefault("otel_enable", false)
	viper.SetDefault("ignored_span_names", []string{})
	viper.SetDefault("graph_port", "8080")
	viper.SetDefault("graph_complexity_limit", DefaultComplexityLimit)
	viper.SetDefault("rpc_host", "rpc:50051")
	viper.SetDefault("rpc_insecure", false)
	viper.SetDefault("cors_allow_origins", []string{})
	viper.SetDefault("cors_allow_methods", []string{})
	viper.SetDefault("cors_allow_headers", []string{})
	viper.SetDefault("cors_expose_headers", []string{})
}
