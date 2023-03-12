package config

import (
	"time"

	"github.com/davidterranova/userstore/pkg/pg"
)

type ServerConfig struct {
	DB DBConfig `envconfig:"DB"`
}

func (cfg ServerConfig) EnvPrefix() string {
	return "USERSTORE"
}

type DBConfig struct {
	ConnString         string        `envconfig:"CONN_STRING" required:"true"`
	MaxOpenConnections int           `envconfig:"MAX_OPEN_CONNECTIONS" default:"25"`
	MaxConnIdleTime    time.Duration `envconfig:"MAX_CONN_IDLE_TIME" default:"1m"`
	MaxConnLifetime    time.Duration `envconfig:"MAX_CONN_LIFETIME" default:"5m"`
}

func (cfg DBConfig) DBConfig() pg.Config {
	return pg.Config{
		ConnString:         cfg.ConnString,
		MaxOpenConnections: cfg.MaxOpenConnections,
		MaxConnIdleTime:    cfg.MaxConnIdleTime,
		MaxConnLifetime:    cfg.MaxConnLifetime,
	}
}
