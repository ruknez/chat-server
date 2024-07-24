-- +goose Up
-- +goose StatementBegin
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";
CREATE TABLE IF NOT EXISTS Messages
(
    id      UUID DEFAULT uuid_generate_v4(),
    chat_id bigint REFERENCES Chats (id) ON DELETE CASCADE,
    user_id bigint NOT NULL,
    text    TEXT,
    PRIMARY KEY (id, chat_id)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE  IF EXISTS Messages;
-- +goose StatementEnd
