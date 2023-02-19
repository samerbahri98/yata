package repository

import (
	"github.com/samerbahri98/yata/internal/models/entities"
)

type Repository struct {
	db      *entities.DBTX
	queries *entities.Queries
}

func New(db *entities.DBTX) *Repository {

	return &Repository{
		db:      db,
		queries: entities.New(*db),
	}
}
