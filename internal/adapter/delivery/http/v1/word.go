package v1

import (
	"errors"
	"net/http"

	"github.com/joqd/ruskee/internal/adapter/delivery/http/mapper"
	"github.com/joqd/ruskee/internal/adapter/delivery/http/response"
	_ "github.com/joqd/ruskee/internal/adapter/delivery/http/response/wrapper"
	"github.com/joqd/ruskee/internal/core/domain"
	"github.com/joqd/ruskee/internal/core/port"

	"github.com/gin-gonic/gin"
)

type wordHTTP struct {
	usecase port.WordUsecase
	xlog    port.Logger
}

func NewWordHTTP(usecase port.WordUsecase, xlog port.Logger) *wordHTTP {
	return &wordHTTP{
		usecase: usecase,
		xlog:    xlog,
	}
}

// Retrieve godoc
// @Summary      Get a word by ID
// @Description  Retrieve a word from the database using its ID
// @Tags         words
// @Param        id   path      string  true  "Word ID"
// @Success      200  {object}  wrapper.RetrievedWordWrapper
// @Failure      404  {object}  wrapper.ErrorNotFoundWrapper
// @Failure      400  {object}  wrapper.ErrorInvalidObjectIdWrapper
// @Failure      500  {object}  wrapper.ErrorInternalServerWrapper
// @Router       /api/v1/words/{id} [get]
func (w *wordHTTP) GetByID(c *gin.Context) {
	id := c.Param("id")

	word, err := w.usecase.GetByID(c.Request.Context(), id)
	if err != nil {
		if errors.Is(err, domain.ErrDataNotFound) {
			response.ErrorResponse(c, http.StatusNotFound, response.DescDataNotFound)
			return
		}

		if errors.Is(err, domain.ErrInvalidObjectID) {
			response.ErrorResponse(c, http.StatusBadRequest, response.DescInvalidObjectID)
			return
		}

		response.ErrorResponse(c, http.StatusInternalServerError, response.DescInternalServerError)
		return
	}

	retrievedWord := mapper.DomainWordToRetrievedWord(word)
	response.SuccessResponse(c, http.StatusOK, retrievedWord)
}
