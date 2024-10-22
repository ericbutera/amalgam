package config

import "github.com/spf13/viper"

// NewFromEnv unmarshals the environment into a struct.
// Make sure to include an init() function to set defaults.
func NewFromEnv[T any]() (*T, error) {
	viper.AutomaticEnv()
	var c T
	if err := viper.Unmarshal(&c); err != nil {
		return nil, err
	}
	return &c, nil
}
