package cmd

import (
	"github.com/samerbahri98/yata/internal/config"
	"github.com/samerbahri98/yata/internal/server"
)

func Execute() {
	v := config.Load()
	server.Bootstrap(&v)
}
