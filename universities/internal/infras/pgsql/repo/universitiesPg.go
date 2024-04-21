package repo

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"github.com/lib/pq"
	"github.com/samber/lo"
	"time"
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

func (u *universitiesPg) CreateReview(ctx context.Context, review *domain.Review) (*domain.Review, error) {
	querier := postgresql.New(u.pg.GetDB())
	params := postgresql.CreateReviewParams{
		ReviewID:     review.Id,
		UniversityID: review.UniversityId,
		AuthorStatus: postgresql.Statuses(review.AuthorStatus),
		Sentiment:    postgresql.Sentiments(review.Sentiment),
		Date:         time.Now(),
		Text:         review.Text,
	}
	if review.ParentId != nil {
		params.ParentReviewID = sql.NullString{
			String: *review.ParentId,
			Valid:  true,
		}
	}
	item, err := querier.CreateReview(ctx, params)
	if err != nil {
		if pqErr, ok := err.(*pq.Error); ok {
			if pqErr.Code == "23503" {
				return nil, errors.New("not found")
			}
			return nil, errors.New(pqErr.Message)
		}
		return nil, err
	}

	return &domain.Review{
		Id:           item.ReviewID,
		UniversityId: item.UniversityID,
		Date:         item.Date,
		Text:         item.Text,
		AuthorStatus: domain.AuthorStatus(item.AuthorStatus),
		RepliesCount: 0,
		Sentiment:    domain.Sentiment(item.Sentiment),
	}, nil
}

func (u *universitiesPg) GetReviews(ctx context.Context, universityID string, limit, offset int) ([]*domain.Review, error) {
	querier := postgresql.New(u.pg.GetDB())
	results, err := querier.GetReviews(ctx, postgresql.GetReviewsParams{Limit: int32(limit), Offset: int32(offset), UniversityID: universityID})
	if err != nil {
		return nil, err
	}

	reviews := lo.Map(results, func(item postgresql.GetReviewsRow, _ int) *domain.Review {
		return &domain.Review{
			Id:           item.ReviewID,
			UniversityId: item.UniversityID,
			Date:         item.Date,
			Text:         item.Text,
			AuthorStatus: domain.AuthorStatus(item.AuthorStatus),
			RepliesCount: int(item.ReplyCount),
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
	db := u.pg.GetDB()
	query := fmt.Sprintf("select * from university_panoramas where university_id = $1")
	if category != "" {
		query += fmt.Sprintf(" and type = '%s'", category)
	}
	panoramas := make([]*domain.Panorama, 0)
	rows, err := db.Query(query, universityID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var panorama domain.Panorama
		if err := rows.Scan(&panorama.UniversityId,
			&panorama.Address,
			&panorama.Name,
			&panorama.FirstLocation,
			&panorama.LastLocation,
			&panorama.Type,
		); err != nil {
			return nil, err
		}
		panoramas = append(panoramas, &panorama)
	}

	return panoramas, nil
}

func (u *universitiesPg) GetUniversitiesTop(ctx context.Context, n int) ([]*domain.University, error) {
	querier := postgresql.New(u.pg.GetDB())
	results, err := querier.GetTopOfUniversities(ctx, int32(n))
	if err != nil {
		return nil, err
	}

	return lo.Map(results, func(item postgresql.University, _ int) *domain.University {
		return &domain.University{
			Id:           item.ID,
			Name:         item.Name,
			LongName:     item.LongName,
			Logo:         item.Logo,
			Rating:       item.Rating,
			Region:       item.Region,
			Type:         string(item.Type),
			StudyFields:  int(item.StudyFields),
			BudgetPlaces: int(item.BudgetPlaces),
		}
	}), nil
}

func (u *universitiesPg) GetUniversities(ctx context.Context, offset, limit int) ([]*domain.University, error) {
	querier := postgresql.New(u.pg.GetDB())
	results, err := querier.GetUniversities(ctx, postgresql.GetUniversitiesParams{
		Offset: int32(offset),
		Limit:  int32(limit),
	})
	if err != nil {
		return nil, err
	}

	return lo.Map(results, func(item postgresql.University, _ int) *domain.University {
		return &domain.University{
			Id:           item.ID,
			Name:         item.Name,
			LongName:     item.LongName,
			Logo:         item.Logo,
			Rating:       item.Rating,
			Region:       item.Region,
			Type:         string(item.Type),
			StudyFields:  int(item.StudyFields),
			BudgetPlaces: int(item.BudgetPlaces),
		}
	}), nil
}

func (u *universitiesPg) GetUniversity(ctx context.Context, universityID string) (*domain.University, error) {
	querier := postgresql.New(u.pg.GetDB())
	item, err := querier.GetUniversity(ctx, universityID)
	if err != nil {

	}
	if len(item.ID) == 0 {
		return nil, errors.New("not found")
	}

	return &domain.University{
		Id:           item.ID,
		Name:         item.Name,
		LongName:     item.LongName,
		Logo:         item.Logo,
		Rating:       item.Rating,
		Region:       item.Region,
		Type:         string(item.Type),
		StudyFields:  int(item.StudyFields),
		BudgetPlaces: int(item.BudgetPlaces),
	}, nil
}

func (u *universitiesPg) SearchUniversities(ctx context.Context, name string) ([]*domain.University, error) {
	querier := postgresql.New(u.pg.GetDB())
	results, err := querier.SearchUniversities(ctx, postgresql.SearchUniversitiesParams{
		LongName: "%" + name + "%",
		Name:     "%" + name + "%",
	})
	if err != nil {
		return nil, err
	}

	return lo.Map(results, func(item postgresql.University, _ int) *domain.University {
		return &domain.University{
			Id:           item.ID,
			Name:         item.Name,
			LongName:     item.LongName,
			Logo:         item.Logo,
			Rating:       item.Rating,
			Region:       item.Region,
			Type:         string(item.Type),
			StudyFields:  int(item.StudyFields),
			BudgetPlaces: int(item.BudgetPlaces),
		}
	}), nil
}

func (u *universitiesPg) GetReplies(ctx context.Context, reviewID string) ([]*domain.Review, error) {
	querier := postgresql.New(u.pg.GetDB())
	results, err := querier.GetReviewsByParent(ctx, sql.NullString{
		String: reviewID,
		Valid:  true,
	})
	if err != nil {
		return nil, err
	}

	return lo.Map(results, func(item postgresql.UniversityReview, _ int) *domain.Review {
		return &domain.Review{
			Id:           item.ReviewID,
			UniversityId: item.UniversityID,
			Text:         item.Text,
			Date:         item.Date,
			Sentiment:    domain.Sentiment(item.Sentiment),
			AuthorStatus: domain.AuthorStatus(item.AuthorStatus),
		}
	}), nil
}
