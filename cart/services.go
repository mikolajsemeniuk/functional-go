package cart

import (
	"context"
	"encoding/json"
	"time"

	"github.com/go-redis/redis/v8"
)

type RedisReader func(context.Context, string) *redis.StringCmd

func RedisRead(r RedisReader) func(string) (string, error) {
	return func(key string) (string, error) {
		return r(context.Background(), key).Result()
	}
}

type RedisWriter func(context.Context, string, interface{}, time.Duration) *redis.StatusCmd

func RedisWrite(w RedisWriter) func(o Order, t time.Duration) error {
	return func(o Order, t time.Duration) error {
		data, err := json.Marshal(o)
		if err != nil {
			return err
		}

		_, err = w(context.Background(), o.Key, data, t).Result()
		return err
	}
}
