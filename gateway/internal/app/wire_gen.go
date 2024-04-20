// Code generated by Wire. DO NOT EDIT.

//go:generate go run -mod=mod github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package app

import (
	"gateway/internal/config"
	"gateway/internal/controllers"
	"gateway/internal/logger"
	"gateway/internal/services"
)

// Injectors from wire.go:

func InitApp() *Application {
	configConfig := config.New()
	universityService := services.NewUniversityService(configConfig)
	slogLogger := logger.New()
	universitiesController := controllers.NewUniversityController(universityService, slogLogger)
	authService := services.NewAuthService(configConfig)
	authController := controllers.NewAuthController(authService, slogLogger)
	application := NewApplication(configConfig, universitiesController, authController, slogLogger)
	return application
}
