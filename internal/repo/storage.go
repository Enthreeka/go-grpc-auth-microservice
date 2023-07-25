package repo

import (
	"context"
	"github.com/NASandGAP/auth-microservice/internal/entity"
)

type Repository interface {
	GetUserByID(ctx context.Context, id string) (*entity.User, error)
	DeleteUserByID(ctx context.Context, id string) error
	CreateUser(ctx context.Context, user *entity.User) (*entity.User, error)
}
