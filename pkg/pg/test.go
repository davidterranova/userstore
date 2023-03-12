package pg

import (
	"sync"
	"testing"

	"github.com/davidterranova/userstore/pkg/config"
	"github.com/jmoiron/sqlx"
	"github.com/kelseyhightower/envconfig"
	"github.com/stretchr/testify/require"
)

var (
	once                = sync.Once{}
	dbInstance *sqlx.DB = nil
)

func TestConnection(t *testing.T) *sqlx.DB {
	t.Helper()

	once.Do(func() {
		var cfg config.ServerConfig
		err := envconfig.Process(cfg.EnvPrefix(), &cfg)
		require.NoError(t, err)

		dbInstance, err = Open(Config{
			ConnString:         cfg.DB.ConnString,
			MaxOpenConnections: cfg.DB.MaxOpenConnections,
			MaxConnIdleTime:    cfg.DB.MaxConnIdleTime,
			MaxConnLifetime:    cfg.DB.MaxConnLifetime,
		})
		require.NoError(t, err)
	})

	return dbInstance
}
