package builder

import (
	"context"
	"log"

	"github.com/samerbahri98/yata/internal/config"
	"github.com/samerbahri98/yata/internal/db"
	"github.com/samerbahri98/yata/internal/repository"
	"github.com/samerbahri98/yata/internal/server"
)

func BuildServer() {

	app := server.New()

	// Context
	ctx := context.Background()
	app.AddContext(ctx)

	// Viper
	v, err := config.Load(ctx)
	if err != nil {
		log.Panicln(err)
	}
	app.AddConfig(v)

	// Repository
	dbtx, err := db.Load(ctx, v)
	if err != nil {
		log.Panicln(err)
	}
	r := repository.New(&dbtx)
	app.AddRepository(r)

	// Start
	if err := app.Start(); err != nil {
		log.Panicln(err)
	}
}
