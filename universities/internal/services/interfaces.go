package services

import (
	"context"
	"universities/internal/domain"
)

type (
	UniversityRepository interface {
		GetOpenDays(ctx context.Context, universityID string) ([]*domain.OpenDay, error)
		GetReviews(ctx context.Context, universityID string, limit, offset int) ([]*domain.Review, error)
	}

	UniversityService interface {
		GetOpenDays(ctx context.Context, universityID string) ([]*domain.OpenDay, error)
		GetReviews(ctx context.Context, universityID string, limit, offset int) ([]*domain.Review, error)
	}
)
