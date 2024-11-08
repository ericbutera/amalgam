package env

import (
	"fmt"

	"github.com/spf13/viper"
)

// NewFromEnv unmarshals the environment into a struct.
// Make sure to include an init() function to set defaults.
func New[T any]() (*T, error) {
	viper.AutomaticEnv()

	var value T

	if err := viper.Unmarshal(&value); err != nil {
		return nil, fmt.Errorf("unable to unmarshal config: %w", err)
	}
	return &value, nil
}
