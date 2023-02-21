package config

import (
	"context"

	"github.com/samerbahri98/yata/internal/db"
	"github.com/samerbahri98/yata/internal/repository"
	"github.com/spf13/viper"
)

func loadDB(ctx context.Context, v *viper.Viper) (error, *repository.Repository) {
	err, dbtx := db.Load(ctx, v)
	if err != nil {
		return err, nil
	}
	r := repository.New(&dbtx)

	return nil, r
}
