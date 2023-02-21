package server

import (
	"context"
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/samerbahri98/yata/internal/repository"
	"github.com/samerbahri98/yata/internal/router/api/v1"
	"github.com/spf13/viper"
)

type ServerProps struct {
	Viper      *viper.Viper
	Repository *repository.Repository
	Context    context.Context
}

func Bootstrap(s *ServerProps) {
	app := fiber.New()
	app.Use(logger.New())
	v1Props := &api.V1Props{
		Repository: s.Repository,
		Context:    s.Context,
	}
	api := api.New(v1Props)
	app.Mount("/api/v1", api)
	port := s.Viper.GetString("server.port")
	address := s.Viper.GetString("server.address")
	listener := fmt.Sprintf("%s:%s", address, port)
	if err := app.Listen(listener); err != nil {
		log.Fatalln(err)
	}
}
