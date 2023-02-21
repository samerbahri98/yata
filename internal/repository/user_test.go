package repository

import (
	"context"
	"database/sql"
	"log"
	"os"
	"path"
	"testing"

	_ "github.com/mattn/go-sqlite3"
	"github.com/pressly/goose/v3"
	"github.com/samerbahri98/yata/internal/models/entities"
)

var ddl string

func setupTestDB(ctx context.Context) entities.DBTX {
	cwd, err := os.Getwd()
	if err != nil {
		log.Panicln(err)
	}
	dbpath := path.Join(cwd, "../../db/tmp.db")
	migpath := path.Join(cwd, "../models/schema")
	log.Println("\033[32m")
	log.Println(migpath)
	log.Println("\033[0m")

	// db
	db, err := sql.Open("sqlite3", dbpath)
	if err != nil {
		log.Panicln(err)
	}
	// goose
	if err := goose.SetDialect("sqlite3"); err != nil {
		log.Panicln(err)
	}

	if err := goose.Up(db, migpath); err != nil {

		log.Panicln(err)
	}

	// sqlc
	if _, err := db.ExecContext(ctx, ddl); err != nil {
		log.Panicln(err)
	}
	return entities.DBTX(db)
}

func setupTestUserRepo(ctx context.Context) UserRepository {
	db := setupTestDB(ctx)
	repo := New(&db)
	userRepository := UserRepository{
		repo,
	}
	return userRepository
}

func TestCreateUser(t *testing.T) {
	ctx := context.Background()
	userRepository := setupTestUserRepo(ctx)
	username := "samer"
	email := "samer.bahri@ieee.org"
	got, err := userRepository.CreateUser(ctx, username, email)
	if err != nil {
		log.Panicln(err)
	}
	if got.Username != username {
		t.Errorf("no")
	}
}
