package chat

import (
	desc "github.com/ruknez/chat-server/pkg/chat_v1"
)

// ChatService Реализация ручек чата.
type ChatService struct {
	desc.UnimplementedChatV1Server
}

// NewChatService возвращает сервис реализующий хэндлеры.
func NewChatService() *ChatService {
	return &ChatService{}
}
