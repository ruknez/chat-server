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

// SendMessage посылает сообщение в чат.
func (c *Service) SendMessage(ctx context.Context, req *desc.SendMessageRequest) (*emptypb.Empty, error) {
	slog.Info("SendMessage method", "req", fmt.Sprintf("%+v", req))

	id, err := c.crudService.SaveMassage(ctx, domain.Massage{ChatID: req.ChatId, UserID: req.UserId, Text: req.Text})
	if err != nil {
		return nil, errors.Wrap(err, "failed to save massage")
	}

	slog.Info("SendMessage method massage", "id", id)
	return nil, nil
}
