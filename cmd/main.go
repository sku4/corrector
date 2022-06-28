package main

import (
	"context"
	app "github.com/sku4/corrector"
	"github.com/sku4/corrector/configs"
	"github.com/sku4/corrector/internal/handler"
	"github.com/sku4/corrector/internal/repository"
	"github.com/sku4/corrector/internal/service"
	"github.com/sku4/corrector/pkg/log"
	"go.uber.org/zap"
	"os"
	"os/signal"
	"syscall"
)

// @title Corrector App API
// @version 1.0
// @description API Server for Corrector application

// @host localhost:8000
// @BasePath /

func main() {
	logger, _ := zap.NewProduction()
	defer func(logger *zap.Logger) {
		_ = logger.Sync()
	}(logger)
	sugar := logger.Sugar()

	cfg, err := configs.Init()
	if err != nil {
		sugar.Fatalf("error init config: %s", err.Error())
	}

	ctx := log.ContextWithLogger(context.Background(), logger)

	repos := repository.NewRepository(ctx, cfg)
	services := service.NewService(ctx, repos)
	handlers := handler.NewHandler(ctx, services)

	srv := new(app.Server)
	go func() {
		if err := srv.Run(cfg.Port, handlers.InitRoutes()); err != nil {
			sugar.Fatalf("error occured while running http server: %s", err.Error())
		}
	}()

	sugar.Info("App Started")

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	<-quit

	sugar.Info("App Shutting Down")

	if err := srv.Shutdown(context.Background()); err != nil {
		sugar.Errorf("error occured on server shutting down: %s", err.Error())
	}
}
