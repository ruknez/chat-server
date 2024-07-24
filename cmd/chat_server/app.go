package main

import (
	"context"
	"fmt"

	"github.com/ruknez/chat-server/internal/api/chat"
	"github.com/ruknez/chat-server/internal/app"
	"github.com/ruknez/chat-server/internal/app/grpc_transport"
	"github.com/ruknez/chat-server/internal/config"
	"github.com/ruknez/chat-server/internal/db/postgres"
)

func main() {
	ctx := context.Background()
	conf, err := config.InitPgConfig()
	if err != nil || conf == nil {
		panic(fmt.Errorf("config.InitConfig: %w", err))
	}

	db, err := postgres.NewPostgresDB(ctx, *conf)
	if err != nil {
		panic(fmt.Errorf("NewPostgresDB: %w", err))
	}

	chatV1 := chat.NewChatService(db)
	transport, err := grpc_transport.NewTransportService()
	if err != nil {
		panic(fmt.Errorf("creating transport service: %w", err))
	}

	if err = app.StartGrpcServer(transport, chatV1); err != nil {
		panic(fmt.Errorf("starting grpc server: %w", err))
	}
}
