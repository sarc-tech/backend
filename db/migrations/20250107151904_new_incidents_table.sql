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

-- +goose StatementEnd
-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS incidents;

-- +goose StatementEnd
