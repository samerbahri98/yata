#! /bin/sh


lower=$(echo "$1" | tr '[A-Z]' '[a-z]';)
capital=$(echo "$1" | tr '[A-Z]' '[a-z]' | sed -e "s/\b\(.\)/\u\1/g";)


cat<< EOF >>internal/models/schema/$(date +%s)_$1.sql
-- +goose Up
CREATE TABLE $(echo $lower)s (
  id   varchar(16) UNIQUE PRIMARY KEY,
);

CREATE INDEX  $(echo $lower)_id on $(echo $lower)s (id)
-- +goose Down
DROP TABLE  $(echo $lower)s
EOF

cat<< EOF >>internal/models/queries/$(date +%s)_$1.sql
-- name: Get$(echo $capital) :one
SELECT * FROM $(echo $lower)s
WHERE id = ? LIMIT 1;

-- name: List$(echo $capital)s :many
SELECT * FROM $(echo $lower)s
ORDER BY name;

-- name: Create$(echo $capital) :one
INSERT INTO $(echo $lower)s (
  name, bio
) VALUES (
  ?, ?
)
RETURNING *;

-- name: Update$(echo $capital) :one
UPDATE $(echo $lower)s
SET
    ? = ?
WHERE id=?
RETURNING *;

-- name: Delete$(echo $capital) :exec
DELETE FROM $(echo $lower)s
WHERE id = ?;
EOF
