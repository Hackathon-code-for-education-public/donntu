package app

import (
	"fmt"
	"gateway/internal/config"
	"gateway/internal/controllers"
	"gateway/internal/domain"
	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/cors"
	"github.com/gofiber/fiber/v3/middleware/logger"
	"log/slog"
)

type Application struct {
	cfg  *config.Config
	http *fiber.App

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
	})

	return &Application{
		cfg:                  cfg,
		http:                 httpServer,
		log:                  log,
		authController:       authController,
		universityController: universityController,
		chatController:       chatController,
	}
}

func (a *Application) Run() error {
	a.http.Use(logger.New())

	a.http.Use(cors.New(cors.Config{
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
	au.Post("/university/sign-up", a.authController.SignUp(domain.UserRoleStudent))
	au.Post("/manager/sign-up", a.authController.SignUp(domain.UserRoleManager))
	au.Post("/sign-out", a.authController.SignOut())
	au.Post("/refresh", a.authController.Refresh())

	v1.Get("/profile", a.authController.GetProfile(), a.authController.AuthRequired(nil))

	users := v1.Group("/users")
	users.Get("/:id", a.authController.GetUser())

	u := v1.Group("/universities")
	u.Get("/open", a.universityController.GetOpenDays())

	r := v1.Group("/reviews")
	r.Get("/", a.universityController.GetReviews())

	p := v1.Group("/panoramas")
	p.Get("/", a.universityController.GetPanorama())
	p.Post("/", a.universityController.CreatePanorama())

	chats := v1.Group("chats")
	chats.Get("/", a.chatController.GetChats(), a.authController.AuthRequired(nil))

	return a.http.Listen(fmt.Sprintf(":%d", a.cfg.HTTP.Port))
}
