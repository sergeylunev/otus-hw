package internalgrpc

import (
	"context"
	"fmt"
	"time"

	ltime "github.com/sergeylunev/otus-hw/hw12_13_14_15_calendar/internal/time"
	"google.golang.org/grpc"
	"google.golang.org/grpc/peer"
)

func (gs *GrpcServer) loggingInterceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
	startTime := time.Now()
	h, err := handler(ctx, req)
	latency := time.Since(startTime)

	var remoteAddr string
	if p, ok := peer.FromContext(ctx); ok {
		remoteAddr = p.Addr.String()
	}

	gs.logger.Info(
		fmt.Sprintf("%s [%s] %s %s %s \"%s\"",
			remoteAddr,
			time.Now().Format(ltime.DateTimeLogFormat),
			info.FullMethod,
			req,
			latency,
			"GRPC-client",
		),
	)

	return h, err
}
