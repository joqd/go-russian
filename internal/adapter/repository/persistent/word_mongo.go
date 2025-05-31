package persistent

import (
	"context"

	"github.com/joqd/ruskee/internal/adapter/repository/model"
	"github.com/joqd/ruskee/internal/core/domain"
	"github.com/joqd/ruskee/internal/core/port"
	"github.com/joqd/ruskee/pkg/mongodb"

	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

type wordPersistent struct {
	collection *mongo.Collection
	xlog       port.Logger
}

func NewWordRespository(mongodb mongodb.MongoDB, xlog port.Logger) port.WordPersistent {
	return &wordPersistent{
		collection: mongodb.Database.Collection("words"),
		xlog:       xlog,
	}
}

func (w *wordPersistent) GetByID(ctx context.Context, id string) (*domain.Word, error) {
	objectID, err := bson.ObjectIDFromHex(id)
	if err != nil {
		return nil, domain.ErrInvalidObjectID
	}

	var wordDocument model.WordDocument

	filter := bson.D{{Key: "_id", Value: objectID}}
	if err := w.collection.FindOne(ctx, filter).Decode(&wordDocument); err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, domain.ErrDataNotFound
		}

		w.xlog.Error("mongo get error, id=%s, err=%v", id, err)
		return nil, err
	}

	return wordDocument.ToDomain(), nil
}
