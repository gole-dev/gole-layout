package service

import (
	"context"
	"github.com/gole-dev/gole-layout/pb/hello"
)

var (
	_ hello.GreeterServer = (*GreeterService)(nil)
)

type GreeterService struct {
	hello.UnimplementedGreeterServer
}

func NewGreeterService() *GreeterService {
	return &GreeterService{}
}

func (s *GreeterService) SayHello(ctx context.Context, req *hello.HelloRequest) (*hello.HelloReply, error) {
	return &hello.HelloReply{}, nil
}
