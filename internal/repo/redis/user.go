package redis

import (
	"context"
	"encoding/json"
	"github.com/NASandGAP/auth-microservice/internal/entity"
	"github.com/NASandGAP/auth-microservice/internal/repo"
	"github.com/NASandGAP/auth-microservice/pkg/logger"
	pkg "github.com/NASandGAP/auth-microservice/pkg/redis"
	"sync"
	"time"
)

type userRedisRepo struct {
	*pkg.Redis
	*logger.Logger

	sync.Mutex
}

func NewUserRedisRepo(redis *pkg.Redis, log *logger.Logger) repo.UserRepository {
	return &userRedisRepo{
		Redis:  redis,
		Logger: log,
	}
}

func (u *userRedisRepo) GetUserByID(ctx context.Context, id string) (*entity.User, error) {
	u.Lock()
	defer u.Unlock()

	user := new(entity.User)

	userBytes, err := u.Rds.Get(ctx, id).Bytes()
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
	err := u.Rds.Del(ctx, id).Err()
	if err != nil {
		return err
	}

	return nil
}

func (u *userRedisRepo) CreateUser(ctx context.Context, user *entity.User) (*entity.User, error) {
	bytesUser, err := json.Marshal(user)
	if err != nil {
		return nil, err
	}

	err = u.Rds.Set(ctx, user.ID.String(), bytesUser, 360*time.Hour).Err()
	if err != nil {
		return nil, err
	}

	return nil, nil
}
