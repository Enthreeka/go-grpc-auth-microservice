package usecase

import (
	"github.com/NASandGAP/auth-microservice/internal/repo"
	"github.com/NASandGAP/auth-microservice/pkg/logger"
)

type userService struct {
	postgres repo.Repository
	redis    repo.Repository
	log      *logger.Logger
}

func NewUserService(pg repo.Repository, redis repo.Repository, log *logger.Logger) UserService {
	return &userService{
		postgres: pg,
		redis:    redis,
		log:      log,
	}
}

func (u userService) Create() {
	//TODO implement me
	panic("implement me")
}

func (u userService) Get() {
	//TODO implement me
	panic("implement me")
}

func (u userService) Delete() {
	//TODO implement me
	panic("implement me")
}
