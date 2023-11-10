package redis

import (
	"context"
	"github.com/Enthreeka/auth-microservice/internal/config"
)

type Redis struct {
	Rds *redis.Client
}

func (r *Redis) Close() {
	if r.Rds != nil {
		r.Rds.Close()
	}
}

func New(ctx context.Context, cfg *config.Config) (*Redis, error) {

	client := redis.NewClient(&redis.Options{
		Addr:         cfg.Redis.Host,
		Password:     cfg.Redis.Password,
		MinIdleConns: cfg.Redis.MinIdleConns,
		DB:           cfg.Redis.Db,
	})

	err := client.Ping(ctx).Err()
	if err != nil {
		return nil, err
	}

	rds := &Redis{
		Rds: client,
	}

	return rds, nil
}
