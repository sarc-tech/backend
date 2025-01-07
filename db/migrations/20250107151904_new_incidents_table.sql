-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS incidents (
 id SERIAL PRIMARY KEY,
 is_training boolean DEFAULT false NOT NULL,
 region VARCHAR(255),
 fio VARCHAR(255),
 status VARCHAR(255),
 date TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS incidents;
-- +goose StatementEnd
