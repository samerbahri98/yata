package api

import (
	"github.com/gofiber/fiber/v2"
	"github.com/samerbahri98/yata/internal/models/entities"
	"github.com/samerbahri98/yata/internal/repository"
)

var userRepository *repository.UserRepository = nil

func userApi(v1Props *V1Props) func(fiber.Router) {
	if userRepository == nil {
		userRepository = &repository.UserRepository{
			Repository: v1Props.Repository,
		}
	}
	return func(r fiber.Router) {
		r.Post("/", createUser(v1Props))
		r.Get("/:id", getUser(v1Props))
	}
}

func createUser(v1Props *V1Props) func(c *fiber.Ctx) error {
	userProps := new(entities.CreateUserParams)

	return func(c *fiber.Ctx) error {
		if err := c.BodyParser(userProps); err != nil {
			return err
		}
		user, err := userRepository.CreateUser(v1Props.Context, userProps.Username, userProps.Email)
		if err != nil {
			return err
		}
		c.JSON(user)
		return nil
	}
}

func getUser(v1Props *V1Props) func(c *fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		id := c.AllParams()["id"]
		user, err := userRepository.GetUser(v1Props.Context, id)
		if err != nil {
			return err
		}
		c.JSON(user)
		return nil
	}
}
