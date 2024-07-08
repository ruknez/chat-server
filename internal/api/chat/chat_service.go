package chat

import (
	desc "github.com/ruknez/chat-server/pkg/chat_v1"
)

// Service Реализация ручек чата.
type Service struct {
	desc.UnimplementedChatV1Server
}

// NewChatService возвращает сервис реализующий хэндлеры.
func NewChatService() *Service {
	return &Service{}
}
