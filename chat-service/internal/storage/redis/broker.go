package redis

import (
	"chat-service/internal/dto"
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

func (r *Broker) Subscribe(ctx context.Context, topic string, ch chan<- *dto.Message) error {
	sub := r.client.Subscribe(ctx, topic)

	for {
		select {
		case <-ctx.Done():
			return nil
		case msg := <-sub.Channel():
			m := &dto.Message{}

			if err := json.Unmarshal([]byte(msg.Payload), &m); err != nil {
				return err
			}

			ch <- m
		}
	}
}

func (r *Broker) Publish(ctx context.Context, topic string, message *dto.Message) error {
	return r.client.Publish(ctx, topic, message).Err()
}
