package service

import (
	"chat-service/internal/dto"
	"chat-service/internal/entity"
	"chat-service/internal/handlers/grpc"
	"chat-service/internal/models"
	"context"
	"github.com/google/uuid"
	"log/slog"
)

var _ grpc.Service = (*Service)(nil)

type ChatStorage interface {
	Create(ctx context.Context, id string, userId string, targetId string) error
	ListChatByUserId(ctx context.Context, userId string) ([]*dto.Chat, error)
	GetHistory(ctx context.Context, chatId string) ([]*models.Message, error)
}

type MessageStorage interface {
	Save(ctx context.Context, message *models.Message) error
	MarkAsRead(ctx context.Context, messageId int64) error
}

type Broker interface {
	Subscribe(ctx context.Context, topic string, ch chan<- *dto.Message) error
	Publish(ctx context.Context, topic string, message *dto.Message) error
}

type Service struct {
	broker         Broker
	chatStorage    ChatStorage
	messageStorage MessageStorage
	//userStorage    UserStorage
}

func NewService(broker Broker, chatStorage ChatStorage, messageStorage MessageStorage) *Service {
	return &Service{
		broker:         broker,
		chatStorage:    chatStorage,
		messageStorage: messageStorage,
	}
}

func (s Service) Send(ctx context.Context, message *entity.Message) error {

	m := &models.Message{
		UserId: message.UserId,
		ChatId: message.ChatId,
		Text:   message.Text,
	}

	if err := s.messageStorage.Save(ctx, m); err != nil {
		return err
	}

	mdto := &dto.Message{
		Id:     m.Id,
		UserId: m.UserId,
		Text:   m.Text,
	}

	if err := s.broker.Publish(ctx, message.ChatId, mdto); err != nil {
		return err
	}

	return nil
}

func (s Service) Attach(ctx context.Context, chatId string, ch chan<- *dto.Message) error {

	tch := make(chan *dto.Message)

	go func() {
		for {
			select {
			case <-ctx.Done():
				close(tch)
			case msg := <-tch:
				if err := s.messageStorage.MarkAsRead(ctx, msg.Id); err != nil {
					slog.Error("error mark as read", slog.String("err", err.Error()))
					return
				}
				ch <- msg
			}
		}
	}()

	if err := s.broker.Subscribe(ctx, chatId, ch); err != nil {
		slog.Error("error subscribe", slog.String("err", err.Error()))
		return err
	}

	return nil
}

func (s Service) CreateChat(ctx context.Context, ch *entity.Chat) error {

	id, err := uuid.NewV7()
	if err != nil {
		return err
	}

	if err := s.chatStorage.Create(ctx, id.String(), ch.Participants[0], ch.Participants[1]); err != nil {
		return err
	}

	return nil
}

func (s Service) ListChat(ctx context.Context, userId string) ([]*dto.Chat, error) {
	return s.chatStorage.ListChatByUserId(ctx, userId)
}

func (s Service) GetHistory(ctx context.Context, chatId string) ([]*models.Message, error) {
	return s.chatStorage.GetHistory(ctx, chatId)
}
