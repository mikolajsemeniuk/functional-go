package main

import (
	"fun/cart"
	"log"

	_ "fun/docs"

	"github.com/go-redis/redis/v8"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/swagger"
)

// @title           Cart
// @version         1.0
// @description     Cart API
// @BasePath /
// @schemes http https
func main() {
	storage := redis.NewClient(&redis.Options{Addr: "localhost:6379"})
	router := fiber.New()
	fatal := log.Fatal

	config, err := cart.NewConfig()
	if err != nil {
		fatal(err)
	}

	cart.Route(storage, router)
	router.Get("/swagger/*", swagger.HandlerDefault)
	router.Use(func(c *fiber.Ctx) error { return c.Status(fiber.StatusNotFound).Redirect("/swagger/index.html") })

	if err := router.Listen(config.Listen); err != nil {
		fatal(err)
	}
}
