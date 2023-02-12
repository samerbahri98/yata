package server

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/spf13/viper"
)

func Bootstrap(v *viper.Viper) {
	app := fiber.New()
	app.Use(logger.New())
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})
	port := v.GetString("server.port")
	address := v.GetString("server.address")
	listener := fmt.Sprintf("%s:%s", address, port)
	app.Listen(listener)
}
