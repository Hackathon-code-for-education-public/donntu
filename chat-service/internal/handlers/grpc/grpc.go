package grpc

import (
	"chat-service/api/chat"
	"chat-service/internal/dto"
	"chat-service/internal/entity"
	"chat-service/internal/models"
	"context"
	"log/slog"
)

type Service interface {
	Send(ctx context.Context, message *entity.Message) error
	Attach(ctx context.Context, chatId string, ch chan<- *dto.Message) error

	CreateChat(ctx context.Context, ch *entity.Chat) error
	ListChat(ctx context.Context, userId string) ([]*dto.Chat, error)
	GetHistory(ctx context.Context, chatId string) ([]*models.Message, error)
}

type Handler struct {
	chat.UnimplementedChatServer
	service Service
}

func NewHandler(service Service) *Handler {
	return &Handler{service: service}
}

func (h Handler) Send(ctx context.Context, request *chat.Message) (*chat.Empty, error) {

	message := &entity.Message{
		UserId: request.UserId,
		ChatId: request.ChatId,
		Text:   request.Text,
	}

	return &chat.Empty{}, h.service.Send(ctx, message)
}

func (h Handler) Attach(request *chat.AttachRequest, stream chat.Chat_AttachServer) error {
	ctx := stream.Context()

	msgCh := make(chan *dto.Message)

	if err := h.service.Attach(ctx, request.ChatId, msgCh); err != nil {
		return err
	}

	for {
		select {
		case <-ctx.Done():
			close(msgCh)
			return nil
		case m := <-msgCh:
			if err := stream.Send(&chat.IncomingMessage{
				Text: m.Text,
			}); err != nil {
				slog.Error("failed to send", slog.String("err", err.Error()))
				return err
			}
		}
	}
}

func (h Handler) Create(ctx context.Context, request *chat.CreateRequest) (*chat.CreateResponse, error) {

	ch := &entity.Chat{
		Participants: []string{request.UserId, request.TargetId},
	}

	if err := h.service.CreateChat(ctx, ch); err != nil {
		return nil, err
	}

	return &chat.CreateResponse{
		Id: ch.Id,
	}, nil
}

func (h Handler) List(request *chat.ListRequest, server chat.Chat_ListServer) error {
	chats, err := h.service.ListChat(server.Context(), request.UserId)
	if err != nil {
		return err
	}

	for _, ch := range chats {
		if err := server.Send(&chat.ChatEntity{
			ChatId: ch.Id,
		}); err != nil {
			return err
		}
	}

	return nil
}
func (h Handler) GetMessageHistory(request *chat.GetMessageHistoryRequest, stream chat.Chat_GetMessageHistoryServer) error {

	ctx := stream.Context()

	history, err := h.service.GetHistory(ctx, request.ChatId)
	if err != nil {
		return err
	}

	for _, msg := range history {
		if err := stream.Send(&chat.GetMessageHistoryItem{
			UserId:  msg.UserId,
			Message: msg.Text,
			SentAt:  msg.CreatedAt.Unix(),
		}); err != nil {
			return err
		}
	}

	return nil
}
