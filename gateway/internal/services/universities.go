package services

import (
	"context"
	"fmt"
	"gateway/api/universities"
	"gateway/internal/config"
	"gateway/internal/domain"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"time"
)

type UniversityService struct {
	cfg    *config.Config
	client universities.UniversitiesClient
}

func NewUniversityService(cfg *config.Config) domain.UniversityService {
	conn, err := grpc.NewClient(fmt.Sprintf("%s:%d", cfg.Services.UniversityService.Host, cfg.Services.UniversityService.Port), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		panic(err)
	}

	client := universities.NewUniversitiesClient(conn)

	return &UniversityService{
		cfg:    cfg,
		client: client,
	}

}

func (s *UniversityService) GetOpenDays(ctx context.Context, universityId string) ([]*domain.OpenDay, error) {
	req := &universities.UniversityId{Id: universityId}

	res, err := s.client.GetOpenDays(ctx, req)
	if err != nil {
		return nil, err
	}

	days := make([]*domain.OpenDay, len(res.Days))
	for i, day := range res.Days {
		days[i] = &domain.OpenDay{
			UniversityName: day.UniversityName,
			Description:    day.Description,
			Address:        day.Address,
			Link:           day.Link,
			Date:           time.UnixMilli(day.Time),
		}
	}

	return days, nil
}

func (s *UniversityService) GetReviews(ctx context.Context, universityId string, offset int, limit int) ([]*domain.Review, error) {
	req := &universities.GetReviewsRequest{
		Params: &universities.Params{
			Offset: int32(offset),
			Limit:  int32(limit),
		},
		UniversityId: universityId,
	}

	res, err := s.client.GetReviews(ctx, req)
	if err != nil {
		return nil, err
	}

	reviews := make([]*domain.Review, len(res.Reviews))
	for i, review := range res.Reviews {
		reviews[i] = &domain.Review{
			UniversityId: universityId,
			Date:         time.UnixMilli(review.Date),
			Text:         review.Text,
			AuthorStatus: review.AuthorStatus,
			RepliesCount: int(review.RepliesCount),
			Sentiment:    review.Sentiment,
		}
	}

	return reviews, nil
}
func (s *UniversityService) CreatePanorama(ctx context.Context, panorama *domain.Panorama) (*domain.Panorama, error) {
	p, err := s.client.CreatePanorama(ctx, &universities.CreatePanoramaRequest{
		Panorama: &universities.Panorama{
			UniversityId:   panorama.UniversityId,
			Address:        panorama.Address,
			Name:           panorama.Name,
			FirstLocation:  panorama.FirstLocation,
			SecondLocation: panorama.SecondLocation,
			Type:           domain.ConvertPanoramaToGrpc(panorama.Type),
		},
	})
	if err != nil {
		return nil, err
	}

	return &domain.Panorama{
		UniversityId:   p.UniversityId,
		Address:        p.Address,
		Name:           p.Name,
		FirstLocation:  p.FirstLocation,
		SecondLocation: p.SecondLocation,
		Type:           domain.ConvertPanoramaTypeFromGrpc(p.Type),
	}, nil
}

func (s *UniversityService) GetPanoramas(ctx context.Context, universityId string, category string) ([]*domain.Panorama, error) {
	rpcPanoramas, err := s.client.GetPanoramas(ctx, &universities.GetPanoramasRequest{
		UniversityId: universityId,
		Category:     category,
	})
	if err != nil {
		return nil, err
	}

	panoramas := make([]*domain.Panorama, len(rpcPanoramas.Panoramas))
	for i, p := range rpcPanoramas.Panoramas {
		panoramas[i] = &domain.Panorama{
			UniversityId:   p.UniversityId,
			Address:        p.Address,
			Name:           p.Name,
			FirstLocation:  p.FirstLocation,
			SecondLocation: p.SecondLocation,
			Type:           domain.ConvertPanoramaTypeFromGrpc(p.Type),
		}
	}

	return panoramas, nil
}
