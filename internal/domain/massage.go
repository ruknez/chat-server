package domain

import (
	"google.golang.org/protobuf/types/known/timestamppb"
)

// Massage общая структура описывающая сообщение.
type Massage struct {
	From      string
	Text      string
	Timestamp *timestamppb.Timestamp
	UserID    int64
	ChatID    int64
}
