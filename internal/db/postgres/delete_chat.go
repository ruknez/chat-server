package postgres

import (
	"context"

	sq "github.com/Masterminds/squirrel"
	"github.com/pkg/errors"
	"github.com/ruknez/chat-server/internal/domain"
)

// DeleteChat удаляет чат.
func (s *service) DeleteChat(ctx context.Context, massage domain.Chat) error {
	builderInsert := sq.Delete(chatDBName).
		PlaceholderFormat(sq.Dollar).
		Where(sq.Eq{"id": massage.ChatID})

	query, args, err := builderInsert.ToSql()
	if err != nil {
		return errors.Wrap(err, "failed to builder query")
	}

	_, err = s.conn.Exec(ctx, query, args...)
	if err != nil {
		return errors.Wrap(err, "failed to Exec")
	}

	return nil
}
