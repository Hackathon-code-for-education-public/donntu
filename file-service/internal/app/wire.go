//go:build wireinject
// +build wireinject

package app

import (
	"file-service/internal/config"
	"file-service/internal/handler/grpc"
	"file-service/internal/service"
	"file-service/internal/storage"
	"github.com/google/wire"
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
	"log/slog"
	"os"
)

func Init() (*App, func(), error) {
	panic(
		wire.Build(
			newApp,
			wire.NewSet(config.New),
			wire.NewSet(initLogger),
			wire.NewSet(initMinio),
			wire.NewSet(storage.New),

			wire.Bind(new(service.Storage), new(*storage.MinioStorage)),
			wire.NewSet(service.New),

			wire.Bind(new(grpc.Service), new(*service.Service)),
			wire.NewSet(grpc.New),
		),
	)
}

func initLogger(cfg *config.Config) *slog.Logger {

	var level slog.Level

	switch cfg.Logger.Level {
	case "debug":
		level = slog.LevelDebug
	case "info":
		level = slog.LevelInfo
	}

	l := slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: level}))
	slog.SetDefault(l)

	return l
}

func initMinio(cfg *config.Config) (*minio.Client, error) {
	endpoint := cfg.Minio.Endpoint
	accessKey := cfg.Minio.AccessKey
	secretKey := cfg.Minio.SecretKey

	client, err := minio.New(endpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(accessKey, secretKey, ""),
		Secure: cfg.Minio.Secure,
	})
	if err != nil {
		return nil, err
	}

	return client, err
}
