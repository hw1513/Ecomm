package server

import (
	"net"
	"os"

	grpc_logrus "github.com/grpc-ecosystem/go-grpc-middleware/logging/logrus"
	grpc_tags "github.com/grpc-ecosystem/go-grpc-middleware/tags"
	"github.com/hw1513/ecomm/common/config"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
)

func RunGRPCServer(config *config.ServiceConfig, registerServer func(server *grpc.Server)) {
	addr := config.GRPCAddr

	if addr == "" {
		//warn
	}

	// 初始化 logrus
	logger := logrus.New()
	logger.SetOutput(os.Stdout)
	logger.SetLevel(logrus.InfoLevel)
	entry := logrus.NewEntry(logger)

	lis, err := net.Listen("tcp", addr) // Change port as needed
	if err != nil {
		logger.Fatalf("failed to listen: %v", err)
	}

	// 创建带中间件的 gRPC 服务器
	grpcServer := grpc.NewServer(
		grpc.ChainUnaryInterceptor(
			grpc_tags.UnaryServerInterceptor(),
			grpc_logrus.UnaryServerInterceptor(entry),
		),
		grpc.ChainStreamInterceptor(
			grpc_tags.StreamServerInterceptor(),
			grpc_logrus.StreamServerInterceptor(entry),
		),
	)

	registerServer(grpcServer)
	// 3. Register your service implementation with the gRPC server
	// pb.RegisterYourServiceServer(grpcServer, &service.YourServiceImpl{})

	// 4. Start serving
	if err := grpcServer.Serve(lis); err != nil {
		logger.Fatalf("failed to serve: %v", err)
	}
}
