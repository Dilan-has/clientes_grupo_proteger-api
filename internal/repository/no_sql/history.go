package no_sql

import (
	"context"

	"github.com/Users/dilperez/Documents/clientes_grupo_proteger/internal"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type historialRepository struct {
	collection *mongo.Collection
}

func NewHistorialRepository(db *mongo.Database) *historialRepository {
	return &historialRepository{
		collection: db.Collection("history"),
	}
}

func (r *historialRepository) SaveHistory(ctx context.Context, h *internal.History) error {
	opts := options.InsertOne()
	_, err := r.collection.InsertOne(ctx, h, opts)
	return err
}
