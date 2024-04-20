package interceptors

import (
	"google.golang.org/grpc"
	"log/slog"
)

func LoggingInterceptor(log *slog.Logger) grpc.StreamServerInterceptor {

	return func(srv any, ss grpc.ServerStream, info *grpc.StreamServerInfo, handler grpc.StreamHandler) error {
		log.Info("new grpc request", slog.String("method", info.FullMethod))
		return handler(srv, ss)
	}
}
