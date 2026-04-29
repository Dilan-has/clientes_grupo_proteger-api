package internal

import "context"

type History struct {
	Id      interface{}   `bson:"_id" json:"id"`
	Name    string        `bson:"name" json:"name"`
	Cc      string        `bson:"cc" json:"cc"`
	History []DateHistory `bson:"history" json:"history"`
}

type DateHistory struct {
	Entry_date string `bson:"entry_date,omitempty" json:"entry_date,omitempty"`
	End_date   string `bson:"end_date,omitempty" json:"end_date,omitempty"`
}

type HistoryRepository interface {
	SaveHistory(ctx context.Context, h *History) error
	FindByID(ctx context.Context, id int) (*History, error)
}

type HistoryService interface {
	SaveHistory(ctx context.Context, history *History) (err error)
	FindByID(ctx context.Context, id int) (*History, error)
}
