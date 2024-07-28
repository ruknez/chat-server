package postgres

import (
	"context"

	sq "github.com/Masterminds/squirrel"
	"github.com/pkg/errors"
	"github.com/ruknez/chat-server/internal/domain"
)

// SaveMassage сохраняет сообщение.
func (s *service) SaveMassage(ctx context.Context, massage domain.Massage) (int64, error) {
	builderInsert := sq.Insert(messageDBName).
		PlaceholderFormat(sq.Dollar).
		Columns("chat_id", "user_id", "text").
		Values(massage.ChatID, massage.UserID, massage.Text).
		Suffix("RETURNING id")

	query, args, err := builderInsert.ToSql()
	if err != nil {
		return 0, errors.Wrap(err, "failed to builder query")
	}

	var id int64
	err = s.conn.QueryRow(ctx, query, args...).Scan(&id)
	if err != nil {
		return 0, errors.Wrap(err, "failed to save chat")
	}

	return id, nil
}
