-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS Chats
(
    id       bigserial PRIMARY KEY,
    name     TEXT NOT NULL,
    user_ids bigint[]
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS Chats;
-- +goose StatementEnd
