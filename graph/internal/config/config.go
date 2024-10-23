package config

import "github.com/spf13/viper"

// GraphQL server configuration
type Config struct {
	Port      string `mapstructure:"port"`
	ApiHost   string `mapstructure:"api_host"`
	ApiScheme string `mapstructure:"api_scheme"`
	RpcHost   string `mapstructure:"rpc_host"`

	// Deprecated: use api_host instead
	ApiBaseUrl string `mapstructure:"api_base_url"`
}

func init() {
	viper.SetDefault("port", "8080")
	viper.SetDefault("api_host", "api:8080")
	viper.SetDefault("api_scheme", "https")
	viper.SetDefault("api_base_url", "api:8080/v1")
	viper.SetDefault("rpc_host", "rpc:50051")
}
