package cmd

import (
	"log"

	"github.com/samerbahri98/yata/internal/config"
	"github.com/samerbahri98/yata/internal/server"
)

func Execute() {
	err, v := config.Load()
	if err != nil {
		log.Panicln(err)
	}
	server.Bootstrap(v)
}
