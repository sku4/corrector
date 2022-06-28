package handler

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type errorResponse struct {
	Message string `json:"message"`
}

func newErrorResponse(ctx *gin.Context, logger *zap.Logger, statusCode int, message string) {
	logger.Error(message)
	ctx.AbortWithStatusJSON(statusCode, errorResponse{message})
}
