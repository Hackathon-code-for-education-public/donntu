package grpc

import (
	"chat-service/api/chat"
	"context"
)

type Service interface {
	Send(context.Context, *chat.Message) error
}

type Handler struct {
	chat.UnimplementedChatServer
	service Service
}

func (h Handler) Send(ctx context.Context, message *chat.Message) (*chat.Empty, error) {

}

func (h Handler) Attach(request *chat.AttachRequest, server chat.Chat_AttachServer) error {

}
