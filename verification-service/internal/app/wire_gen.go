// Code generated by Wire. DO NOT EDIT.

//go:generate go run -mod=mod github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package app

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"log/slog"
	"os"
	"verification-service/internal/config"
	"verification-service/internal/handlers/grpc"
	"verification-service/internal/services"
	"verification-service/internal/storage/api"
	"verification-service/internal/storage/pg"
)

import (
	_ "github.com/lib/pq"
)

// Injectors from wire.go:

func Init() (*App, func(), error) {
	configConfig := config.New()
	logger := initLogger(configConfig)
	db, cleanup, err := initDB(configConfig)
	if err != nil {
		return nil, nil, err
	}
	requestStorage := pg.NewRequestStorage(db)
	reasonStorage := pg.NewReasonStorage(db)
	authService := api.NewAuthService(configConfig)
	service := services.New(requestStorage, reasonStorage, authService)
	handler := grpc.New(service)
	app := newApp(configConfig, logger, handler)
	return app, func() {
		cleanup()
	}, nil
}

// wire.go:

func initLogger(cfg *config.Config) *slog.Logger {

	var level slog.Level

	switch cfg.Logger.Level {
	case "debug":
		level = slog.LevelDebug
	case "info":
		level = slog.LevelInfo
	}

	l := slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: level}))
	slog.SetDefault(l)
	return l
}

func initDB(cfg *config.Config) (*sqlx.DB, func(), error) {

	host := cfg.DB.Host
	port := cfg.DB.Port
	user := cfg.DB.User
	pass := cfg.DB.Pass
	name := cfg.DB.Name

	cs := fmt.Sprintf("postgres://%s:%s@%s:%d/%s?sslmode=disable", user, pass, host, port, name)
	slog.Info("connecting to database", slog.String("conn", cs))

	db, err := sqlx.Open("postgres", cs)
	if err != nil {
		return nil, nil, err
	}

	if err := db.Ping(); err != nil {
		slog.Error("failed to connect to database", slog.String("err", err.Error()), slog.String("conn", cs))
		return nil, func() { db.Close() }, err
	}
	slog.Info("connected to database", slog.String("conn", cs))

	return db, func() { db.Close() }, nil
}
