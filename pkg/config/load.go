package config

import (
	"fmt"

	"github.com/kelseyhightower/envconfig"
)

type PrefixedConfig interface {
	EnvPrefix() string
}

func LoadEnv(cfg PrefixedConfig) error {
	err := envconfig.Process(cfg.EnvPrefix(), cfg)
	if err != nil {
		return fmt.Errorf("failed to parse environment variables for type %T: %w", cfg, err)
	}
	return nil
}
