package cart

import (
	"context"
	"testing"
	"time"

	"github.com/go-redis/redis/v8"
	"github.com/stretchr/testify/assert"
)

func TestReadOk(t *testing.T) {
	reader := func(context.Context, string) *redis.StringCmd {
		cmd := &redis.StringCmd{}
		cmd.SetVal(`{ "key": "123", "name": "test" }`)
		cmd.SetErr(nil)
		return cmd
	}

	read := RedisRead(reader)
	value, err := read("123")

	assert.Equal(t, nil, err, "err not nil")
	assert.Equal(t, `{ "key": "123", "name": "test" }`, value, "value not equal")
}

func TestWriteOk(t *testing.T) {
	writer := func(context.Context, string, interface{}, time.Duration) *redis.StatusCmd {
		cmd := &redis.StatusCmd{}
		cmd.SetVal(`{ "key": "123", "name": "test" }`)
		cmd.SetErr(nil)
		return cmd
	}

	write := RedisWrite(writer)
	err := write(Order{}, 0)

	assert.Equal(t, nil, err, "err not nil")
}
