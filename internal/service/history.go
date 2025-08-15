package service

import (
	"context"

	"github.com/Users/dilperez/Documents/clientes_grupo_proteger/internal"
)

func NewHistoryDefault(rp internal.HistoryRepository) *HistoryDefault {
	return &HistoryDefault{
		rp: rp,
	}
}

type HistoryDefault struct {
	rp internal.HistoryRepository
}

func (s *HistoryDefault) SaveHistory(ctx context.Context, history *internal.History) (err error) {
	err = s.rp.SaveHistory(ctx, history)
	if err != nil {
		return err
	}
	return
}
