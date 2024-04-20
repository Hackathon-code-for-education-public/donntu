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
