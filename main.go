package main

import (
	"fun/cart"
	"log"

	_ "fun/docs"

	"github.com/go-redis/redis/v8"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/swagger"
	"github.com/kelseyhightower/envconfig"
)

type Settings struct {
	Listen string `envconfig:"LISTEN" default:":5000"`
	Redis  string `envconfig:"REDIS" default:"localhost:6379"`
}

// @title           Cart
// @version         1.0
// @description     Cart API
// @BasePath /
// @schemes http https
func main() {
	var settings Settings
	if err := envconfig.Process("", &settings); err != nil {
		log.Fatal(err)
	}

	storage := redis.NewClient(&redis.Options{Addr: settings.Redis})
	router := fiber.New()

	if err := cart.Route(storage, router); err != nil {
		log.Fatal(err)
	}

	router.Get("/swagger/*", swagger.HandlerDefault)
	router.Use(func(c *fiber.Ctx) error { return c.Status(fiber.StatusNotFound).Redirect("/swagger/index.html") })

	if err := router.Listen(settings.Listen); err != nil {
		log.Fatal(err)
	}
}
