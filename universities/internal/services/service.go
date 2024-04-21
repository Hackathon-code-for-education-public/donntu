package services

import (
	"context"
	"errors"
	"fmt"
	"github.com/google/uuid"
	"strings"
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
	universityID = strings.TrimSpace(universityID)
	return s.universityRepo.GetOpenDays(ctx, universityID)
}

func (s *universityService) GetReviews(ctx context.Context, universityID string, limit, offset int) ([]*domain.Review, error) {
	universityID = strings.TrimSpace(universityID)
	return s.universityRepo.GetReviews(ctx, universityID, limit, offset)
}

func (s *universityService) CreatePanorama(ctx context.Context, panorama *domain.Panorama) (*domain.Panorama, error) {
	return s.universityRepo.CreatePanorama(ctx, panorama)
}

func (s *universityService) GetPanoramas(ctx context.Context, universityID string, category string) ([]*domain.Panorama, error) {
	universityID = strings.TrimSpace(universityID)
	return s.universityRepo.GetPanoramas(ctx, universityID, category)
}

func (s *universityService) GetUniversitiesTop(ctx context.Context, n int) ([]*domain.University, error) {
	return s.universityRepo.GetUniversitiesTop(ctx, n)
}

func (s *universityService) GetUniversities(ctx context.Context, offset, limit int) ([]*domain.University, error) {
	return s.universityRepo.GetUniversities(ctx, offset, limit)
}

func (s *universityService) GetUniversity(ctx context.Context, universityID string) (*domain.University, error) {
	universityID = strings.TrimSpace(universityID)
	return s.universityRepo.GetUniversity(ctx, universityID)
}

func (s *universityService) SearchUniversities(ctx context.Context, name string) ([]*domain.University, error) {
	name = strings.TrimSpace(name)
	return s.universityRepo.SearchUniversities(ctx, name)
}

func (s *universityService) CreateReview(ctx context.Context, review *domain.Review) (*domain.Review, error) {
	id, err := uuid.NewV7()
	if err != nil {
		return nil, fmt.Errorf("failed to generate uuid: %w", err)
	}
	review.Id = id.String()

	return s.universityRepo.CreateReview(ctx, review)
}

func (s *universityService) GetReplies(ctx context.Context, reviewID string) ([]*domain.Review, error) {
	reviewID = strings.TrimSpace(reviewID)
	if reviewID == "" {
		return nil, errors.New("reviewID is required")
	}

	return s.universityRepo.GetReplies(ctx, reviewID)
}
