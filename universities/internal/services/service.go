package services

import (
	"context"
	"universities/internal/domain"
)

type universityService struct {
	universityRepo UniversityRepository
}

func NewUniversityService(universityRepo UniversityRepository) UniversityService {
	return &universityService{
		universityRepo: universityRepo,
	}
}

func (s *universityService) GetOpenDays(ctx context.Context, universityID string) ([]*domain.OpenDay, error) {
	return s.universityRepo.GetOpenDays(ctx, universityID)
}

func (s *universityService) GetReviews(ctx context.Context, universityID string, limit, offset int) ([]*domain.Review, error) {
	return s.universityRepo.GetReviews(ctx, universityID, limit, offset)
}

func (s *universityService) CreatePanorama(ctx context.Context, panorama *domain.Panorama) (*domain.Panorama, error) {
	return s.universityRepo.CreatePanorama(ctx, panorama)
}

func (s *universityService) GetPanoramas(ctx context.Context, universityID string, category string) ([]*domain.Panorama, error) {
	return s.universityRepo.GetPanoramas(ctx, universityID, category)
}

func (s *universityService) GetUniversitiesTop(ctx context.Context, n int) ([]*domain.University, error) {
	return s.universityRepo.GetUniversitiesTop(ctx, n)
}

func (s *universityService) GetUniversities(ctx context.Context, offset, limit int) ([]*domain.University, error) {
	return s.universityRepo.GetUniversities(ctx, offset, limit)
}

func (s *universityService) GetUniversity(ctx context.Context, universityID string) (*domain.University, error) {
	return s.universityRepo.GetUniversity(ctx, universityID)
}

func (s *universityService) SearchUniversities(ctx context.Context, name string) ([]*domain.University, error) {
	return s.universityRepo.SearchUniversities(ctx, name)
}
