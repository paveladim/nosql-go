package server

import (
	"fmt"

	"go.uber.org/config"
)

type Config struct {
	Port string `yaml:"port"`
}

func NewServerConfig(provider config.Provider) (*Config, error) {
	var config Config
	err := provider.Get("server").Populate(&config)

	if err != nil {
		return nil, fmt.Errorf("server config: %w", err)
	}

	return &config, nil
}
