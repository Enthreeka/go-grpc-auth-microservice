package server

import (
	"context"
	"fmt"
	"github.com/NASandGAP/auth-microservice/internal/config"
	pb "github.com/NASandGAP/auth-microservice/internal/delivery/grpc"
	"github.com/NASandGAP/auth-microservice/internal/repo/postgres"
	redisRepo "github.com/NASandGAP/auth-microservice/internal/repo/redis"
	"github.com/NASandGAP/auth-microservice/internal/usecase"
	"github.com/NASandGAP/auth-microservice/pkg/logger"
	"github.com/NASandGAP/auth-microservice/pkg/redis"
	"github.com/NASandGAP/auth-microservice/pkg/relationDB"
	"google.golang.org/grpc"
	"net"
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

	userRepoPG := postgres.NewUserPostgresRepo(psql.Pool, log)
	userRepoRDS := redisRepo.NewUserRedisRepo(rds, log)

	userService := usecase.NewUserService(userRepoPG, userRepoRDS, log)

	userService.Get(context.Background(), "1")
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
