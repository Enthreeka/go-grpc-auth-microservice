package usecase

import (
	"context"
	"github.com/NASandGAP/auth-microservice/internal/entity"
)

type UserService interface {
	Create(ctx context.Context, user *entity.User) (*entity.User, error)
	Get(ctx context.Context, id string) (*entity.User, error)
	Delete(ctx context.Context, id string) error
}
