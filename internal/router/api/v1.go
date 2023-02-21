package api

import (
	"context"

	"github.com/gofiber/fiber/v2"
	"github.com/samerbahri98/yata/internal/models/entities"
	"github.com/samerbahri98/yata/internal/repository"
)

type V1Props struct {
	Repository *repository.Repository
	Context    context.Context
}

var userRepository *repository.UserRepository = nil

func New(v1Props *V1Props) *fiber.App {
	if userRepository == nil {
		userRepository = &repository.UserRepository{
			Repository: v1Props.Repository,
		}
	}
	api := fiber.New()
	api.Use(func(c *fiber.Ctx) error {
		c.Accepts("application/json")
		return c.Next()
	})
	userApi := api.Group("/user")
	userApi.Post("/", func(c *fiber.Ctx) error {
		userProps := new(entities.CreateUserParams)
		if err := c.BodyParser(userProps); err != nil {
			return err
		}
		user, err := userRepository.CreateUser(v1Props.Context, userProps.Username, userProps.Email)
		if err != nil {
			return err
		}
		return c.JSON(user)
	})

	return api
}
