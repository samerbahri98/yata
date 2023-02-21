package config

import (
	"context"

	"github.com/samerbahri98/yata/internal/server"
)

func Load(ctx context.Context) (err error, c *server.ServerProps) {
	err, v := loadViper()
	if err != nil {
		return err, nil
	}

	err, r := loadDB(ctx, v)
	if err != nil {
		return err, nil
	}
	serverProps := &server.ServerProps{
		Viper:      v,
		Repository: r,
		Context:    ctx,
	}
	return nil, serverProps

}
