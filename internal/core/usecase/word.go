package usecase

import (
	"context"

	"github.com/joqd/ruskee/internal/core/domain"
	"github.com/joqd/ruskee/internal/core/port"
)

type wordUsecase struct {
	persistent port.WordPersistent
	cache      port.WordCache
	xlog       port.Logger
}

func NewWordUsecase(persistent port.WordPersistent, cache port.WordCache, xlog port.Logger) port.WordUsecase {
	return &wordUsecase{
		persistent: persistent,
		cache:      cache,
		xlog:       xlog,
	}
}

func (w *wordUsecase) GetByID(ctx context.Context, id string) (*domain.Word, error) {
	word, err := w.cache.Get(ctx, id)
	if err == nil {
		return word, nil
	}

	word, err = w.persistent.GetByID(ctx, id)
	if err != nil {
		return nil, err
	}

	w.cache.Set(ctx, word)

	return word, nil
}
