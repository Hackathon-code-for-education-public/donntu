package app

import (
	"log/slog"
	"universities/api/universities"
	"universities/internal/config"
	"universities/internal/services"
	"universities/pkg/engine"
)

type Application struct {
	Cfg *config.Config
	PG  engine.DBEngine
	Log *slog.Logger

	UniversityService      services.UniversityService
	UniversitiesGRPCServer universities.UniversitiesServer
}

func NewApplication(cfg *config.Config, log *slog.Logger, pg engine.DBEngine, universityService services.UniversityService, universitiesGRPCServer universities.UniversitiesServer) *Application {
	return &Application{
		Cfg:                    cfg,
		PG:                     pg,
		UniversityService:      universityService,
		UniversitiesGRPCServer: universitiesGRPCServer,
	}
}
