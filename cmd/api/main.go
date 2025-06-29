package main

import (
	"os"
	"os/signal"
	"syscall"

	"github.com/Kingpant/tipster/internal/infrastructure/config"
	"github.com/Kingpant/tipster/internal/infrastructure/db"
	"github.com/Kingpant/tipster/internal/infrastructure/logger"
	"github.com/Kingpant/tipster/internal/infrastructure/repository"
	"github.com/Kingpant/tipster/internal/interface/handler"
	"github.com/Kingpant/tipster/internal/interface/router"
	"github.com/Kingpant/tipster/internal/usecase"
)

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

	f := router.NewFiberRouter(router.WithPinger(bunDb))
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
