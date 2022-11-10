package cart

import (
	"context"
	"net/http/httptest"
	"strings"
	"testing"
	"time"

	"github.com/go-redis/redis/v8"
	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
)

var (
	get = func(context.Context, string) *redis.StringCmd { return nil }
	set = func(context.Context, string, interface{}, time.Duration) *redis.StatusCmd { return nil }
)

type mock struct{}

func (m *mock) Get(c context.Context, s string) *redis.StringCmd {
	return get(c, s)
}
func (m *mock) Set(c context.Context, s string, i interface{}, t time.Duration) *redis.StatusCmd {
	return set(c, s, i, t)
}

func TestReadOrderOk(t *testing.T) {
	m := &mock{}
	get = func(context.Context, string) *redis.StringCmd {
		cmd := &redis.StringCmd{}
		cmd.SetVal(`{ "key": "123", "name": "test" }`)
		cmd.SetErr(nil)
		return cmd
	}

	route := "/test"
	app := fiber.New()
	handler := ReadOrder(m)
	app.Get(route, handler)

	request := httptest.NewRequest(fiber.MethodGet, route, nil)
	response, err := app.Test(request, -1)

	assert.Equal(t, fiber.StatusOK, response.StatusCode, "http status code not equal")
	assert.Equal(t, nil, err, "err not nil")
}

func TestWriteOrderOk(t *testing.T) {
	m := &mock{}
	set = func(context.Context, string, interface{}, time.Duration) *redis.StatusCmd {
		cmd := &redis.StatusCmd{}
		cmd.SetErr(nil)
		return cmd
	}

	route := "/test"
	app := fiber.New()
	handler := WriteOrder(m)
	app.Post(route, handler)

	request := httptest.NewRequest(fiber.MethodPost, route, strings.NewReader(`{ "key": "123", "name": "test" }`))
	response, err := app.Test(request, -1)

	assert.Equal(t, fiber.StatusOK, response.StatusCode, "http status code not equal")
	assert.Equal(t, nil, err, "err not nil")
}
