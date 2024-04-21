package router

import (
	"context"
	"github.com/samber/lo"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/reflection"
	"google.golang.org/grpc/status"
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
				ReviewId:     review.Id,
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

func (s *UniversitiesService) CreatePanorama(ctx context.Context, request *universities.CreatePanoramaRequest) (*universities.Panorama, error) {
	s.log.Info("create_panorama request received", slog.String("university_id", request.Panorama.UniversityId))

	_, err := s.service.CreatePanorama(ctx, &domain.Panorama{
		UniversityId:  request.Panorama.UniversityId,
		Address:       request.Panorama.Address,
		Name:          request.Panorama.Name,
		FirstLocation: request.Panorama.FirstLocation,
		LastLocation:  request.Panorama.SecondLocation,
		Type:          domain.ConvertPanoramaTypeFromGrpc(request.Panorama.Type),
	})
	if err != nil {
		s.log.Error(err.Error())
		return nil, err
	}

	return request.Panorama, nil
}

func (s *UniversitiesService) GetPanoramas(ctx context.Context, request *universities.GetPanoramasRequest) (*universities.Panoramas, error) {
	s.log.Info("get_panoramas request received", slog.String("university_id", request.UniversityId))
	panoramas, err := s.service.GetPanoramas(ctx, request.UniversityId, request.Category)
	if err != nil {
		return nil, err
	}

	return &universities.Panoramas{
		Panoramas: lo.Map(panoramas, func(panorama *domain.Panorama, _ int) *universities.Panorama {
			return &universities.Panorama{
				UniversityId:   panorama.UniversityId,
				Address:        panorama.Address,
				Name:           panorama.Name,
				FirstLocation:  panorama.FirstLocation,
				SecondLocation: panorama.LastLocation,
				Type:           domain.ConvertPanoramaToGrpc(panorama.Type),
			}
		}),
	}, nil
}

func (s *UniversitiesService) GetUniversity(ctx context.Context, id *universities.UniversityId) (*universities.University, error) {
	s.log.Info("get_university request received", slog.String("university_id", id.Id))
	u, err := s.service.GetUniversity(ctx, id.Id)
	if err != nil {
		return nil, status.Error(codes.NotFound, err.Error())
	}

	return &universities.University{
		Id:           u.Id,
		Name:         u.Name,
		LongName:     u.LongName,
		Logo:         u.Logo,
		Rating:       float32(u.Rating),
		Region:       u.Region,
		Type:         u.Type,
		StudyFields:  int32(u.StudyFields),
		BudgetPlaces: int32(u.BudgetPlaces),
	}, nil
}

func (s *UniversitiesService) GetUniversities(ctx context.Context, params *universities.PageParams) (*universities.UniversitiesSchema, error) {
	s.log.Info("get_universities request received")
	us, err := s.service.GetUniversities(ctx, int(params.Offset), int(params.Limit))
	if err != nil {
		return nil, err
	}

	return &universities.UniversitiesSchema{
		Universities: lo.Map(us, func(u *domain.University, _ int) *universities.University {
			return &universities.University{
				Id:           u.Id,
				Name:         u.Name,
				LongName:     u.LongName,
				Logo:         u.Logo,
				Rating:       float32(u.Rating),
				Region:       u.Region,
				Type:         u.Type,
				StudyFields:  int32(u.StudyFields),
				BudgetPlaces: int32(u.BudgetPlaces),
			}
		}),
	}, nil
}

func (s *UniversitiesService) SearchUniversities(ctx context.Context, request *universities.SearchUniversitiesRequest) (*universities.UniversitiesSchema, error) {
	s.log.Info("search_universities request received", slog.String("name", request.Name))
	us, err := s.service.SearchUniversities(ctx, request.Name)
	if err != nil {
		return nil, err
	}

	return &universities.UniversitiesSchema{
		Universities: lo.Map(us, func(u *domain.University, _ int) *universities.University {
			return &universities.University{
				Id:           u.Id,
				Name:         u.Name,
				LongName:     u.LongName,
				Logo:         u.Logo,
				Rating:       float32(u.Rating),
				Region:       u.Region,
				Type:         u.Type,
				StudyFields:  int32(u.StudyFields),
				BudgetPlaces: int32(u.BudgetPlaces),
			}
		}),
	}, nil
}

func (s *UniversitiesService) GetTopOfUniversities(ctx context.Context, request *universities.GetTopOfUniversitiesRequest) (*universities.UniversitiesSchema, error) {
	s.log.Info("get_top_of_universities request received")
	us, err := s.service.GetUniversitiesTop(ctx, int(request.Count))
	if err != nil {
		return nil, err
	}

	return &universities.UniversitiesSchema{
		Universities: lo.Map(us, func(u *domain.University, _ int) *universities.University {
			return &universities.University{
				Id:           u.Id,
				Name:         u.Name,
				LongName:     u.LongName,
				Logo:         u.Logo,
				Rating:       float32(u.Rating),
				Region:       u.Region,
				Type:         u.Type,
				StudyFields:  int32(u.StudyFields),
				BudgetPlaces: int32(u.BudgetPlaces),
			}
		}),
	}, nil
}
