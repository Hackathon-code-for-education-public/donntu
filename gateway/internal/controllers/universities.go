package controllers

import (
	"gateway/internal/domain"
	"github.com/gofiber/fiber/v3"
	"log/slog"
)

type UniversitiesController struct {
	universityService domain.UniversityService
	log               *slog.Logger
}

func NewUniversityController(universityService domain.UniversityService, log *slog.Logger) *UniversitiesController {
	return &UniversitiesController{
		universityService: universityService,
		log:               log,
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
