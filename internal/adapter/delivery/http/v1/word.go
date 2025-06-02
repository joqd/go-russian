package v1

import (
	"errors"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/joqd/go-russian/internal/adapter/delivery/http/mapper"
	"github.com/joqd/go-russian/internal/adapter/delivery/http/request"
	"github.com/joqd/go-russian/internal/adapter/delivery/http/response"
	_ "github.com/joqd/go-russian/internal/adapter/delivery/http/response/wrapper"
	"github.com/joqd/go-russian/internal/core/domain"
	"github.com/joqd/go-russian/internal/core/port"
	"go.mongodb.org/mongo-driver/v2/bson"

	"github.com/gin-gonic/gin"
)

type wordHandler struct {
	usecase   port.WordUsecase
	xlog      port.Logger
	validator *validator.Validate
}

func NewWordHandler(usecase port.WordUsecase, xlog port.Logger) *wordHandler {
	return &wordHandler{
		usecase:   usecase,
		xlog:      xlog,
		validator: validator.New(validator.WithRequiredStructEnabled()),
	}
}

func (w *wordHandler) Get(c *gin.Context) {
	query := c.Param("query")

	if _, err := bson.ObjectIDFromHex(query); err == nil {
		w.GetByID(c)
	} else {
		w.GetByBare(c)
	}
}

// @Summary      Get a word by ID
// @Description  Retrieve a word from the database using its ID
// @Tags         words
// @Param        id   path      string  true  "Word ID"
// @Success      200  {object}  wrapper.RetrievedWordWrapper
// @Failure      404  {object}  wrapper.ErrorNotFoundWrapper
// @Failure      400  {object}  wrapper.ErrorInvalidObjectIdWrapper
// @Failure      500  {object}  wrapper.ErrorInternalServerWrapper
// @Router       /api/v1/words/{id} [get]
func (w *wordHandler) GetByID(c *gin.Context) {
	id := c.Param("query")

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

	retrievedWord := mapper.WordToRetrievedWord(word)
	response.SuccessResponse(c, http.StatusOK, retrievedWord)
}

// @Summary      Create a word
// @Description  Create a word with payload
// @Tags         words
// @Accept       json
// @Produce      json
// @Param        request  body  request.CreateWord  true  "Word formation data"
// @Success      201  {object}  wrapper.RetrievedWordWrapper
// @Failure      400  {object}  wrapper.ErrorBadRequestWrapper
// @Failure      422  {object}  wrapper.ErrorUnprocessableEntityWrapper
// @Failure      500  {object}  wrapper.ErrorInternalServerWrapper
// @Router       /api/v1/words  [post]
func (w *wordHandler) Create(c *gin.Context) {
	var body request.CreateWord
	if err := c.ShouldBind(&body); err != nil {
		response.ErrorResponse(c, http.StatusBadRequest, response.DescBadRequest)
		return
	}

	if err := w.validator.Struct(body); err != nil {
		response.ErrorResponse(c, http.StatusUnprocessableEntity, response.UnprocessableEntity)
		return
	}

	payload := mapper.CreateWordToWord(&body)

	word, err := w.usecase.Create(c.Request.Context(), payload)
	if err != nil {
		response.ErrorResponse(c, http.StatusInternalServerError, response.DescInternalServerError)
		return
	}

	retrievedWord := mapper.WordToRetrievedWord(word)
	response.SuccessResponse(c, http.StatusOK, retrievedWord)
}

// @Summary      Get a word by Bare
// @Description  Retrieve a word from the database using its Bare (raw word)
// @Tags         words
// @Param        bare path      string  true  "Raw Word"
// @Success      200  {object}  wrapper.RetrievedWordWrapper
// @Failure      404  {object}  wrapper.ErrorNotFoundWrapper
// @Failure      500  {object}  wrapper.ErrorInternalServerWrapper
// @Router       /api/v1/words/{bare} [get]
func (w *wordHandler) GetByBare(c *gin.Context) {
	bare := c.Param("query")

	word, err := w.usecase.GetByBare(c.Request.Context(), bare)
	if err != nil {
		if errors.Is(err, domain.ErrDataNotFound) {
			response.ErrorResponse(c, http.StatusNotFound, response.DescDataNotFound)
			return
		}

		response.ErrorResponse(c, http.StatusInternalServerError, response.DescInternalServerError)
		return
	}

	retrievedWord := mapper.WordToRetrievedWord(word)
	response.SuccessResponse(c, http.StatusOK, retrievedWord)
}
