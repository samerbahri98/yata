package api

import (
	"context"

	"github.com/gofiber/fiber/v2"
	"github.com/samerbahri98/yata/internal/repository"
)

type V1Props struct {
	Repository *repository.Repository
	Context    context.Context
}

func New(v1Props *V1Props) *fiber.App {
	api := fiber.New()
	api.Use(func(c *fiber.Ctx) error {
		c.Accepts("application/json")
		return c.Next()
	})
	api.Route("/user", userApi(v1Props))
	return api
}
