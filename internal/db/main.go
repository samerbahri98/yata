package db

import (
	"context"
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
	"github.com/samerbahri98/yata/internal/models/entities"
	"github.com/spf13/viper"
)

var ddl string

func Load(ctx context.Context, v *viper.Viper) (error, entities.DBTX) {
	conn := v.GetString("database.conn")
	driver := v.GetString("database.driver")
	db, err := sql.Open(driver, conn)
	if err != nil {
		return err, nil
	}
	if _, err := db.ExecContext(ctx, ddl); err != nil {
		return err, nil
	}
	return nil, entities.DBTX(db)
}
