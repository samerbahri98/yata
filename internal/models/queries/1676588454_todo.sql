-- name: GetTodo :one
SELECT id,
  title,
  description,
  created_at,
  last_modified_at,
  user_id
FROM todos
WHERE id = ?
LIMIT 1;
-- name: ListTodos :many
SELECT id,
  title,
  description,
  created_at,
  last_modified_at,
  user_id
FROM todos
ORDER BY ?;
-- name: FindByUser :many
SELECT id,
  title,
  description,
  created_at,
  last_modified_at,
  user_id
FROM todos
WHERE user_id = ?
ORDER BY ?;
-- name: CreateTodo :one
INSERT INTO todos (id, title, description, user_id)
VALUES (?, ?, ?, ?)
RETURNING *;
-- name: UpdateTodo :one
UPDATE todos
SET title = ?,
  description = ?,
  last_modified_at = CURRENT_TIMESTAMP
WHERE id = ?
RETURNING *;
-- name: DeleteTodo :exec
DELETE FROM todos
WHERE id = ?;
