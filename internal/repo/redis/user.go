package redis

import (
	"context"
	"encoding/json"

	"github.com/NASandGAP/auth-microservice/internal/entity"
	"github.com/NASandGAP/auth-microservice/pkg/logger"
	"github.com/redis/go-redis/v9"
)

type userRedisRepo struct {
	redisClient *redis.Client
	log         *logger.Logger
}

func NewUserRedisRepo(redisClient *redis.Client, log *logger.Logger) Repository {
	return &userRedisRepo{
		redisClient: redisClient,
		log:         log,
	}
}

func (u *userRedisRepo) GetUserByID(ctx context.Context, id string) (*entity.User, error) {
	user := new(entity.User)

	userBytes, err := u.redisClient.Get(ctx, id).Bytes()
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(userBytes, user)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (u *userRedisRepo) DeleteUserByID(ctx context.Context, id string) error {
	err := u.redisClient.Del(ctx, id).Err()
	if err != nil {
		return err
	}

	return nil
}

func (u *userRedisRepo) CreateUser(ctx context.Context, user entity.User) error {
	bytesUser, err := json.Marshal(user)
	if err != nil {
		return err
	}

	err = u.redisClient.Set(ctx, user.ID, bytesUser, 5).Err()
	if err != nil {
		return err
	}
	return nil
}
