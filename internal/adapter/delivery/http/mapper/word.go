package mapper

import (
	"github.com/joqd/ruskee/internal/adapter/delivery/http/response"
	"github.com/joqd/ruskee/internal/core/domain"
)

func WordToRetrievedWord(word *domain.Word) *response.RetrievedWord {
	return &response.RetrievedWord{
		ID:       word.ID,
		Bare:     word.Bare,
		Accented: word.Accented,
		Type:     word.Type,
		Level:    word.Level,
	}
}
