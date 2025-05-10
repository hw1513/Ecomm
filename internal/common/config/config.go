package config

import (
	"fmt"

	"github.com/spf13/viper"
)

// Config 应用配置结构体
type Config struct {
	Order            ServiceConfig `mapstructure:"order"`
	Kitchen          ServiceConfig `mapstructure:"kitchen"`
	Payment          ServiceConfig `mapstructure:"payment"`
	Stock            ServiceConfig `mapstructure:"stock"`
	FallbackGRPCAddr string        `mapstructure:"fallback-grpc-addr"`
}

// ServiceConfig 服务配置
type ServiceConfig struct {
	ServiceName string `mapstructure:"service_name"`
	ServerToRun string `mapstructure:"server_to_run"`
	HTTPAddr    string `mapstructure:"http_addr"`
	GRPCAddr    string `mapstructure:"grpc_addr"`
}

// LoadConfig 加载配置文件
func LoadConfig(configPath string) (*Config, error) {
	v := viper.New()

	v.SetConfigName("global")
	v.SetConfigType("yaml")
	// 设置配置文件路径
	v.SetConfigFile(configPath)
	v.AutomaticEnv()

	// 读取配置文件
	if err := v.ReadInConfig(); err != nil {
		return nil, fmt.Errorf("failed to read config file: %w", err)
	}

	// 解析配置到结构体
	var config Config
	if err := v.Unmarshal(&config); err != nil {
		return nil, fmt.Errorf("failed to unmarshal config: %w", err)
	}

	return &config, nil
}
