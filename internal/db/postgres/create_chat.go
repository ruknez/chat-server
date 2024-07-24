package postgres

import (
	"context"

	sq "github.com/Masterminds/squirrel"
	"github.com/pkg/errors"
	"github.com/ruknez/chat-server/internal/domain"
)

// CreateChat создает чат.
func (s *service) CreateChat(ctx context.Context, massage domain.Chat) (int64, error) {
	builderInsert := sq.Insert(chatDBName).
		PlaceholderFormat(sq.Dollar).
		Columns("name", "user_ids").
		Values("testName", massage.UserIDs).
		Suffix("RETURNING id")

	query, args, err := builderInsert.ToSql()
	if err != nil {
		return 0, errors.Wrap(err, "failed to builder query")
	}

	var id int64
	err = s.conn.QueryRow(ctx, query, args...).Scan(&id)
	if err != nil {
		return 0, errors.Wrap(err, "failed to save message")
	}

	return id, nil
}
