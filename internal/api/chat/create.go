package chat

import (
	"context"
	"fmt"
	"log/slog"

	desc "github.com/ruknez/chat-server/pkg/chat_v1"
)

func (c *ChatService) Create(_ context.Context, req *desc.CreateRequest) (*desc.CreateResponse, error) {
	slog.Info("Create method", "req", fmt.Sprintf("%+v", req))

	return nil, nil
}
