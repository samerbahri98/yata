package cmd

import (
	"context"
	"log"

	"github.com/samerbahri98/yata/internal/config"
	"github.com/samerbahri98/yata/internal/server"
)

func Execute() {
	ctx := context.Background()
	err, serverProps := config.Load(ctx)
	if err != nil {
		log.Panicln(err)
	}
	server.Bootstrap(serverProps)
}
