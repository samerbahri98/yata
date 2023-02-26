package apiV1

import (
	"github.com/gofiber/fiber/v2"
	"github.com/samerbahri98/yata/internal/models/entities"
	"github.com/samerbahri98/yata/internal/repository"
)

var userRepository *repository.UserRepository = nil

func (a *APIBuilder) user(r fiber.Router) {
	if userRepository == nil {
		userRepository = &repository.UserRepository{
			Repository: a.repository,
		}
	}
	r.Post("/", a.createUser)
	r.Get("/:id", a.getUser)
}

func (a *APIBuilder) createUser(c *fiber.Ctx) error {
	userProps := new(entities.CreateUserParams)
	if err := c.BodyParser(userProps); err != nil {
		return err
	}
	user, err := userRepository.CreateUser(a.context, userProps.Username, userProps.Email)
	if err != nil {
		return err
	}
	return c.JSON(user)
}
func (a *APIBuilder) getUser(c *fiber.Ctx) error {
	id := c.AllParams()["id"]
	user, err := userRepository.GetUser(a.context, id)
	if err != nil {
		return err
	}
	return c.JSON(user)
}
