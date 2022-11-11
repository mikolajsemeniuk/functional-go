package cart

import (
	"context"
	"encoding/json"
	"time"

	"github.com/go-redis/redis/v8"
	"github.com/gofiber/fiber/v2"
	"github.com/kelseyhightower/envconfig"
)

// Route
func Route(s *redis.Client, r fiber.Router) error {
	if err := envconfig.Process("", &configuration); err != nil {
		return err
	}

	r.Get("/cart/:key", ReadOrder(s))
	r.Post("/cart", WriteOrder(s))
	return nil
}

// OrderReader
type OrderReader interface {
	Get(context.Context, string) *redis.StringCmd
}

// ReadOrder
//
// @Summary Read order
// @Schemes
// @Description Read order from redis
// @Tags order
// @Accept application/json
// @Param key path string true "key"
// @Success 200 {object} Order
// @Failure 404
// @Failure 500
// @Failure 503
// @Router /cart/{key} [get]
func ReadOrder(r OrderReader) fiber.Handler {
	read := RedisRead(r.Get)

	return func(c *fiber.Ctx) error {
		key := c.Params("key")
		order, err := read(key)

		switch err {
		case nil:
			return c.SendString(order)
		case redis.Nil:
			return fiber.NewError(fiber.StatusNotFound, "order with this key not found")
		case redis.ErrClosed:
			return fiber.NewError(fiber.StatusServiceUnavailable, err.Error())
		default:
			return fiber.NewError(fiber.StatusInternalServerError, err.Error())
		}
	}
}

// OrderWriter
type OrderWriter interface {
	Set(context.Context, string, interface{}, time.Duration) *redis.StatusCmd
}

// WriteOrder
//
// @Summary Write order
// @Schemes
// @Description Write order to redis
// @Tags order
// @Accept application/json
// @Param payload body Order true "body"
// @Success 200 {object} string
// @Success 400
// @Failure 503
// @Router /cart [post]
func WriteOrder(w OrderWriter) fiber.Handler {
	write := RedisWrite(w.Set)

	return func(c *fiber.Ctx) error {
		var order Order
		if err := json.Unmarshal(c.Body(), &order); err != nil {
			return fiber.NewError(fiber.StatusBadRequest, err.Error())
		}

		if err := write(order, 0); err != nil {
			return err
		}

		return c.SendString("cart updated")
	}
}
