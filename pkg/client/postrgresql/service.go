package postrgresql

import (
	"context"
	"github.com/jackc/pgx/v5/pgxpool"
	"time"
)

func NewClient(ctx context.Context, connectionURL string) (*pgxpool.Pool, error) {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	pool, err := pgxpool.New(ctx, connectionURL)
	if err != nil {
		return nil, err
	}

	return pool, nil
}
