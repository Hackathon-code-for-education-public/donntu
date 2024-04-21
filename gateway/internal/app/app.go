package app

import (
	"errors"
	"fmt"
	"gateway/internal/config"
	"gateway/internal/controllers"
	"gateway/internal/domain"
	fiberv2 "github.com/gofiber/fiber/v2"
	corsv2 "github.com/gofiber/fiber/v2/middleware/cors"
	loggerv2 "github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/cors"
	"github.com/gofiber/fiber/v3/middleware/logger"
	"log/slog"
)

type Application struct {
	cfg   *config.Config
	http  *fiber.App
	http2 *fiberv2.App

	authController       *controllers.AuthController
	universityController *controllers.UniversitiesController
	chatController       *controllers.ChatController
	log                  *slog.Logger
}

func NewApplication(cfg *config.Config, universityController *controllers.UniversitiesController, authController *controllers.AuthController, chatController *controllers.ChatController, log *slog.Logger) *Application {
	httpServer := fiber.New(fiber.Config{
		AppName:       "gateway",
		CaseSensitive: false,
		BodyLimit:     10 << 20,
		ErrorHandler: func(ctx fiber.Ctx, err error) error {
			code := fiber.StatusInternalServerError

			var e *fiber.Error
			if errors.As(err, &e) {
				code = e.Code
			}

			err = ctx.Status(code).JSON(fiber.Map{
				"message": e.Error(),
			})

			return nil
		},
	})

	http2Server := fiberv2.New(fiberv2.Config{
		AppName:       "gateway",
		CaseSensitive: false,
		BodyLimit:     10 << 20,
		ErrorHandler: func(ctx *fiberv2.Ctx, err error) error {
			code := fiber.StatusInternalServerError

			var e *fiber.Error
			if errors.As(err, &e) {
				code = e.Code
			}

			err = ctx.Status(code).JSON(fiber.Map{
				"message": e.Message,
			})

			return nil
		},
	})

	return &Application{
		cfg:                  cfg,
		http:                 httpServer,
		http2:                http2Server,
		log:                  log,
		authController:       authController,
		universityController: universityController,
		chatController:       chatController,
	}
}

// НЕИСПОЛЬЗУЙТЕ FIBER V3 НИКОГДА
func (a *Application) Run() error {
	a.http.Use(logger.New())

	a.http.Use(cors.New(cors.Config{
		AllowOrigins:     "*",
		AllowMethods:     "GET,POST,HEAD,PUT,DELETE,PATCH",
		AllowHeaders:     "Origin, Content-Type, Accept, Authorization",
		AllowCredentials: false,
	}))

	a.http2.Use(loggerv2.New())

	a.http2.Use(corsv2.New(corsv2.Config{
		AllowOrigins:     "*",
		AllowMethods:     "GET,POST,HEAD,PUT,DELETE,PATCH",
		AllowHeaders:     "Origin, Content-Type, Accept, Authorization",
		AllowCredentials: false,
	}))

	v1 := a.http.Group("/api/v1")

	au := v1.Group("/auth")
	au.Post("/sign-in", a.authController.SignIn())
	au.Post("/students/sign-up", a.authController.SignUp(domain.UserRoleStudent))
	au.Post("/applicant/sign-up", a.authController.SignUp(domain.UserRoleApplicant))
	au.Post("/university/sign-up", a.authController.SignUp(domain.UserRoleManager))
	au.Post("/manager/sign-up", a.authController.SignUp(domain.UserRoleManager))
	au.Post("/sign-out", a.authController.SignOut(), a.authController.AuthRequired(domain.UserRoleAny))
	au.Post("/refresh", a.authController.Refresh())

	v1.Get("/profile", a.authController.GetProfile(), a.authController.AuthRequired(domain.UserRoleAny))

	users := v1.Group("/users")
	users.Get("/:id", a.authController.GetUser())

	u := v1.Group("/universities")
	u.Get("/open", a.universityController.GetOpenDays())
	u.Get("/", a.universityController.GetUniversities())
	u.Get("/search", a.universityController.SearchUniversities())
	u.Get("/top", a.universityController.GetUniversitiesTop())
	u.Get("/:id", a.universityController.GetUniversity())

	r := v1.Group("/reviews")
	r.Get("/", a.universityController.GetReviews())
	r.Post("/", a.universityController.CreateReview(), a.authController.AuthRequired(domain.UserRoleStudent))
	r.Get("/:id", a.universityController.GetReplies())

	p := v1.Group("/panoramas")
	p.Get("/", a.universityController.GetPanorama())
	p.Post("/", a.universityController.CreatePanorama(), a.authController.AuthRequired(domain.UserRoleManager))

	v1v2 := a.http2.Group("/api/v1")
	chats := v1v2.Group("chats")
	chats.Get("/", a.authController.AuthRequiredV2(domain.UserRoleAny), a.chatController.GetChats())
	chats.Post("/", a.authController.AuthRequiredV2(domain.UserRoleAny), a.chatController.CreateChat(a.authController))
	chats.Get("/:id", a.authController.AuthRequiredV2(domain.UserRoleAny), a.chatController.Attach())
	chats.Post("/:id", a.authController.AuthRequiredV2(domain.UserRoleAny), a.chatController.SendMessage())
	chats.Get("/history/:id", a.authController.AuthRequiredV2(domain.UserRoleAny), a.chatController.GetHistory())

	go func() {
		err := a.http2.Listen(fmt.Sprintf(":%d", a.cfg.HTTP.Port+1))
		if err != nil {
			slog.Error(err.Error())
		}
	}()

	return a.http.Listen(fmt.Sprintf(":%d", a.cfg.HTTP.Port))
}
