package grpc

import (
	"chat-service/api/chat"
	"chat-service/internal/entity"
	"context"
)

type Service interface {
	Send(context.Context, *entity.Message) error

	Attach(ctx context.Context, chatId string, ch chan *entity.Message) error

	CreateChat(ctx context.Context, ch *entity.Chat) error
	ListChat(ctx context.Context, userId string) ([]*entity.Chat, error)
	GetHistory(ctx context.Context, chatId string) ([]*entity.Message, error)
}

type Handler struct {
	chat.UnimplementedChatServer
	service Service
}

func (h Handler) Send(ctx context.Context, request *chat.Message) (*chat.Empty, error) {

	message := &entity.Message{
		UserId: request.UserId,
		ChatId: request.ChatId,
		Text:   request.Text,
	}

	return &chat.Empty{}, h.service.Send(ctx, message)
}

func (h Handler) Attach(request *chat.AttachRequest, server chat.Chat_AttachServer) error {

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
