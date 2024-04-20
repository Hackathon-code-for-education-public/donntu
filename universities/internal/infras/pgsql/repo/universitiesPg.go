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
