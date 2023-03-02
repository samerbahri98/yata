// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.16.0
// source: 1676588454_todo.sql

package entities

import (
	"context"
	"database/sql"
)

const createTodo = `-- name: CreateTodo :one
INSERT INTO todos (id, title, description, user_id)
VALUES (?, ?, ?, ?)
RETURNING id, title, description, created_at, last_modified_at, user_id
`

type CreateTodoParams struct {
	ID          string
	Title       sql.NullString
	Description sql.NullString
	UserID      string
}

func (q *Queries) CreateTodo(ctx context.Context, arg CreateTodoParams) (Todo, error) {
	row := q.db.QueryRowContext(ctx, createTodo,
		arg.ID,
		arg.Title,
		arg.Description,
		arg.UserID,
	)
	var i Todo
	err := row.Scan(
		&i.ID,
		&i.Title,
		&i.Description,
		&i.CreatedAt,
		&i.LastModifiedAt,
		&i.UserID,
	)
	return i, err
}

const deleteTodo = `-- name: DeleteTodo :exec
DELETE FROM todos
WHERE id = ?
`

func (q *Queries) DeleteTodo(ctx context.Context, id string) error {
	_, err := q.db.ExecContext(ctx, deleteTodo, id)
	return err
}

const findTodoByUserId = `-- name: FindTodoByUserId :many
SELECT id,
  title,
  description,
  created_at,
  last_modified_at,
  user_id
FROM todos
WHERE user_id = ?
ORDER BY ?
`

func (q *Queries) FindTodoByUserId(ctx context.Context, userID string) ([]Todo, error) {
	rows, err := q.db.QueryContext(ctx, findTodoByUserId, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Todo
	for rows.Next() {
		var i Todo
		if err := rows.Scan(
			&i.ID,
			&i.Title,
			&i.Description,
			&i.CreatedAt,
			&i.LastModifiedAt,
			&i.UserID,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getTodo = `-- name: GetTodo :one
SELECT id,
  title,
  description,
  created_at,
  last_modified_at,
  user_id
FROM todos
WHERE id = ?
LIMIT 1
`

func (q *Queries) GetTodo(ctx context.Context, id string) (Todo, error) {
	row := q.db.QueryRowContext(ctx, getTodo, id)
	var i Todo
	err := row.Scan(
		&i.ID,
		&i.Title,
		&i.Description,
		&i.CreatedAt,
		&i.LastModifiedAt,
		&i.UserID,
	)
	return i, err
}

const listTodos = `-- name: ListTodos :many
SELECT id,
  title,
  description,
  created_at,
  last_modified_at,
  user_id
FROM todos
ORDER BY ?
`

func (q *Queries) ListTodos(ctx context.Context) ([]Todo, error) {
	rows, err := q.db.QueryContext(ctx, listTodos)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Todo
	for rows.Next() {
		var i Todo
		if err := rows.Scan(
			&i.ID,
			&i.Title,
			&i.Description,
			&i.CreatedAt,
			&i.LastModifiedAt,
			&i.UserID,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const updateTodo = `-- name: UpdateTodo :one
UPDATE todos
SET title = ?,
  description = ?,
  last_modified_at = CURRENT_TIMESTAMP
WHERE id = ?
RETURNING id, title, description, created_at, last_modified_at, user_id
`

type UpdateTodoParams struct {
	Title       sql.NullString
	Description sql.NullString
	ID          string
}

func (q *Queries) UpdateTodo(ctx context.Context, arg UpdateTodoParams) (Todo, error) {
	row := q.db.QueryRowContext(ctx, updateTodo, arg.Title, arg.Description, arg.ID)
	var i Todo
	err := row.Scan(
		&i.ID,
		&i.Title,
		&i.Description,
		&i.CreatedAt,
		&i.LastModifiedAt,
		&i.UserID,
	)
	return i, err
}
