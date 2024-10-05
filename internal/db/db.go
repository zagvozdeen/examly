package db

import (
	"context"
	"github.com/jackc/pgx/v5/pgxpool"
)

func New(ctx context.Context, addr string) (*pgxpool.Pool, error) {
	conn, err := pgxpool.New(ctx, addr)
	if err != nil {
		return nil, err
	}

	err = conn.Ping(ctx)
	if err != nil {
		return nil, err
	}

	return conn, nil
}
