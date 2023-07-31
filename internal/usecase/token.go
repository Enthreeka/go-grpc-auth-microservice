package usecase

import (
	"context"
	"github.com/NASandGAP/auth-microservice/internal/entity"
)

type tokenService struct {
}

func NewTokenService() Token {
	return &tokenService{}
}

func (t tokenService) Create(ctx context.Context, token *entity.Token) error {
	//TODO implement me
	panic("implement me")
}

func (t tokenService) Get(ctx context.Context, id string) (*entity.Token, error) {
	//TODO implement me
	panic("implement me")
}

func (t tokenService) Delete(ctx context.Context, id string) error {
	//TODO implement me
	panic("implement me")
}
