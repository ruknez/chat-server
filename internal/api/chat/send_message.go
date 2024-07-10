package chat

import (
	"context"
	"fmt"
	"log/slog"

	desc "github.com/ruknez/chat-server/pkg/chat_v1"

	"google.golang.org/protobuf/types/known/emptypb"
)

// SendMessage посылает сообщение в чат.
func (c *Service) SendMessage(_ context.Context, req *desc.SendMessageRequest) (*emptypb.Empty, error) {
	slog.Info("SendMessage method", "req", fmt.Sprintf("%+v", req))

	return nil, nil
}
