package chat

import (
	"context"

	"github.com/ruknez/chat-server/internal/domain"
	desc "github.com/ruknez/chat-server/pkg/chat_v1"
)

type workWithStorage interface {
	CreateChat(ctx context.Context, massage domain.Chat) (int64, error)
	DeleteChat(ctx context.Context, massage domain.Chat) error
	SaveMassage(ctx context.Context, massage domain.Massage) (int64, error)
}

// Service Реализация ручек чата.
type Service struct {
	desc.UnimplementedChatV1Server
	crudService workWithStorage
}

// NewChatService возвращает сервис реализующий хэндлеры.
func NewChatService(c workWithStorage) *Service {
	return &Service{crudService: c}
}
