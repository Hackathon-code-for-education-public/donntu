package repo

import (
	"context"
	"github.com/samber/lo"
	"universities/internal/domain"
	postgresql "universities/internal/infras/pgsql"
	"universities/internal/services"
	"universities/pkg/engine"
)

type universitiesPg struct {
	pg engine.DBEngine
}

func NewUniversitiesPgRepo(pg engine.DBEngine) services.UniversityRepository {
	return &universitiesPg{pg: pg}
}

func (u *universitiesPg) GetOpenDays(ctx context.Context, universityID string) ([]*domain.OpenDay, error) {
	querier := postgresql.New(u.pg.GetDB())
	results, err := querier.GetOpenDays(ctx, universityID)
	if err != nil {
		return nil, err
	}

	days := lo.Map(results, func(item postgresql.GetOpenDaysRow, _ int) *domain.OpenDay {
		return &domain.OpenDay{
			UniversityName: item.Name,
			Description:    item.Description,
			Address:        item.Address,
			Link:           item.Link,
			Date:           item.Date,
		}
	})

	return days, nil
}

func (u *universitiesPg) GetReviews(ctx context.Context, universityID string, limit, offset int) ([]*domain.Review, error) {
	querier := postgresql.New(u.pg.GetDB())
	results, err := querier.GetReviews(ctx, postgresql.GetReviewsParams{Limit: int32(limit), Offset: int32(offset), UniversityID: universityID})
	if err != nil {
		return nil, err
	}

	reviews := lo.Map(results, func(item postgresql.UniversityReview, _ int) *domain.Review {
		return &domain.Review{
			UniversityId: item.UniversityID,
			Date:         item.Date,
			Text:         item.Text,
			AuthorStatus: domain.AuthorStatus(item.AuthorStatus),
			RepliesCount: int(item.Repliescount),
			Sentiment:    domain.Sentiment(item.Sentiment),
		}
	})

	return reviews, nil
}

func (u *universitiesPg) CreatePanorama(ctx context.Context, p *domain.Panorama) (*domain.Panorama, error) {
	querier := postgresql.New(u.pg.GetDB())
	_, err := querier.AddPanorama(ctx,
		postgresql.AddPanoramaParams{
			Type:           postgresql.PanoramaTypes(p.Type),
			UniversityID:   p.UniversityId,
			Address:        p.Address,
			Name:           p.Name,
			Firstlocation:  p.FirstLocation,
			Secondlocation: p.LastLocation})
	if err != nil {
		return nil, err
	}

	return p, nil
}

func (u *universitiesPg) GetPanoramas(ctx context.Context, universityID string, category string) ([]*domain.Panorama, error) {
	querier := postgresql.New(u.pg.GetDB())
	panoramas, err := querier.GetPanoramas(ctx, postgresql.GetPanoramasParams{
		UniversityID: universityID,
		Type:         postgresql.PanoramaTypes(category),
	})
	if err != nil {
		return nil, err
	}

	return lo.Map(panoramas, func(item postgresql.UniversityPanorama, _ int) *domain.Panorama {
		return &domain.Panorama{
			UniversityId:  item.UniversityID,
			Address:       item.Address,
			Name:          item.Name,
			FirstLocation: item.Firstlocation,
			LastLocation:  item.Secondlocation,
			Type:          string(item.Type),
		}
	}), nil
}
