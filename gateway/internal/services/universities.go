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
