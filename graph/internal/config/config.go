package config

import "github.com/spf13/viper"

// GraphQL server configuration
type Config struct {
	Port      string `mapstructure:"port"`
	ApiHost   string `mapstructure:"api_host"`
	ApiScheme string `mapstructure:"api_scheme"`

	// Deprecated: use api_host instead
	ApiBaseUrl string `mapstructure:"api_base_url"`
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
	viper.SetDefault("port", "8080")
	viper.SetDefault("api_host", "api:8080")
	viper.SetDefault("api_scheme", "https")
	viper.SetDefault("api_base_url", "api:8080/v1")
}
