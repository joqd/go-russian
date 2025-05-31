package response

import (
	"github.com/gin-gonic/gin"
)

type ErrorResponseWrapper struct {
	Ok          bool   `json:"ok" example:"false"`
	ErrorCode   int    `json:"error_code" example:"404"`
	Description string `json:"description" example:"data not found"`
}

func ErrorResponse(c *gin.Context, status int, description string) {
	c.JSON(status, ErrorResponseWrapper{
		Ok:          false,
		ErrorCode:   status,
		Description: description,
	})
}

type SuccessResponseWrapper[T any] struct {
	Ok     bool `json:"ok"`
	Result *T   `json:"result"`
}

func SuccessResponse[T any](c *gin.Context, status int, result T) {
	c.JSON(status, SuccessResponseWrapper[T]{
		Ok:     true,
		Result: &result,
	})
}
