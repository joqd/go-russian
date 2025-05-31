package port

import (
	"context"

	"github.com/joqd/ruskee/internal/core/domain"
)

type WordPersistent interface {
	GetByID(ctx context.Context, id string) (*domain.Word, error)
	Create(ctx context.Context, word *domain.Word) (id string, err error)
}

type WordCache interface {
	Set(ctx context.Context, word *domain.Word) error
	Get(ctx context.Context, id string) (*domain.Word, error)
	Del(ctx context.Context, id string) error
}

type WordUsecase interface {
	GetByID(ctx context.Context, id string) (*domain.Word, error)
	Create(ctx context.Context, word *domain.Word) (*domain.Word, error)
}
