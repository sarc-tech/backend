-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS users (
 id SERIAL PRIMARY KEY,
 login VARCHAR(255) NOT NULL,
 ver INTEGER NOT NULL DEFAULT '0',
 email VARCHAR(255) NOT NULL,
 password bytea NOT NULL,
 created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
 updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE IF NOT EXISTS sessions (
 id VARCHAR(32) NOT NULL,
 user_id INTEGER NOT NULL
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS users;

DROP TABLE IF EXISTS sessions;
-- +goose StatementEnd