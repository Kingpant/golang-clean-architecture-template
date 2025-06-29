package main

import (
	"os"
	"os/signal"
	"syscall"

	"github.com/Kingpant/golang-clean-architecture-template/internal/infrastructure/config"
	"github.com/Kingpant/golang-clean-architecture-template/internal/infrastructure/db"
	"github.com/Kingpant/golang-clean-architecture-template/internal/infrastructure/logger"
	"github.com/Kingpant/golang-clean-architecture-template/internal/infrastructure/repository"
	"github.com/Kingpant/golang-clean-architecture-template/internal/interface/handler"
	"github.com/Kingpant/golang-clean-architecture-template/internal/interface/router"
	"github.com/Kingpant/golang-clean-architecture-template/internal/usecase"
)

// @title			template API
// @version		1.0
// @description	This is the API for the template application.
// @BasePath		/
func main() {
	cfg, loadConfigErr := config.LoadAPIConfig()
	if loadConfigErr != nil {
		panic(loadConfigErr)
	}

	logger := logger.InitLogger(cfg.AppEnv)

	bunDb := db.NewDB(
		cfg.AppEnv,
		cfg.PostgresqlUsername,
		cfg.PostgresqlPassword,
		cfg.PostgresqlHost,
		cfg.PostgresqlDatabase,
		cfg.PostgresqlSchema,
		cfg.PostgresqlSSL,
	)

	// Initialize Repositories
	userRepo := repository.NewUserPGRepository(bunDb)

	// Initialize Usecases
	userUseCase := usecase.NewUserUseCase(userRepo, logger)

	// Initialize Handlers
	userHandler := handler.NewUserHandler(userUseCase)

	f := router.NewFiberRouter(cfg.AppEnv, router.WithPinger(bunDb))
	router.RegisterUserRouter(f.App(), userHandler)

	go func() {
		quit := make(chan os.Signal, 1)
		signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
		logger.Info("Received shutdown signal, shutting down gracefully")

		<-quit
		if err := f.Shutdown(); err != nil {
			panic(err)
		}
		os.Exit(0)
	}()

	logger.Infow("Starting API server", "port", cfg.AppPort)
	if err := f.Listen(cfg.AppPort); err != nil {
		panic(err)
	}
}
