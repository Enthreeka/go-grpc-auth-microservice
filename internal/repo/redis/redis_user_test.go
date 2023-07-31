package redis

import (
	"context"
	"github.com/NASandGAP/auth-microservice/internal/entity"
	"github.com/NASandGAP/auth-microservice/internal/repo"
	pkg "github.com/NASandGAP/auth-microservice/pkg/redis"
	"github.com/alicebob/miniredis/v2"
	"github.com/google/uuid"
	"github.com/redis/go-redis/v9"
	"github.com/stretchr/testify/require"
	"log"
	"testing"
)

func SetupRedis() repo.UserRepository {
	mr, err := miniredis.Run()
	if err != nil {
		log.Fatal(err)
	}

	client := redis.NewClient(&redis.Options{
		Addr: mr.Addr(),
	})

	rds := &pkg.Redis{
		Rds: client,
	}

	userRedisRepo := NewUserRedisRepo(rds, nil)

	return userRedisRepo
}

func TestUserRedisRepo_CreateUser(t *testing.T) {
	t.Parallel()

	userRedisRepo := SetupRedis()

	t.Run("Create User in Redis", func(t *testing.T) {

		UUID := uuid.New().String()

		userMock := &entity.User{
			ID:       UUID,
			Email:    "TestGmail",
			Password: "TESTPASSWORD",
		}

		_, err := userRedisRepo.CreateUser(context.Background(), userMock)
		require.NoError(t, err)
	})
}

func TestUserRedisRepo_GetUserByID(t *testing.T) {
	t.Parallel()

	userRedisRepo := SetupRedis()

	t.Run("Get user from Redis", func(t *testing.T) {

		UUID := uuid.New().String()
		userMock := &entity.User{
			ID:       UUID,
			Email:    "@FKNIZHNYNOVGOROD",
			Password: "TOP",
		}

		_, err := userRedisRepo.CreateUser(context.Background(), userMock)
		require.NoError(t, err)

		user, err := userRedisRepo.GetUserByID(context.Background(), userMock.ID)
		require.NoError(t, err)
		require.NotNil(t, user)
	})
}

func TestUserRedisRepo_DeleteUserByID(t *testing.T) {
	t.Parallel()

	userRedisRepo := SetupRedis()

	t.Run("Delete user from Redis", func(t *testing.T) {

		UUID := uuid.New().String()
		userMock := &entity.User{
			ID: UUID,
		}

		err := userRedisRepo.DeleteUserByID(context.Background(), userMock.ID)
		require.NoError(t, err)
	})
}
