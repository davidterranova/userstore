package pg

import (
	"time"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/stdlib"
	"github.com/jmoiron/sqlx"
)

type Config struct {
	ConnString         string
	MaxOpenConnections int
	MaxConnIdleTime    time.Duration
	MaxConnLifetime    time.Duration
}

func Open(cfg Config) (*sqlx.DB, error) {
	connConfig, err := pgx.ParseConfig(cfg.ConnString)
	if err != nil {
		return nil, err
	}
	connConfig.ConnectTimeout = 10 * time.Second

	db := sqlx.NewDb(stdlib.OpenDB(*connConfig), "pgx")
	db.SetConnMaxIdleTime(cfg.MaxConnIdleTime)
	db.SetConnMaxLifetime(cfg.MaxConnLifetime)
	db.SetMaxOpenConns(cfg.MaxOpenConnections)

	if err := db.Ping(); err != nil {
		return nil, err
	}

	return db, err
}
