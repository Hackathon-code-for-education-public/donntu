//go:build wireinject
// +build wireinject

package app

import (
	"chat-service/internal/config"
	"chat-service/internal/handlers/grpc"
	"chat-service/internal/service"
	"chat-service/internal/storage/pg"
	redis2 "chat-service/internal/storage/redis"
	"context"
	"fmt"
	"github.com/google/wire"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/redis/go-redis/v9"
	"log/slog"
	"os"
	"time"
)

func Init() (*App, func(), error) {
	panic(
		wire.Build(
			newApp,
			wire.NewSet(config.New),
			wire.NewSet(initLogger),
			wire.NewSet(initDB),
			wire.NewSet(initRedis),

			wire.NewSet(service.New),

			wire.Bind(new(grpc.Service), new(*service.Service)),

			wire.NewSet(pg.NewChatStorage),
			wire.NewSet(pg.NewMessageStorage),
			wire.NewSet(redis2.NewBroker),

			wire.Bind(new(service.ChatStorage), new(*pg.ChatStorage)),
			wire.Bind(new(service.MessageStorage), new(*pg.MessageStorage)),
			wire.Bind(new(service.Broker), new(*redis2.Broker)),

			wire.NewSet(grpc.NewHandler),
		))
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

func initRedis(cfg *config.Config, log *slog.Logger) (*redis.Client, func(), error) {
	host := cfg.Redis.Host
	port := cfg.Redis.Port
	pass := cfg.Redis.Pass
	db := cfg.Redis.DB

	client := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%d", host, port),
		Password: pass,
		DB:       db,
	})

	ctx, _ := context.WithTimeout(context.Background(), 30*time.Second)

	log.Info("connecting to redis", slog.Int("db", db), slog.String("host", host), slog.Int("port", port))

	if _, err := client.Ping(ctx).Result(); err != nil {
		log.Error("failed to connect to redis", slog.String("err", err.Error()), slog.Int("db", db), slog.String("host", host), slog.Int("port", port))
		return nil, func() { client.Close() }, err
	}

	log.Info("connected to redis", slog.Int("db", db), slog.String("host", host), slog.Int("port", port))

	return client, func() { client.Close() }, nil
}
