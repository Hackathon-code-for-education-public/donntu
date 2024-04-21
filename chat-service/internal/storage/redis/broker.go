package redis

import (
	"chat-service/internal/dto"
	"context"
	"encoding/json"
	"github.com/redis/go-redis/v9"
	"log/slog"
)

type Broker struct {
	client *redis.Client
}

func NewBroker(client *redis.Client) *Broker {
	return &Broker{client: client}
}

func (r *Broker) Subscribe(ctx context.Context, topic string, ch chan<- *dto.Message) error {

	log := slog.With(slog.String("topic", topic))
	log.Debug("subscribe topic")

	sub := r.client.Subscribe(ctx, topic)

	go func() {
		for {
			select {
			case <-ctx.Done():
				return
			case msg := <-sub.Channel():
				log.Debug("received message")
				m := &dto.Message{}
				if err := json.Unmarshal([]byte(msg.Payload), &m); err != nil {
					return
				}
				ch <- m
			}
		}
	}()

	return nil
}

func (r *Broker) Publish(ctx context.Context, topic string, message *dto.Message) error {
	payload, err := json.Marshal(message)
	if err != nil {
		return err
	}
	return r.client.Publish(ctx, topic, payload).Err()
}
