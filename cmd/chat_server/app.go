package main

import (
	"fmt"

	"chat-server/internal/api/chat"
	"chat-server/internal/app"
	"chat-server/internal/app/grpc_transpotr"
)

func main() {
	chatV1 := chat.NewChatService()
	transport, err := grpc_transpotr.NewTransportService()
	if err != nil {
		panic(fmt.Sprint("Error creating transport service: ", err))
	}

	if err = app.StartGrpcServer(transport, chatV1); err != nil {
		panic(fmt.Sprint("Error starting grpc server: ", err))
	}
}
