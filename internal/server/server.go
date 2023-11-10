package server

import (
	"context"
	"fmt"
	"github.com/Enthreeka/auth-microservice/internal/config"
	pb "github.com/Enthreeka/auth-microservice/internal/delivery/v1/grpc"
	"github.com/Enthreeka/auth-microservice/internal/interceptor"
	"github.com/Enthreeka/auth-microservice/pkg/logger"
	"github.com/Enthreeka/auth-microservice/pkg/postgres"
	"github.com/Enthreeka/auth-microservice/pkg/redis"
	"google.golang.org/grpc"
	"net"
)

func Run(cfg *config.Config, log *logger.Logger) error {
	// Connect to PostgreSQL
	psql, err := postgres.New(context.Background(), cfg.Postgres.URL)
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

	lis, err := net.Listen("tcp", fmt.Sprintf(":%s", cfg.ServerGrpc.Port))
	if err != nil {
		log.Fatal("failed to listen: %v", err)
		return err
	}

	s := grpc.NewServer(
		grpc.UnaryInterceptor(interceptor.UnaryInterceptor),
	)
	authServiceServer := pb.NewAuthServerGRPC(log, nil, nil)

	pb.RegisterAuthServiceServer(s, authServiceServer)

	log.Info("Starting gRPC listener on port :" + cfg.ServerGrpc.Port)
	if err := s.Serve(lis); err != nil {
		log.Fatal("failed to serve: %v", err)
		return err
	}

	return nil
}
