package redis

import (
	"chat-service/internal/models"
	"context"
	"encoding/json"
	"github.com/redis/go-redis/v9"
)

type Broker struct {
	client *redis.Client
}

func NewBroker(client *redis.Client) *Broker {
	return &Broker{client: client}
}

func (r *Broker) Subscribe(ctx context.Context, topic string, handleMessage func(m *models.Message) error) error {
	sub := r.client.Subscribe(ctx, topic)

	for {
		select {
		case <-ctx.Done():
			return nil
		case msg := <-sub.Channel():
			m := &models.Message{}

			if err := json.Unmarshal([]byte(msg.Payload), &m); err != nil {
				return err
			}

			if err := handleMessage(m); err != nil {
				return err
			}
		}
	}
}

func (r *Broker) Publish(ctx context.Context, topic string, message *models.MessageDto) error {
	return r.client.Publish(ctx, topic, message).Err()
}
