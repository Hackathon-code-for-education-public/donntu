package services

import (
	"context"
	"universities/internal/domain"
)

type (
	UniversityRepository interface {
		GetOpenDays(ctx context.Context, universityID string) ([]*domain.OpenDay, error)
	}

	UniversityService interface {
		GetOpenDays(ctx context.Context, universityID string) ([]*domain.OpenDay, error)
	}
)
