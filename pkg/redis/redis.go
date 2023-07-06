package redis

import (
	"context"

	"github.com/NASandGAP/auth-microservice/internal/config"
	"github.com/redis/go-redis/v9"
)

type Redis struct {
	rds *redis.Client
}

func (r *Redis) Close() {
	if r.rds != nil {
		r.rds.Close()
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
		rds: client,
	}

	return rds, nil
}
