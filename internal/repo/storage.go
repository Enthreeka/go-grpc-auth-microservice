package repo

import (
	"context"
	"github.com/NASandGAP/auth-microservice/internal/entity"
)

type UserRepository interface {
	GetUserByID(ctx context.Context, id string) (*entity.User, error)
	DeleteUserByID(ctx context.Context, id string) error
	CreateUser(ctx context.Context, user *entity.User) (*entity.User, error)
}

type TokenRepository interface {
	GetTokenByID(ctx context.Context, id string) (*entity.Token, error)
	DeleteTokenByID(ctx context.Context, id string) error
	CreateToken(ctx context.Context, token *entity.Token) (*entity.Token, error)
}
