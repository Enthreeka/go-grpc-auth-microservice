package main

import (
	"github.com/Enthreeka/auth-microservice/internal/config"
	"github.com/Enthreeka/auth-microservice/internal/server"
	"github.com/Enthreeka/auth-microservice/pkg/logger"
)

func main() {
	path := `configs/config.json`

	log := logger.New()

	cfg, err := config.New(path)
	if err != nil {
		log.Fatal("failed to load config: %v", err)
	}

	if err := server.Run(cfg, log); err != nil {
		log.Fatal("failed to run server: %v", err)
	}
}
