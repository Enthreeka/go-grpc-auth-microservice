package grpc

import (
	"context"
	"github.com/Enthreeka/auth-microservice/internal/usecase"
	"github.com/Enthreeka/auth-microservice/pkg/logger"
)

type userHandlerGRPC struct {
	log          *logger.Logger
	userUsecase  usecase.User
	tokenUsecase usecase.Token

	UnimplementedAuthServiceServer
}

func NewAuthServerGRPC(log *logger.Logger, userUsecase usecase.User, tokenUsecase usecase.Token) AuthServiceServer {
	return &userHandlerGRPC{
		log:          log,
		userUsecase:  userUsecase,
		tokenUsecase: tokenUsecase,
	}
}

func (u *userHandlerGRPC) SignUp(context.Context, *SignUpRequest) (*SignUpResponse, error) {

	panic("error")
}

func (u *userHandlerGRPC) SignIn(context.Context, *SignInRequest) (*SignInResponse, error) {

	panic("error")
}

func (u *userHandlerGRPC) Logout(context.Context, *LogoutRequest) (*LogoutResponse, error) {

	panic("error")
}

func (u *userHandlerGRPC) FindByID(context.Context, *FindByIdRequest) (*FindByIdResponse, error) {

	panic("error")
}

func (u *userHandlerGRPC) FindByEmail(context.Context, *FindByEmailRequest) (*FindByEmailResponse, error) {

	panic("error")
}
