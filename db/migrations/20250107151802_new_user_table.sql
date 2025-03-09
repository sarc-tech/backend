-- +goose Up
-- +goose StatementBegin
CREATE TYPE gender as enum ('male', 'female');

CREATE TABLE IF NOT EXISTS users(
    id serial PRIMARY KEY,
    yandex_id VARCHAR(30) NOT NULL,
    surname VARCHAR(255),
    name VARCHAR(255) NOT NULL,
    patronymic VARCHAR(255),
    call_sign VARCHAR(255),
    gender gender,
    birthdate DATE,
    vk VARCHAR(255),
    telegram  VARCHAR(255),
    email varchar(255),
    phone varchar(255),
    approval boolean,
    created_at timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE IF NOT EXISTS sessions(
    id varchar(59) NOT NULL,
    user_id integer NOT NULL
);

-- +goose StatementEnd
-- +goose Down
-- +goose StatementBegin
DROP TYPE gender;

DROP TABLE IF EXISTS users;

DROP TABLE IF EXISTS sessions;

-- +goose StatementEnd
