-- +goose Up
CREATE TABLE todos (
  id   varchar(16) UNIQUE PRIMARY KEY,
  title TEXT,
  description TEXT,
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  last_modified_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  user_id VARCHAR(16),
  FOREIGN KEY(user_id) REFERENCES users(id)
);

CREATE INDEX todo_id on todos (id)

-- +goose Down
DROP TABLE  todos
