package postgres

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v4"
	"github.com/pkg/errors"
	"github.com/ruknez/chat-server/internal/db/postgres/entity"
)

type service struct {
	conn *pgx.Conn
}

// NewPostgresDB возвращает объект подключения к постгре.
func NewPostgresDB(ctx context.Context, config entity.PostgresConfig) (*service, error) {
	// Создаем соединение с базой данных
	dbDSN := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=%s", config.Host, config.Port, config.UserName, config.Password, config.DBName, "disable")
	con, err := pgx.Connect(ctx, dbDSN)
	if err != nil {
		return nil, errors.Wrap(err, "failed to connect to database")
	}

	err = con.Ping(ctx)
	if err != nil {
		return nil, errors.Wrap(err, "failed to Ping")
	}

	return &service{conn: con}, nil
}
