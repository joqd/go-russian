package port

import (
	"context"

	"github.com/joqd/go-russian/internal/core/domain"
)

type WordPersistent interface {
	GetByID(ctx context.Context, id string) (*domain.Word, error)
	GetByBare(ctx context.Context, bare string) (*domain.Word, error)
	Create(ctx context.Context, word *domain.Word) (id string, err error)
	DeleteByID(ctx context.Context, id string) error
	DeleteByBare(ctx context.Context, bare string) error
}

type WordCache interface {
	Set(ctx context.Context, word *domain.Word) error
	Get(ctx context.Context, id string) (*domain.Word, error)
	Del(ctx context.Context, id string) error
}

type WordUsecase interface {
	GetByID(ctx context.Context, id string) (*domain.Word, error)
	GetByBare(ctx context.Context, word string) (*domain.Word, error)
	Create(ctx context.Context, word *domain.Word) (*domain.Word, error)
	DeleteByID(ctx context.Context, id string) error
	DeleteByBare(ctx context.Context, bare string) error
}
