package chat

import (
	"context"
	"fmt"
	"log/slog"

	"github.com/pkg/errors"
	"github.com/ruknez/chat-server/internal/domain"
	desc "github.com/ruknez/chat-server/pkg/chat_v1"
)

// Create создает чат.
func (c *Service) Create(ctx context.Context, req *desc.CreateRequest) (*desc.CreateResponse, error) {
	slog.Info("Create method", "req", fmt.Sprintf("%+v", req))
	id, err := c.crudService.CreateChat(ctx, domain.Chat{UserIDs: req.UserId})
	if err != nil {
		return nil, errors.Wrap(err, "CreateChat failed")
	}

	return &desc.CreateResponse{Id: id}, nil
}
