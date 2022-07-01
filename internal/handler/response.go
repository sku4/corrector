package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/sku4/corrector/pkg/log"
)

type errorResponse struct {
	Message string `json:"message"`
}

func newErrorResponse(ctx *gin.Context, statusCode int, message string) {
	logger := log.LoggerFromGinContext(ctx)
	logger.Error(message)
	ctx.AbortWithStatusJSON(statusCode, errorResponse{message})
}
