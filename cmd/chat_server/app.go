package main

import (
	"fmt"

	"github.com/ruknez/chat-server/internal/api/chat"
	"github.com/ruknez/chat-server/internal/app"
	"github.com/ruknez/chat-server/internal/app/grpc_transport"
)

func main() {
	chatV1 := chat.NewChatService()
	transport, err := grpc_transport.NewTransportService()
	if err != nil {
		panic(fmt.Errorf("creating transport service: %w", err))
	}

	if err = app.StartGrpcServer(transport, chatV1); err != nil {
		panic(fmt.Errorf("starting grpc server: %w", err))
	}
}
