package server

import (
	"context"
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/samerbahri98/yata/internal/repository"
	"github.com/spf13/viper"
)

type ServerBuilder struct {
	app        *fiber.App
	viper      *viper.Viper
	repository *repository.Repository
	context    context.Context
	listener   string
}

func New() *ServerBuilder {
	app := fiber.New()
	app.Use(logger.New())
	return &ServerBuilder{app: app}
}

func (b *ServerBuilder) AddConfig(v *viper.Viper) *ServerBuilder {
	b.viper = v
	port := b.viper.GetString("server.port")
	address := b.viper.GetString("server.address")
	b.listener = fmt.Sprintf("%s:%s", address, port)
	return b
}

func (b *ServerBuilder) AddContext(ctx context.Context) *ServerBuilder {
	b.context = ctx
	return b
}

func (b *ServerBuilder) AddRepository(r *repository.Repository) *ServerBuilder {
	b.repository = r
	return b
}

func (b *ServerBuilder) Start() error {
	b.buildAPIV1()
	return b.app.Listen(b.listener)
}
