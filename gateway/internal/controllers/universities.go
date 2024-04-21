package controllers

import (
	"gateway/internal/domain"
	"gateway/pkg/file"
	"github.com/gofiber/fiber/v3"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"log/slog"
)

type UniversitiesController struct {
	universityService domain.UniversityService
	fileService       domain.FileService
	log               *slog.Logger
}

func NewUniversityController(universityService domain.UniversityService,
	fileService domain.FileService,
	log *slog.Logger) *UniversitiesController {
	return &UniversitiesController{
		universityService: universityService,
		log:               log,
		fileService:       fileService,
	}
}

func (a *UniversitiesController) GetOpenDays() fiber.Handler {
	type request struct {
		Id string `query:"universityId"`
	}

	return func(c fiber.Ctx) error {
		var req request
		if err := c.Bind().Query(&req); err != nil {
			return bad(err.Error())
		}
		a.log.Info("get open days request: ", slog.String("universityId", req.Id))
		if req.Id == "" {
			return bad("university id is required")
		}

		days, err := a.universityService.GetOpenDays(c.Context(), req.Id)
		if err != nil {
			return internal(err.Error())
		}

		return ok(c, days)
	}
}

func (a *UniversitiesController) GetReviews() fiber.Handler {
	type request struct {
		Id     string `query:"universityId"`
		Offset int    `query:"offset"`
		Limit  int    `query:"limit"`
	}

	return func(ctx fiber.Ctx) error {
		var req request
		if err := ctx.Bind().Query(&req); err != nil {
			a.log.Error("error while bind request: ", slog.String("universityId", req.Id))
			return bad(err.Error())
		}
		a.log.Info("get reviews request: ", slog.String("universityId", req.Id))

		reviews, err := a.universityService.GetReviews(ctx.Context(), req.Id, req.Offset, req.Limit)
		if err != nil {
			a.log.Error("error while get reviews: ", slog.String("universityId", req.Id))
			return internal(err.Error())
		}

		return ok(ctx, reviews)
	}
}

func (a *UniversitiesController) CreatePanorama() fiber.Handler {
	type request struct {
		UniversityId string `json:"universityId"`
		Name         string `json:"name"`
		Address      string `json:"address"`
		Type         string `json:"type"`
	}

	return func(ctx fiber.Ctx) error {
		var req request
		if err := ctx.Bind().Body(&req); err != nil {
			a.log.Error("error while bind request: ", slog.String("name", req.Name))
			return bad(err.Error())
		}

		firstLocation, err := ctx.FormFile("firstLocation")
		if err != nil {
			a.log.Error("error while get first location: ", slog.String("name", req.Name))
			return internal(err.Error())
		}
		fLocationCType := firstLocation.Header.Get("Content-Type")
		reader, err := firstLocation.Open()
		if err != nil {
			a.log.Error("error while open first location: ", slog.String("name", req.Name))
			return internal(err.Error())
		}
		r := file.NewReader(reader, firstLocation.Size, fLocationCType)
		firstUrl, err := a.fileService.Upload(ctx.Context(), r)
		if err != nil {
			a.log.Error("error while upload first location: ", slog.String("name", req.Name))
			return internal(err.Error())
		}

		secondLocation, err := ctx.FormFile("secondLocation")
		if err != nil {
			a.log.Error("error while get second location: ", slog.String("name", req.Name))
			return internal(err.Error())
		}
		sLocationCType := secondLocation.Header.Get("Content-Type")
		reader, err = secondLocation.Open()
		if err != nil {
			a.log.Error("error while open second location: ", slog.String("name", req.Name))
			return internal(err.Error())
		}
		r = file.NewReader(reader, secondLocation.Size, sLocationCType)
		secondUrl, err := a.fileService.Upload(ctx.Context(), r)
		if err != nil {
			a.log.Error("error while upload second location: ", slog.String("name", req.Name))
			return internal(err.Error())
		}

		a.log.Info("get reviews count request: ", slog.String("name", req.Name))
		p, err := a.universityService.CreatePanorama(ctx.Context(), &domain.Panorama{
			UniversityId:   req.UniversityId,
			Address:        req.Address,
			Name:           req.Name,
			FirstLocation:  firstUrl,
			SecondLocation: secondUrl,
			Type:           req.Type,
		})

		return ok(ctx, p)
	}
}

