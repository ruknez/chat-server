package chat

import (
	"context"
	"fmt"
	"log/slog"

	desc "github.com/ruknez/chat-server/pkg/chat_v1"
)

// Create создает чат.
func (c *Service) Create(_ context.Context, req *desc.CreateRequest) (*desc.CreateResponse, error) {
	slog.Info("Create method", "req", fmt.Sprintf("%+v", req))

	return nil, nil
}
