-- +goose Up
CREATE TABLE event (id uuid PRIMARY KEY, name text NOT NULL);
-- +goose Down
DROP TABLE event