-- +goose Up
CREATE TABLE users (
  id   VARCHAR(16) UNIQUE PRIMARY KEY,
  username TEXT,
  email TEXT,
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  last_modified_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE INDEX  user_id on users (id);
CREATE INDEX user_created_at on users (created_at);
CREATE INDEX user_modified_at on users (last_modified_at);

-- +goose Down
DROP TABLE  users;
