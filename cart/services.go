package cart

import (
	"context"
	"time"

	"github.com/go-redis/redis/v8"
)

type Order struct {
	Key  string `json:"key"`
	Name string `json:"name"`
}

type reader func(context.Context, string) *redis.StringCmd

func read(r reader) func(string) (string, error) {
	return func(key string) (string, error) {
		result, err := r(context.Background(), key).Result()
		if err != nil {
			return "", err
		}
		return result, nil
	}
}

type writer func(ctx context.Context, key string, value interface{}, expiration time.Duration) *redis.StatusCmd

func write(w writer) func(o Order, t time.Duration) error {
	return func(o Order, t time.Duration) error {
		_, err := w(context.Background(), o.Key, o, t).Result()
		if err != nil {
			return err
		}
		return nil
	}
}
