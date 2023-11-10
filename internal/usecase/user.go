package usecase

import (
	"context"
	"github.com/Enthreeka/auth-microservice/internal/apperror"
	"github.com/Enthreeka/auth-microservice/internal/entity"
	"github.com/Enthreeka/auth-microservice/internal/repo"
	"github.com/Enthreeka/auth-microservice/pkg/logger"
)

type userService struct {
	postgres repo.UserRepository
	redis    repo.UserRepository
	log      *logger.Logger
}

func NewUserService(pg repo.UserRepository, redis repo.UserRepository, log *logger.Logger) User {
	return &userService{
		postgres: pg,
		redis:    redis,
		log:      log,
	}
}

func (u *userService) Create(ctx context.Context, user *entity.User) (*entity.User, error) {
	u.log.Info("Сreating a user")

	argon := NewArgonPassword(user.ID.String())
	hashPassword, err := argon.generateHashFromPassword(user.Password)
	if err != nil {
		return nil, err
	}
	user.Password = hashPassword

	createdUser, err := u.postgres.CreateUser(ctx, user)
	if err != nil {
		//return nil, apperror.ErrUserExist
		return nil, err
	}

	_, err = u.redis.CreateUser(ctx, user)
	if err != nil {
		//return nil, apperror.ErrUserExist
		return nil, err
	}

	return createdUser, nil
}

func (u *userService) Get(ctx context.Context, id string) (*entity.User, error) {
	userRedis, errRDS := u.redis.GetUserByID(ctx, id)
	//TODO Create validation to id

	u.log.Info("Starting search user in Redis with [id:%v]", id)
	if errRDS != nil && userRedis == nil {

		u.log.Error("User not exist in Redis with [id:%s]", id)
		userPostgres, errPG := u.postgres.GetUserByID(ctx, id)
		if errPG != nil && userPostgres == nil {
			u.log.Error("User not exist in Postgres with [id:%s]", id)
			return nil, apperror.ErrUserNotExist
		}

		_, err := u.redis.CreateUser(ctx, userPostgres)
		if err != nil {
			return nil, err
		}
		u.log.Info("Postgres has user with [id:%s]", id)
		return userPostgres, nil
	}

	u.log.Info("Redis has user with [id:%s]", id)
	return userRedis, nil
}

func (u *userService) Delete(ctx context.Context, id string) error {
	err := u.postgres.DeleteUserByID(ctx, id)
	if err != nil {
		return err
	}

	err = u.redis.DeleteUserByID(ctx, id)
	if err != nil {
		return err
	}

	return nil
}
