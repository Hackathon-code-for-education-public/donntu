package controllers

import (
	"errors"
	"fmt"
	"gateway/internal/domain"
	"gateway/internal/services"
	"github.com/go-playground/validator/v10"
	fiberv2 "github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v3"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"log/slog"
	"strings"
)

type AuthController struct {
	authService domain.AuthService
	log         *slog.Logger
	validator   *validator.Validate
}

func NewAuthController(authService domain.AuthService, log *slog.Logger) *AuthController {
	var val = validator.New()
	return &AuthController{
		authService: authService,
		log:         log,
		validator:   val,
	}
}

func (a *AuthController) SignIn() fiber.Handler {
	type request struct {
		Email    string `json:"email" validate:"required,len=11"`
		Password string `json:"password" validate:"required"`
	}

	return func(ctx fiber.Ctx) error {
		var req request
		if err := ctx.Bind().Body(&req); err != nil {
			a.log.Error("signIn error", slog.String("err", err.Error()))
			return bad(err.Error())
		}

		a.log.Info("signIn request", slog.String("email", req.Email))

		credentials := &domain.Credentials{
			Email:    req.Email,
			Password: req.Password,
		}

		tokens, err := a.authService.SignIn(ctx.Context(), credentials)
		if err != nil {
			a.log.Error("signIn error", slog.Any("err", err))
			return internal(err.Error())
		}

		return ok(ctx, tokens)
	}
}

func (a *AuthController) SignUp(role domain.UserRole) fiber.Handler {
	type request struct {
		Email      string `json:"email" validate:"required,email"`
		Password   string `json:"password" validate:"required"`
		FirstName  string `json:"firstName" validate:"required"`
		LastName   string `json:"lastName" validate:"required"`
		MiddleName string `json:"middleName" validate:"required"`
	}

	return func(ctx fiber.Ctx) error {
		var req request
		if err := ctx.Bind().Body(&req); err != nil {
			return bad(err.Error())
		}
		a.log.Debug("sign-up request", slog.Any("req", req))

		u := &domain.User{
			Email:      req.Email,
			LastName:   req.LastName,
			FirstName:  req.FirstName,
			MiddleName: req.MiddleName,
			Role:       role,
		}

		if errs := a.Validate(req); errs != "" {
			return bad(errs)
		}

		tokens, err := a.authService.SignUp(ctx.Context(), u, req.Password)
		if err != nil {
			if e, ok := status.FromError(err); ok {
				switch e.Code() {
				case codes.AlreadyExists:
					a.log.Error("user already exists", slog.Any("err", e))
					return bad(e.Message())
				}
			}

			a.log.Error("internal error", slog.Any("err", err))
			return internal(err.Error())
		}

		return ok(ctx, tokens)
	}
}

func (a *AuthController) SignOut() fiber.Handler {
	return func(ctx fiber.Ctx) error {
		a.log.Debug("sign-out request")

		accessToken, k := ctx.Locals("accessToken").(string)
		if !k {
			a.log.Error("accessToken not found")
			return internal("accessToken not found")
		}

		if err := a.authService.SignOut(ctx.Context(), accessToken); err != nil {
			a.log.Error("internal error", slog.Any("err", err))
			return internal(err.Error())
		}

		return ok(ctx)
	}
}

func (a *AuthController) Refresh() fiber.Handler {
	type req struct {
		RefreshToken string `json:"refreshToken" validate:"required"`
	}
	return func(ctx fiber.Ctx) error {
		var r req
		if err := ctx.Bind().Body(&r); err != nil {
			a.log.Error("internal error", slog.Any("err", err))
			return bad(err.Error())
		}
		a.log.Debug("refresh request", slog.Any("req", r))

		tokens, err := a.authService.Refresh(ctx.Context(), r.RefreshToken)
		if err != nil {
			a.log.Error("internal error", slog.Any("err", err))
			return internal(err.Error())
		}

		return ok(ctx, tokens)
	}
}

