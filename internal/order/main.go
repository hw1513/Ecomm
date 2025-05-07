package main

import (
	"log"
	"net/http"

	"github.com/hw1513/ecomm/common/config"
)

var cfg *config.Config

func init() {
	var err error
	if cfg, err = config.LoadConfig("../../internal/common/config/config.dev.yaml"); err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}
}

func main() {

	// 使用配置的服务名称和地址
	log.Printf("Starting %s on %s", cfg.Order.ServiceName, cfg.Order.HTTPAddr)
	mux := http.NewServeMux()

	// 使用配置的HTTP地址
	if err := http.ListenAndServe(cfg.Order.HTTPAddr, mux); err != nil {
		log.Fatal(err)
	}
}
