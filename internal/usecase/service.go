package usecase

import (
	"context"
	"github.com/Enthreeka/auth-microservice/internal/entity"
)

type User interface {
	Create(ctx context.Context, user *entity.User) (*entity.User, error)
	Get(ctx context.Context, id string) (*entity.User, error)
	Delete(ctx context.Context, id string) error
}

type Token interface {
	Create(ctx context.Context, token *entity.Token) error
	Get(ctx context.Context, id string) (*entity.Token, error)
	Delete(ctx context.Context, id string) error
}

type Authentication interface {
}
