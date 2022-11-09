package cart

import (
	"context"
	"testing"

	"github.com/go-redis/redis/v8"
)

func TestRead(t *testing.T) {
	getter := func(context.Context, string) *redis.StringCmd {
		return &redis.StringCmd{}
	}

	read(getter)
}

func TestWrite(t *testing.T) {
	getter := func(context.Context, string) *redis.StringCmd {
		return &redis.StringCmd{}
	}

	read(getter)
}
