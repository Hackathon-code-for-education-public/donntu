package router

import (
	"context"
	"github.com/samber/lo"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log/slog"
	"universities/api/universities"
	"universities/internal/domain"
	"universities/internal/services"
)

type UniversitiesService struct {
	universities.UnimplementedUniversitiesServer
	service services.UniversityService
	log     *slog.Logger
}

func NewGRPCServer(grpcServer *grpc.Server, service services.UniversityService, log *slog.Logger) universities.UniversitiesServer {
	srv := &UniversitiesService{
		service: service,
		log:     log,
	}
	universities.RegisterUniversitiesServer(grpcServer, srv)
	reflection.Register(grpcServer)

	return srv
}

func (s *UniversitiesService) GetOpenDays(ctx context.Context, uni *universities.UniversityId) (*universities.OpenDays, error) {
	s.log.Info("get_open_days request received", slog.String("university_id", uni.Id))
	days, err := s.service.GetOpenDays(ctx, uni.Id)
	if err != nil {
		return nil, err
	}

	return &universities.OpenDays{
		Days: lo.Map(days, func(day *domain.OpenDay, _ int) *universities.OpenDay {
			return &universities.OpenDay{
				Link:           day.Link,
				Address:        day.Address,
				Description:    day.Description,
				UniversityName: day.UniversityName,
				Time:           day.Date.Unix(),
			}
		}),
	}, nil
}

func (s *UniversitiesService) GetReviews(ctx context.Context, request *universities.GetReviewsRequest) (*universities.Reviews, error) {
	s.log.Info("get_reviews request received", slog.String("university_id", request.UniversityId))

	reviews, err := s.service.GetReviews(ctx, request.UniversityId, int(request.Params.Limit), int(request.Params.Offset))
	if err != nil {
		return nil, err
	}

	return &universities.Reviews{
		Reviews: lo.Map(reviews, func(review *domain.Review, _ int) *universities.Review {
			return &universities.Review{
				Sentiment:    string(review.Sentiment),
				RepliesCount: int32(review.RepliesCount),
				AuthorStatus: string(review.AuthorStatus),
				Text:         review.Text,
				Date:         review.Date.Unix(),
				UniversityId: review.UniversityId,
			}
		}),
	}, nil
}
