//go:build integration

package pg_test

import (
	"testing"

	"github.com/davidterranova/userstore/pkg/config"
	"github.com/davidterranova/userstore/pkg/pg"
	"github.com/kelseyhightower/envconfig"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestOpen(t *testing.T) {
	t.Run("should connect to existing db", func(t *testing.T) {
		var cfg config.ServerConfig
		err := envconfig.Process(cfg.EnvPrefix(), &cfg)
		require.NoError(t, err)

		db, err := pg.Open(pg.Config{
			ConnString:         cfg.DB.ConnString,
			MaxOpenConnections: cfg.DB.MaxOpenConnections,
			MaxConnIdleTime:    cfg.DB.MaxConnIdleTime,
			MaxConnLifetime:    cfg.DB.MaxConnLifetime,
		})
		require.NoError(t, err)
		assert.NotNil(t, db)
	})

	t.Run("should not connect with empty url", func(t *testing.T) {
		_, err := pg.Open(pg.Config{})
		require.Error(t, err)
		require.Contains(t, err.Error(), "dial error")
	})
}
