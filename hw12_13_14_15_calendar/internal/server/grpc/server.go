package internalgrpc

import (
	"github.com/sergeylunev/otus-hw/hw12_13_14_15_calendar/api/grpcpb"
	"github.com/sergeylunev/otus-hw/hw12_13_14_15_calendar/internal/logger"
	"google.golang.org/grpc"
)

type GrpcServer struct {
	logger logger.Logger
	server *grpc.Server
}

func New(logger logger.Logger) *GrpcServer {
	server := &GrpcServer{
		logger: logger,
	}

	opt := make([]grpc.ServerOption, 0)
	opt = append(opt, grpc.UnaryInterceptor(server.loggingInterceptor))

	grpcServer := grpc.NewServer(opt...)
	grpcpb.Register

	return server
}
