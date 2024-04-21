package services

import (
	"context"
	"universities/internal/domain"
)

type (
	UniversityRepository interface {
		GetOpenDays(ctx context.Context, universityID string) ([]*domain.OpenDay, error)
		GetReviews(ctx context.Context, universityID string, limit, offset int) ([]*domain.Review, error)
		CreatePanorama(ctx context.Context, panorama *domain.Panorama) (*domain.Panorama, error)
		GetPanoramas(ctx context.Context, universityID string, category string) ([]*domain.Panorama, error)
	}

	UniversityService interface {
		GetOpenDays(ctx context.Context, universityID string) ([]*domain.OpenDay, error)
		GetReviews(ctx context.Context, universityID string, limit, offset int) ([]*domain.Review, error)
		CreatePanorama(ctx context.Context, panorama *domain.Panorama) (*domain.Panorama, error)
		GetPanoramas(ctx context.Context, universityID string, category string) ([]*domain.Panorama, error)
	}
)
