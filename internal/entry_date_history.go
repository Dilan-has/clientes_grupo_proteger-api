package internal

import "context"

type History struct {
	Id      interface{}   `bson:"_id"` // puede ser int, string, u ObjectID
	Name    string        `bson:"name"`
	Cc      string        `bson:"cc"`
	History []DateHistory `bson:"history"`
}

type DateHistory struct {
	Entry_date string `bson:"entry_date"`
	End_date   string `bson:"end_date,omitempty"`
}

type HistoryRepository interface {
	SaveHistory(ctx context.Context, h *History) error
}

type HistoryService interface {
	SaveHistory(ctx context.Context, history *History) (err error)
}
