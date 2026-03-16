package db

import (
	"auth/config"
	"context"
	"fmt"

	"github.com/jackc/pgx/v5/pgxpool"
)

func ConnectDB(cfg *config.Config) (*pgxpool.Pool, error) {

	pool, err := pgxpool.New(context.TODO(), cfg.DBURL)
	if err != nil {
		return nil, fmt.Errorf("unable to connect to database : %w", err)
	}

	if err := pool.Ping(context.Background()); err != nil {
		return nil, fmt.Errorf("database ping failed : %w", err)
	}

	return pool, nil
}
