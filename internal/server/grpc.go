package server

import (
	"github.com/gole-dev/gole-layout/pb/hello"
	"time"

	"github.com/gole-dev/gole/pkg/app"
	"github.com/gole-dev/gole/pkg/transport/grpc"

	"github.com/gole-dev/gole-layout/internal/service"
)

// NewGRPCServer creates a gRPC server
func NewGRPCServer(cfg *app.ServerConfig) *grpc.Server {

	grpcServer := grpc.NewServer(
		grpc.Network("tcp"),
		grpc.Address(cfg.Addr),
		grpc.Timeout(3*time.Second),
	)

	// register biz service
	hello.RegisterGreeterServer(grpcServer, service.NewGreeterService())

	return grpcServer
}
