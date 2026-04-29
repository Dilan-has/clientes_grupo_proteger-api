package no_sql

import (
	"context"

	"github.com/Users/dilperez/Documents/clientes_grupo_proteger/internal"
	"go.mongodb.org/mongo-driver/bson"
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
	if len(h.History) == 0 {
		return nil
	}

	newHistoryItem := h.History[0]

	// 1. Intentar actualizar una entrada existente que coincida con entry_date pero no tenga end_date
	if newHistoryItem.End_date != "" {
		filter := bson.M{
			"_id": h.Id,
			"history": bson.M{
				"$elemMatch": bson.M{
					"entry_date": newHistoryItem.Entry_date,
					"end_date":   bson.M{"$exists": false},
				},
			},
		}
		update := bson.M{
			"$set": bson.M{
				"history.$.end_date": newHistoryItem.End_date,
				"name":               h.Name,
				"cc":                 h.Cc,
			},
		}
		result, err := r.collection.UpdateOne(ctx, filter, update)
		if err == nil && result.ModifiedCount > 0 {
			return nil // Actualizado con éxito
		}
	}

	// 2. Si no se encontró una entrada para actualizar (o no hay end_date), se agrega una nueva
	filter := bson.M{"_id": h.Id}
	update := bson.M{
		"$set": bson.M{
			"name": h.Name,
			"cc":   h.Cc,
		},
		"$push": bson.M{
			"history": newHistoryItem,
		},
	}

	opts := options.Update().SetUpsert(true)
	_, err := r.collection.UpdateOne(ctx, filter, update, opts)
	return err
}

func (r *historialRepository) FindByID(ctx context.Context, id int) (*internal.History, error) {
	var h internal.History
	err := r.collection.FindOne(ctx, bson.M{"_id": id}).Decode(&h)
	if err != nil {
		return nil, err
	}
	return &h, nil
}
