package postgres

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
)

type Postgres struct {
	Pool *pgxpool.Pool
}

func (p *Postgres) Close() {
	if p.Pool != nil {
		p.Pool.Close()
	}

}

func New(ctx context.Context, url string) (*Postgres, error) {

	pool, err := pgxpool.New(ctx, url)
	if err != nil {
		return nil, err
	}

	err = pool.Ping(ctx)
	if err != nil {
		return nil, err
	}

	db := &Postgres{
		Pool: pool,
	}

	return db, nil
}
