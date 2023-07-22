package server

import (
	"context"
	"fmt"
	"github.com/NASandGAP/auth-microservice/internal/config"
	pb "github.com/NASandGAP/auth-microservice/internal/delivery/grpc"
	"github.com/NASandGAP/auth-microservice/pkg/logger"
	"github.com/NASandGAP/auth-microservice/pkg/postgres"
	"github.com/NASandGAP/auth-microservice/pkg/redis"
	"google.golang.org/grpc"
	"net"
)

type server struct {
	pb.UnimplementedAuthServiceServer
}

func Run(cfg *config.Config, log *logger.Logger) error {

	// Connect to PostgreSQL
	psql, err := postgres.New(context.Background(), cfg.Postgres.URL)
	if err != nil {
		log.Fatal("failed to connect postgres: %v", err)
	}
	defer psql.Close()

	// Connect to Redis
	rds, err := redis.New(context.Background(), cfg)
	if err != nil {
		log.Fatal("redis is not working: %v", err)
	}
	defer rds.Close()

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
