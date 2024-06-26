//go:build wireinject
// +build wireinject

package app

import (
	"gateway/internal/config"
	"gateway/internal/controllers"
	"gateway/internal/logger"
	"gateway/internal/services"
	"github.com/google/wire"
)

func InitApp() *Application {
	panic(wire.Build(
		NewApplication,
		wire.NewSet(config.New),
		wire.NewSet(logger.New),
		wire.NewSet(services.NewAuthService),
		wire.NewSet(controllers.NewAuthController),
		wire.NewSet(services.NewUniversityService),
		wire.NewSet(services.NewFileService),
		wire.NewSet(controllers.NewUniversityController),
		wire.NewSet(controllers.NewChatController),
	))
}
