package controllers

import (
	"fmt"
	"gateway/api/chat"
	"gateway/internal/config"
	"gateway/internal/domain"
	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/log"
	"github.com/gorilla/websocket"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"io"
	"log/slog"
)

type ChatController struct {
	client   chat.ChatClient
	upgrader websocket.Upgrader
}

func NewChatController(cfg *config.Config) *ChatController {
	conn, err := grpc.NewClient(fmt.Sprintf("%s:%d", cfg.Services.ChatService.Host, cfg.Services.ChatService.Port), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		panic(err)
	}

	return &ChatController{
		upgrader: websocket.Upgrader{
			ReadBufferSize:  1024,
			WriteBufferSize: 1024,
		},
		client: chat.NewChatClient(conn),
	}
}

func (c *ChatController) GetChats() fiber.Handler {

	type chats struct {
		UserId string `json:"userId"`
		ChatId string `json:"chatId"`
	}

	return func(ctx fiber.Ctx) error {

		l := slog.With(slog.String("handler", "ChatController.GetChats"))

		u, okk := ctx.Locals("user").(*domain.UserClaims)
		if !okk {
			l.Error("cannot get user claims from context")
			return internal("internal")
		}

		stream, err := c.client.List(ctx.Context(), &chat.ListRequest{
			UserId: u.Id,
		})
		if err != nil {
			l.Error("stream cant get", slog.String("err", err.Error()))
			return internal(err.Error())
		}

		res := make([]*chats, 0)

		recv, err := stream.Recv()
		if err != nil {
			l.Error("stream cant received", slog.String("err", err.Error()))
			return internal(err.Error())
		}
		for recv != nil {
			res = append(res, &chats{
				//UserId: recv.UserId,
				ChatId: recv.ChatId,
			})
			recv, err = stream.Recv()
			if err != nil {
				if err == io.EOF {
					log.Debug("stream closed")
					break
				}
				l.Error("stream cant received", slog.String("err", err.Error()))
				return internal(err.Error())
			}
		}

		return ok(ctx, res)
	}
}

func (c *ChatController) CreateChat() fiber.Handler {

	type request struct {
		TargetId string `json:"targetId"`
	}

	return func(ctx fiber.Ctx) error {
		l := slog.With(slog.String("handler", "ChatController.GetChats"))

		var req request
		if err := ctx.Bind().Body(&req); err != nil {
			l.Error("cant parse request", slog.String("err", err.Error()))
			return bad(err.Error())
		}

		u, okk := ctx.Locals("user").(*domain.UserClaims)
		if !okk {
			l.Error("cannot get user claims from context")
			return internal("internal")
		}

		res, err := c.client.Create(ctx.Context(), &chat.CreateRequest{
			UserId:   u.Id,
			TargetId: req.TargetId,
		})
		if err != nil {
			l.Error("cant create chat", slog.String("err", err.Error()))
			return internal(err.Error())
		}

		return ok(ctx, res)
	}
}

func (c *ChatController) Attach() fiber.Handler {
	return func(ctx fiber.Ctx) error {
		l := slog.With(slog.String("handler", "ChatController.Attach"))

		u, okk := ctx.Locals("user").(*domain.UserClaims)
		if !okk {
			l.Error("cannot get user claims from context")
			return internal("internal")
		}

	}
}
