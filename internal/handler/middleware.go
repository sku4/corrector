package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/sku4/corrector/pkg/log"
)

func (h *Handler) logger(c *gin.Context) {
	l := log.LoggerFromContext(h.ctx)
	c.Set(log.LoggerKey, l)
}
