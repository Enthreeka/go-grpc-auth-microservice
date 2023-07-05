package server

import (
	"fmt"
	"net"

	"github.com/NASandGAP/auth-microservice/internal/config"
	pb "github.com/NASandGAP/auth-microservice/internal/delivery/grpc"
	"github.com/NASandGAP/auth-microservice/pkg/logger"
	"google.golang.org/grpc"
)

type server struct {
	pb.UnimplementedAuthServiceServer
}

func Run(cfg *config.Config, log *logger.Logger) error {

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
