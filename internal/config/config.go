package config

import (
	"encoding/json"
	"os"
)

type (
	Config struct {
		ServerGrpc ServerGrpc `json:"grpc_server"`
		ClientGrpc ClientGrpc `json:"grpc_client"`
		Postgres   Postgres   `json:"postgres"`
		Redis      Redis      `json:"redis"`
		HTTPServer HTTPServer `json:"http_server"`
	}

	ServerGrpc struct {
		Port string `json:"port"`
	}

	ClientGrpc struct {
		Port string `json:"port"`
	}

	Postgres struct {
		URL string `json:"url"`
	}

	Redis struct {
		//		URL string `json:"url"`
	}

	HTTPServer struct {
		Hostname   string `json:"hostname"`
		Port       string `json:"port"`
		TypeServer string `json:"type_server"`
	}
)

func New(path string) (*Config, error) {

	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}

	defer file.Close()

	decoder := json.NewDecoder(file)
	config := &Config{}
	if err := decoder.Decode(config); err != nil {
		return nil, err
	}

	return config, nil
}
