-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS incidents(
    id serial PRIMARY KEY,
    is_training boolean DEFAULT FALSE NOT NULL,
    region varchar(255),
    fio varchar(255),
    status varchar(255),
    date timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE IF NOT EXISTS statuses(
    id serial PRIMARY KEY,
    name varchar(255),
);
-- +goose StatementEnd
-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS incidents;
DROP TABLE IF EXISTS statuses;

-- +goose StatementEnd
