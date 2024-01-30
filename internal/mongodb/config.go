package mongodb

import (
	"fmt"

	"go.uber.org/config"
)

type Config struct {
	Uri string `yaml:"URI"`
}

func NewMongoConfig(provider config.Provider) (*Config, error) {
	var config Config
	err := provider.Get("mongodb").Populate(&config)

	if err != nil {
		return nil, fmt.Errorf("mongodb config: %w", err)
	}

	return &config, nil
}
