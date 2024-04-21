package grpc

import (
	"chat-service/api/chat"
	"chat-service/internal/entity"
	"context"
)

type Service interface {
	Send(context.Context, *chat.Message) error

	CreateChat(ctx context.Context, chat2 *entity.Chat) error
}

type Handler struct {
	chat.UnimplementedChatServer
	service Service
}

func (h Handler) Send(ctx context.Context, message *chat.Message) (*chat.Empty, error) {

}

func (h Handler) Attach(request *chat.AttachRequest, server chat.Chat_AttachServer) error {
	messagesChan := make(chan *entity.Message)
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
	//TODO implement me
	panic("implement me")
}
