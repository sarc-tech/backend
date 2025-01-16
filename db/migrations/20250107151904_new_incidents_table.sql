-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS incidents(
    id serial PRIMARY KEY,
    region varchar(255),
    fio varchar(255),
    statusId varchar(255),
    date timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE IF NOT EXISTS statuses(
    id serial PRIMARY KEY,
    name varchar(255)
);
-- +goose StatementEnd
-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS incidents;
DROP TABLE IF EXISTS statuses;

-- +goose StatementEnd
