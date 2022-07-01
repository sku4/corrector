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

	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt)
	ctx = log.ContextWithLogger(ctx, logger)
	defer stop()

	repos := repository.NewRepository()
	services := service.NewService(repos)
	handlers := handler.NewHandler(ctx, services)

	quit := make(chan os.Signal, 1)
	srv := new(app.Server)
	go func() {
		if err := srv.Run(cfg.Port, handlers.InitRoutes()); err != nil {
			sugar.Fatalf("error occured while running http server: %s", err.Error())
			quit <- nil
		}
	}()

	sugar.Info("App Started")
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	<-quit

	sugar.Info("App Shutting Down")

	if err := srv.Shutdown(context.Background()); err != nil {
		sugar.Errorf("error occured on server shutting down: %s", err.Error())
	}
}
