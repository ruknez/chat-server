package chat

import (
	"context"
	"fmt"
	"log/slog"

	"github.com/pkg/errors"
	"github.com/ruknez/chat-server/internal/domain"
	desc "github.com/ruknez/chat-server/pkg/chat_v1"

	"google.golang.org/protobuf/types/known/emptypb"
)

// Delete удаляет чат.
func (c *Service) Delete(ctx context.Context, req *desc.DeleteRequest) (*emptypb.Empty, error) {
	slog.Info("Delete method", "req", fmt.Sprintf("%+v", req))

	err := c.crudService.DeleteChat(ctx, domain.Chat{ChatID: req.Id})
	if err != nil {
		return nil, errors.Wrap(err, "DeleteChat failed")
	}

	return nil, nil
}
