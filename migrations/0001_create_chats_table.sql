-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS Chats
(
    id       BIGSERIAL PRIMARY KEY,
    name     TEXT NOT NULL,
    user_ids BIGINT[]
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS Chats;
-- +goose StatementEnd
