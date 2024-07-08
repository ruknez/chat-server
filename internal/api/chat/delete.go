package chat

import (
	"context"
	"fmt"
	"log/slog"

	desc "github.com/ruknez/chat-server/pkg/chat_v1"

	"google.golang.org/protobuf/types/known/emptypb"
)

func (c *ChatService) Delete(_ context.Context, req *desc.DeleteRequest) (*emptypb.Empty, error) {
	slog.Info("Delete method", "req", fmt.Sprintf("%+v", req))

	return nil, nil
}
