package services

import (
	"context"
	"universities/internal/domain"
)

type (
	UniversityRepository interface {
		GetOpenDays(ctx context.Context, universityID string) ([]*domain.OpenDay, error)

		CreateReview(ctx context.Context, review *domain.Review) (*domain.Review, error)
		GetReplies(ctx context.Context, reviewID string) ([]*domain.Review, error)
		GetReviews(ctx context.Context, universityID string, limit, offset int) ([]*domain.Review, error)

		CreatePanorama(ctx context.Context, panorama *domain.Panorama) (*domain.Panorama, error)
		GetPanoramas(ctx context.Context, universityID string, category string) ([]*domain.Panorama, error)

		GetUniversitiesTop(ctx context.Context, n int) ([]*domain.University, error)
		GetUniversities(ctx context.Context, offset, limit int) ([]*domain.University, error)
		GetUniversity(ctx context.Context, universityID string) (*domain.University, error)
		SearchUniversities(ctx context.Context, name string) ([]*domain.University, error)
	}

	UniversityService interface {
		GetOpenDays(ctx context.Context, universityID string) ([]*domain.OpenDay, error)

		CreateReview(ctx context.Context, review *domain.Review) (*domain.Review, error)
		GetReplies(ctx context.Context, reviewID string) ([]*domain.Review, error)
		GetReviews(ctx context.Context, universityID string, limit, offset int) ([]*domain.Review, error)

		CreatePanorama(ctx context.Context, panorama *domain.Panorama) (*domain.Panorama, error)
		GetPanoramas(ctx context.Context, universityID string, category string) ([]*domain.Panorama, error)

		GetUniversitiesTop(ctx context.Context, n int) ([]*domain.University, error)
		GetUniversities(ctx context.Context, offset, limit int) ([]*domain.University, error)
		GetUniversity(ctx context.Context, universityID string) (*domain.University, error)
		SearchUniversities(ctx context.Context, name string) ([]*domain.University, error)
	}
)
