package apiV1

import (
	"context"

	"github.com/gofiber/fiber/v2"
	"github.com/samerbahri98/yata/internal/repository"
)

type APIBuilder struct {
	api        *fiber.App
	repository *repository.Repository
	context    context.Context
}

func New() *APIBuilder {
	api := fiber.New()
	api.Use(func(c *fiber.Ctx) error {
		c.Accepts("application/json")
		return c.Next()
	})
	return &APIBuilder{
		api: api,
	}
}

func (a *APIBuilder) AddRepository(r *repository.Repository) *APIBuilder {
	a.repository = r
	return a
}

func (a *APIBuilder) AddContext(ctx context.Context) *APIBuilder {
	a.context = ctx
	return a
}

func (a *APIBuilder) Build() *APIBuilder {
	a.api.Route("/user", a.user)
	return a
}

func (a *APIBuilder) GetAPI() *fiber.App {
	return a.api
}
