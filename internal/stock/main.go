package main

import (
	"log"

	"github.com/hw1513/ecomm/common/config"
	"github.com/hw1513/ecomm/common/server"
	"google.golang.org/grpc"
)

var cfg *config.Config

func init() {
	var err error
	if cfg, err = config.LoadConfig("../../internal/common/config/config.dev.yaml"); err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}
}

func main() {
	server.RunGRPCServer(&cfg.Stock, func(*grpc.Server) {

	})
}
