package usecase

import (
	"context"
	"github.com/NASandGAP/auth-microservice/internal/entity"
	"github.com/NASandGAP/auth-microservice/internal/repo"
)

type tokenService struct {
	postgres repo.TokenRepository
	redis    repo.TokenRepository
}

func NewTokenService(postgres repo.TokenRepository, redis repo.TokenRepository) Token {
	return &tokenService{
		postgres: postgres,
		redis:    redis,
	}
}

func (t *tokenService) Create(ctx context.Context, token *entity.Token) error {
	//TODO implement me
	panic("implement me")
}

func (t *tokenService) Get(ctx context.Context, id string) (*entity.Token, error) {
	//TODO implement me
	panic("implement me")
}

func (t *tokenService) Delete(ctx context.Context, id string) error {
	//TODO implement me
	panic("implement me")
}
