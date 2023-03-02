package repository

import (
	"context"

	"github.com/google/uuid"
	"github.com/samerbahri98/yata/internal/models/entities"
)

type UserRepository struct {
	*Repository
}

func (r *UserRepository) GetUser(ctx context.Context, id string) (entities.User, error) {
	return r.queries.GetUser(ctx, id)
}

func (r *UserRepository) FindUserByEmail(ctx context.Context, email string) (entities.User, error) {
	return r.queries.FindUserByEmail(ctx, email)
}

func (r *UserRepository) FindUserByUsername(ctx context.Context, username string) (entities.User, error) {
	return r.queries.FindUserByUsername(ctx, username)
}

func (r *UserRepository) CreateUser(ctx context.Context, username string, email string) (entities.User, error) {
	var p entities.CreateUserParams

	p.ID = uuid.New().String()
	p.Username = username
	p.Email = email
	return r.queries.CreateUser(ctx, p)

}

func (r *UserRepository) DeleteUser(ctx context.Context, id string) error {
	return r.queries.DeleteUser(ctx, id)
}
