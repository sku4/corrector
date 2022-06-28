package handler

import (
	"context"
	"github.com/gin-gonic/gin"
	_ "github.com/sku4/corrector/docs"
	"github.com/sku4/corrector/internal/service"
	"github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
)

type Handler struct {
	ctx      context.Context
	services service.Service
}

func NewHandler(ctx context.Context, services *service.Service) *Handler {
	return &Handler{
		ctx:      ctx,
		services: *services,
	}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	list := router.Group("/corrector")
	{
		list.POST("/", h.correctorRequest)
	}

	return router
}
