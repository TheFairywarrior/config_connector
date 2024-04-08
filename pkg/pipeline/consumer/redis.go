package consumer

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/redis/go-redis/v9"
	"github.com/thefairywarrior/config_con/pkg/pipeline"
)

type RedisMessageData struct {
	ID   string         `json:id`
	Data map[string]any `json:data`
}

type RedisConsumer struct {
	host     string
	port     int
	database int
	channel  string
	passwrod string
}

func (r *RedisConsumer) Consume(ctx context.Context, c chan pipeline.Message) error {
	rdb := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%d", r.host, r.port),
		Password: r.passwrod,
		DB:       r.database,
	})

	pubsub := rdb.Subscribe(ctx, r.channel)
	defer pubsub.Close()

	_, err := pubsub.Receive(ctx)
	if err != nil {
		return err
	}

	ch := pubsub.Channel()

	for {
		select {
		case <-ctx.Done():
			return nil;
		case msg := <-ch:
			result := RedisMessageData{}
			err := json.Unmarshal([]byte(msg.Payload), &result)
			if err != nil {
				return err
			}
			c <- pipeline.NewDefaultMessage(result)
		}
	}
	return nil
}

func NewRedisConsumer(host string, port, database int, channel, password string) *RedisConsumer {
	return &RedisConsumer{
		host:     host,
		port:     port,
		database: database,
		channel:  channel,
		passwrod: password,
	}
}
