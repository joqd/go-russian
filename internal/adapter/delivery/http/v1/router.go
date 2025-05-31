package v1

import (
	"github.com/joqd/ruskee/internal/core/port"

	"github.com/gin-gonic/gin"
)

func RegisterWordRouter(rg *gin.RouterGroup, usecase port.WordUsecase, xlog port.Logger) {
	wordHTTP := NewWordHTTP(usecase, xlog)

	wordGroup := rg.Group("/words")

	{
		wordGroup.GET("/:id", wordHTTP.GetByID)
	}
}
