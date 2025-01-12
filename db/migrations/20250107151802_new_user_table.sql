-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS users(
    id serial PRIMARY KEY,
    login VARCHAR(255) NOT NULL,
    ver integer NOT NULL DEFAULT '0',
    email varchar(255) NOT NULL,
    phone varchar(255) NOT NULL,
    password bytea NOT NULL,
    created_at timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE IF NOT EXISTS sessions(
    id varchar(32) NOT NULL,
    user_id integer NOT NULL
);

-- +goose StatementEnd
-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS users;

DROP TABLE IF EXISTS sessions;

-- +goose StatementEnd
