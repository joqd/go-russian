package mapper

import (
	"github.com/joqd/go-russian/internal/adapter/delivery/http/request"
	"github.com/joqd/go-russian/internal/adapter/delivery/http/response"
	"github.com/joqd/go-russian/internal/core/domain"
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

func CreateWordToWord(createWord *request.CreateWord) *domain.Word {
	word := &domain.Word{
		Bare:     createWord.Bare,
		Accented: createWord.Accented,
		Type:     createWord.Type,
		Level:    createWord.Level,
	}

	if createWord.Disable == nil {
		word.Disable = false
	} else {
		word.Disable = *createWord.Disable
	}

	return word
}