func (a *AuthController) AuthRequiredV2(role domain.UserRole) fiberv2.Handler {
	return func(ctx *fiberv2.Ctx) error {
		auth := ctx.Get("Authorization")
		s := strings.Split(auth, " ")
		if len(s) != 2 {
			a.log.Error("Authorization not found")
			return bad("Authorization not found")
		}

		accessToken := s[1]

		a.log.Info("auth-required request", slog.Any("accessToken", accessToken))

		u, err := a.authService.Verify(ctx.Context(), accessToken, role)
		if err != nil {
			if errors.Is(err, services.ErrUnauthorized) {
				a.log.Error("unauthorized", slog.Any("err", err))
				return unauthorized(err.Error())
			}
			if errors.Is(err, services.ErrForbidden) {
				a.log.Error("forbidden", slog.Any("err", err))

				return forbidden(err.Error())
			}
			if errors.Is(err, services.ErrNotFound) {
				a.log.Error("not found", slog.Any("err", err))
				return notFound(err.Error())
			}

			if errors.Is(err, services.ErrInvalidRequest) {
				a.log.Error("invalid request", slog.Any("err", err))
				return bad(err.Error())
			}

			a.log.Error("internal error", slog.Any("err", err))
			return internal(err.Error())
		}

		ctx.Locals("accessToken", accessToken)
		ctx.Locals("user", u)

		return ctx.Next()
	}
}

func (a *AuthController) AuthRequired(role domain.UserRole) fiber.Handler {
	return func(ctx fiber.Ctx) error {
		auth := ctx.Get("Authorization")
		s := strings.Split(auth, " ")
		if len(s) != 2 {
			a.log.Error("Authorization not found")
			return bad("Authorization not found")
		}

		accessToken := s[1]

		a.log.Info("auth-required request", slog.Any("accessToken", accessToken))

		u, err := a.authService.Verify(ctx.Context(), accessToken, role)
		if err != nil {
			if errors.Is(err, services.ErrUnauthorized) {
				a.log.Error("unauthorized", slog.Any("err", err))
				return unauthorized(err.Error())
			}
			if errors.Is(err, services.ErrForbidden) {
				a.log.Error("forbidden", slog.Any("err", err))

				return forbidden(err.Error())
			}
			if errors.Is(err, services.ErrNotFound) {
				a.log.Error("not found", slog.Any("err", err))
				return notFound(err.Error())
			}

			if errors.Is(err, services.ErrInvalidRequest) {
				a.log.Error("invalid request", slog.Any("err", err))
				return bad(err.Error())
			}

			a.log.Error("internal error", slog.Any("err", err))
			return internal(err.Error())
		}

		ctx.Locals("accessToken", accessToken)
		ctx.Locals("user", u)

		return ctx.Next()
	}
}
func (a *AuthController) GetUser() fiber.Handler {
	return func(c fiber.Ctx) error {

		userId := c.Params("id", "")
		if userId == "" {
			return bad("invalid user id")
		}

		user, err := a.authService.GetUser(c.Context(), userId)
		if err != nil {
			return err
		}

		return ok(c, user)
	}
}

func (a *AuthController) GetProfile() fiber.Handler {

	type response struct {
		Id         string          `json:"id"`
		Email      string          `json:"email"`
		LastName   string          `json:"lastName"`
		FirstName  string          `json:"firstName"`
		MiddleName string          `json:"middleName"`
		Role       domain.UserRole `json:"role"`
	}

	return func(c fiber.Ctx) error {
		user, k := c.Locals("user").(*domain.UserClaims)
		if !k {
			return internal("cannot parse locals")
		}

		u, err := a.authService.GetUser(c.Context(), user.Id)
		if err != nil {
			return internal(err.Error())
		}

		return ok(c, response{
			Id:         user.Id,
			Email:      u.Email,
			LastName:   u.LastName,
			FirstName:  u.FirstName,
			MiddleName: u.MiddleName,
			Role:       u.Role,
		})
	}
}

func (a *AuthController) Validate(data any) string {
	sb := &strings.Builder{}

	errs := a.validator.Struct(data)
	if errs != nil {
		for _, err := range errs.(validator.ValidationErrors) {
			sb.WriteString(fmt.Sprintf("%s\n", err.Error()))
		}
	}

	return sb.String()
}
