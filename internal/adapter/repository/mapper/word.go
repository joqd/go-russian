package mapper

import (
	"github.com/joqd/slovo/internal/adapter/repository/model"
	"github.com/joqd/slovo/internal/core/domain"
)

func WordToWordPayload(word *domain.Word) *model.WordPayload {
	return &model.WordPayload{
		Bare:     word.Bare,
		Accented: word.Accented,
		Type:     word.Type,
		Level:    word.Level,
		Disable:  word.Disable,
	}
}
