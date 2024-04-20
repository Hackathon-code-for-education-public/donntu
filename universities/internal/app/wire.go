//go:build wireinject
// +build wireinject

package app

import (
	"fmt"
	"github.com/google/wire"
	"google.golang.org/grpc"
	"log/slog"
	"universities/internal/app/router"
	"universities/internal/config"
	"universities/internal/infras/pgsql/repo"
	"universities/internal/services"
	"universities/pkg/engine"
)

func InitApp(grpcServer *grpc.Server, log *slog.Logger) (*Application, func(), error) {
	panic(wire.Build(NewApplication,
		wire.NewSet(config.New),
		initDB,
		wire.NewSet(router.NewGRPCServer),
		wire.NewSet(repo.NewUniversitiesPgRepo),
		wire.NewSet(services.NewUniversityService),
	))
}

func initDB(cfg *config.Config) (engine.DBEngine, func(), error) {
	connStr := engine.DBConnString(fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable", cfg.Database.User, cfg.Database.Pass, cfg.Database.Host, cfg.Database.Port, cfg.Database.Name))
	fmt.Println(connStr)
	db, err := engine.NewPostgresDB(connStr, 5, 2)
	if err != nil {
		return nil, nil, err
	}

	return db, func() { db.Close() }, nil
}
