package handler

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/magiconair/properties/assert"
	"github.com/sku4/corrector/pkg/log"
	"go.uber.org/zap"
	"net/http/httptest"
	"testing"
)

func Test_newErrorResponse(t *testing.T) {
	ctx := context.Background()
	ginCtx, _ := gin.CreateTestContext(httptest.NewRecorder())
	tests := []struct {
		name           string
		ctx            *gin.Context
		logger         *zap.Logger
		statusCode     int
		message        string
		wantStatusCode int
	}{
		{
			name:           "error 500",
			ctx:            ginCtx,
			logger:         log.LoggerFromContext(ctx),
			statusCode:     500,
			message:        "error 500",
			wantStatusCode: 500,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			newErrorResponse(tt.ctx, tt.logger, tt.statusCode, tt.message)
			assert.Equal(t, tt.ctx.Writer.Status(), tt.wantStatusCode)
		})
	}
}
