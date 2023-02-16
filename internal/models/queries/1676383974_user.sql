-- name: GetUser :one
SELECT id,
  username,
  email,
  created_at,
  last_modified_at
FROM users
WHERE id = ?
LIMIT 1;
-- name: ListUsers :many
SELECT id,
  username,
  email,
  created_at,
  last_modified_at
FROM users
ORDER BY ?;
-- name: CreateUser :one
INSERT INTO users (id, username, email)
VALUES (?, ?, ?)
RETURNING *;
-- name: UpdateUser :one
UPDATE users
SET username = ?,
  email = ?,
  last_modified_at = CURRENT_TIMESTAMP
WHERE id = ?
RETURNING *;
-- name: DeleteUser :exec
DELETE FROM users
WHERE id = ?;