func (a *UniversitiesController) GetPanorama() fiber.Handler {
	type request struct {
		UniversityId string `query:"universityId"`
		Category     string `query:"category"`
	}

	return func(ctx fiber.Ctx) error {
		var req request
		if err := ctx.Bind().Query(&req); err != nil {
			a.log.Error("error while bind request: ", slog.String("universityId", req.UniversityId))
			return bad(err.Error())
		}

		a.log.Info("get panorama request: ", slog.String("universityId", req.UniversityId))

		panoramas, err := a.universityService.GetPanoramas(ctx.Context(), req.UniversityId, req.Category)
		if err != nil {
			a.log.Error("error while get reviews: ", slog.String("universityId", req.UniversityId))
			return internal(err.Error())
		}

		return ok(ctx, panoramas)
	}
}

func (a *UniversitiesController) GetUniversitiesTop() fiber.Handler {
	type request struct {
		Count int `query:"count"`
	}

	return func(ctx fiber.Ctx) error {
		var req request
		if err := ctx.Bind().Query(&req); err != nil {
			a.log.Error("error while bind request: ", slog.Int("count", req.Count))
			return bad(err.Error())
		}

		a.log.Info("get universities top request: ", slog.Int("count", req.Count))

		universities, err := a.universityService.GetUniversitiesTop(ctx.Context(), req.Count)
		if err != nil {
			a.log.Error("error while get universities: ", slog.Int("count", req.Count))
			return internal(err.Error())
		}

		return ok(ctx, universities)
	}
}

func (a *UniversitiesController) SearchUniversities() fiber.Handler {
	type request struct {
		Name string `query:"name"`
	}

	return func(ctx fiber.Ctx) error {
		var req request
		if err := ctx.Bind().Query(&req); err != nil {
			a.log.Error("error while bind request: ", slog.String("name", req.Name))
			return bad(err.Error())
		}
		a.log.Info("get universities request: ")

		universities, err := a.universityService.SearchUniversities(ctx.Context(), req.Name)
		if err != nil {
			a.log.Error("error while get universities: ")
			return internal(err.Error())
		}

		return ok(ctx, universities)
	}
}

func (a *UniversitiesController) GetUniversity() fiber.Handler {
	return func(ctx fiber.Ctx) error {
		uid := ctx.Params("id")
		if uid == "" {
			return bad("university id is required")
		}

		a.log.Info("get university request: ", slog.String("universityId", uid))
		university, err := a.universityService.GetUniversity(ctx.Context(), uid)
		if err != nil {
			s, ok := status.FromError(err)
			if ok {
				switch s.Code() {
				case codes.NotFound:
					return notFound(err.Error())
				}
			}

			return internal(err.Error())
		}

		return ok(ctx, university)
	}
}

func (a *UniversitiesController) GetUniversities() fiber.Handler {
	type request struct {
		Offset int `query:"offset"`
		Limit  int `query:"limit"`
	}

	return func(ctx fiber.Ctx) error {
		var req request
		if err := ctx.Bind().Query(&req); err != nil {
			a.log.Error("error while bind request: ", slog.Int("offset", req.Offset), slog.Int("limit", req.Limit))
			return bad(err.Error())
		}
		a.log.Info("get universities request: ", slog.Int("offset", req.Offset), slog.Int("limit", req.Limit))
		universities, err := a.universityService.GetUniversities(ctx.Context(), req.Offset, req.Limit)
		if err != nil {
			a.log.Error("error while get universities: ", slog.Int("offset", req.Offset), slog.Int("limit", req.Limit))
			return internal(err.Error())
		}

		return ok(ctx, universities)
	}
}
