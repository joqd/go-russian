package usecase

import (
	"context"

	"github.com/joqd/slovo/internal/core/domain"
	"github.com/joqd/slovo/internal/core/port"
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

func (w *wordUsecase) Create(ctx context.Context, word *domain.Word) (*domain.Word, error) {
	oid, err := w.persistent.Create(ctx, word)
	if err != nil {
		return nil, err
	}

	word.ID = oid
	w.cache.Set(ctx, word)

	return word, nil
}

func (w *wordUsecase) GetByBare(ctx context.Context, bare string) (*domain.Word, error) {
	word, err := w.persistent.GetByBare(ctx, bare)
	if err != nil {
		return nil, err
	}

	if word.Disable {
		return nil, domain.ErrDataNotFound
	}

	return word, nil
}

func (w *wordUsecase) DeleteByID(ctx context.Context, id string) error {
	err := w.persistent.DeleteByID(ctx, id)
	if err != nil {
		return err
	}

	w.cache.Del(ctx, id)

	return nil
}

func (w *wordUsecase) DeleteByBare(ctx context.Context, bare string) error {
	err := w.persistent.DeleteByBare(ctx, bare)
	if err != nil {
		return err
	}

	return nil
}
