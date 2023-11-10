package server

import (
	"context"
	"fmt"
	"github.com/NASandGAP/auth-microservice/internal/config"
	pb "github.com/NASandGAP/auth-microservice/internal/delivery/grpc"
	"github.com/NASandGAP/auth-microservice/internal/entity"
	redisRepo "github.com/NASandGAP/auth-microservice/internal/repo/redis"
	"github.com/NASandGAP/auth-microservice/pkg/logger"
	"github.com/NASandGAP/auth-microservice/pkg/redis"
	"github.com/NASandGAP/auth-microservice/pkg/relationDB"
	"github.com/google/uuid"
	"google.golang.org/grpc"
	"net"
	"time"
)

type server struct {
	pb.UnimplementedAuthServiceServer
}

func Run(cfg *config.Config, log *logger.Logger) error {

	// Connect to PostgreSQL
	psql, err := relationDB.New(context.Background(), cfg.Postgres.URL)
	if err != nil {
		log.Fatal("failed to connect PostgreSQL: %v", err)
	}
	defer psql.Close()

	// Connect to Redis
	rds, err := redis.New(context.Background(), cfg)
	if err != nil {
		log.Error("redis is not working: %v", err)
	}
	defer rds.Close()
	////////////////////////TEST////////////////////////////////////////////////

	//userRepoPG := postgres.NewUserPostgresRepo(psql.Pool, log)
	//userRepoRDS := redisRepo.NewUserRedisRepo(rds, log)
	//
	//userService := usecase.NewUserService(userRepoPG, userRepoRDS, log)

	//tokenRepoPG := postgres.NewTokenPostgresRepo(psql.Pool)
	//
	//tokEN := &entity.Token{
	//	UserID:       "000be075-062d-4721-b208-0516b961ac56",
	//	RefreshToken: "000be075-062d-4721-b208-0516b961ac56",
	//	ExpiresAt:    time.Now().Add(120 * time.Second),
	//}
	//
	//createdToken, err := tokenRepoPG.CreateToken(context.Background(), tokEN)
	//if err != nil {
	//	log.Error("%v", err)
	//}
	//log.Info("%#v", createdToken)
	//
	//newToken, err := tokenRepoPG.GetTokenByID(context.Background(), createdToken.ID)
	//if err != nil {
	//	log.Error("%v", err)
	//}
	//log.Info("%#v", newToken)
	//
	//err = tokenRepoPG.DeleteTokenByID(context.Background(), createdToken.ID)
	//if err != nil {
	//	log.Error("%v", err)
	//}
	//log.Info("Token was deleted succesfuly!")

	userRepoRDS := redisRepo.NewUserRedisRepo(rds, log)

	uuid := uuid.New()
	usr := &entity.User{
		ID:       uuid,
		Email:    "testik@mail.sos",
		Role:     "user",
		Password: "testikebanniy",
	}

	_, err = userRepoRDS.CreateUser(context.Background(), usr)
	if err != nil {
		log.Error("%v", err)
	}

	tokenRepoRDS := redisRepo.NewTokenRedisRepo(rds)
	tokEN := &entity.Token{
		UserID:       uuid,
		RefreshToken: "000be075-062d-4721-b208-0516b961ac56",
		ExpiresAt:    time.Now().Add(120 * time.Second),
	}

	_, err = tokenRepoRDS.CreateToken(context.Background(), tokEN)
	if err != nil {
		log.Error("%v", err)
	}

	////////////////////////TEST////////////////////////////////////////////////
	lis, err := net.Listen("tcp", fmt.Sprintf(":%s", cfg.ServerGrpc.Port))
	if err != nil {
		log.Fatal("failed to listen: %v", err)
		return err
	}

	s := grpc.NewServer()
	pb.RegisterAuthServiceServer(s, &server{})

	log.Info("Starting gRPC listener on port :" + cfg.ServerGrpc.Port)
	if err := s.Serve(lis); err != nil {
		log.Fatal("failed to serve: %v", err)
		return err
	}

	return nil
}
