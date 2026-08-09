// Package db
package db

import (
	"context"
	"fmt"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
)

type Config struct {
	DSN          string
	MaxIdleTime  string
	MaxIdleConns int
	MinIdleConns int
}

func StartConnection(c *Config) (*pgxpool.Pool, error) {
	pc, err := pgxpool.ParseConfig(c.DSN)
	if err != nil {
		return nil, err
	}

	duration, err := time.ParseDuration(c.MaxIdleTime)
	if err != nil {
		return nil, err
	}

	pc.MaxConnIdleTime = duration
	pc.MaxConns = int32(c.MaxIdleConns)
	pc.MinIdleConns = int32(c.MinIdleConns)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	pool, err := pgxpool.NewWithConfig(ctx, pc)
	if err != nil {
		return nil, err
	}
	if err := pool.Ping(ctx); err != nil {
		return nil, fmt.Errorf("database unreachable %w", err)
	}

	return pool, nil
}
