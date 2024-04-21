//go:build wireinject
// +build wireinject

package app

import (
	"fmt"
	"github.com/google/wire"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"log/slog"
	"os"
	//"verification-service/internal/config"
	//"verification-service/internal/handlers/grpc"
	//"verification-service/internal/services"
	//"verification-service/internal/storage/api"
	//"verification-service/internal/storage/pg"
)

func Init() (*App, func(), error) {
	panic(
		wire.Build(
			newApp,
			wire.NewSet(config.New),
			wire.NewSet(initLogger),
			wire.NewSet(initDB),
			//
			//wire.NewSet(pg.NewReasonStorage),
			//wire.NewSet(pg.NewRequestStorage),
			//wire.NewSet(api.NewAuthService),
			//wire.NewSet(services.New),
			//
			//wire.Bind(new(services.ReasonStorage), new(*pg.ReasonStorage)),
			//wire.Bind(new(services.RequestStorage), new(*pg.RequestStorage)),
			//wire.Bind(new(services.RoleUpdater), new(*api.AuthService)),
			//
			//wire.Bind(new(grpc.Service), new(*services.Service)),
			//wire.NewSet(grpc.New),
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
