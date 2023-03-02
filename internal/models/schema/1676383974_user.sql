-- +goose Up
CREATE TABLE users (
  id   VARCHAR(16) UNIQUE PRIMARY KEY,
  username TEXT UNIQUE NOT NULL,
  email TEXT UNIQUE NOT NULL,
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  last_modified_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE INDEX user_id on users (id);

-- +goose Down
DROP TABLE  users;
