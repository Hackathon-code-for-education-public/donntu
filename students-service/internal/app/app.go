package app

import (
	"log/slog"
	"students-service/internal/config"
)

type App struct {
	cfg *config.Config
	log *slog.Logger

	impl *srv.Handler
}
