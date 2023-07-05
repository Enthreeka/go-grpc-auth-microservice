package main

import (
	"github.com/NASandGAP/auth-microservice/internal/config"
	"github.com/NASandGAP/auth-microservice/internal/server"
	"github.com/NASandGAP/auth-microservice/pkg/logger"
)

func main() {
	path := `configs/config.json`

	log := logger.New()

	cfg, err := config.New(path)
	if err != nil {
		log.Error("failed to load config: %v", err)
	}

	if err := server.Run(cfg, log); err != nil {
		log.Fatal("failed to run server: %v", err)
	}
}
