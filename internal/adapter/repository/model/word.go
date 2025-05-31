package model

import (
	"github.com/joqd/ruskee/internal/core/domain"

	"go.mongodb.org/mongo-driver/v2/bson"
)

type WordDocument struct {
	ID       bson.ObjectID `bson:"_id"`
	Bare     string        `bson:"bare"`
	Accented string        `bson:"accented"`
	Type     *string       `bson:"type,omitempty"`
	Level    *string       `bson:"level,omitempty"`
	Disable  bool          `bson:"disable"`
}

func (w *WordDocument) ToDomain() *domain.Word {
	return &domain.Word{
		ID:       w.ID.Hex(),
		Bare:     w.Bare,
		Accented: w.Accented,
		Type:     w.Type,
		Level:    w.Level,
		Disable:  w.Disable,
	}
}
