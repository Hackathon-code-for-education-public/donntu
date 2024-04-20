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
