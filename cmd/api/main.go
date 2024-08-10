package main

import (
	"context"
	"github.com/Den4ik117/examly/config"
	"github.com/Den4ik117/examly/internal/handler"
	"github.com/Den4ik117/examly/internal/repository"
	"github.com/Den4ik117/examly/internal/service"
	"log"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	go initLogger(ctx)

	db, err := repository.NewPostgresDB()
	if err != nil {
		log.Fatalf("Failed to connect to db: %s", err)
	}

	repos := repository.NewRepository(db)
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)

	router := handlers.InitRoutes()

	go func() {
		slog.Info("Starting server on :8080")
		if err = http.ListenAndServe(":8080", router); err != nil {
			log.Fatalf("Failed to start server: %s", err)
		}
	}()

	<-ctx.Done()
}

func initLogger(ctx context.Context) {
	logger := slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{
		AddSource: true,
		Level:     slog.LevelDebug,
	}))

	if config.Envs.AppEnv == "production" {
		f, err := os.OpenFile(
			time.Now().Format("logs/examly-2006-01-02.log"),
			os.O_APPEND|os.O_RDWR|os.O_CREATE,
			0666,
		)
		if err != nil {
			log.Fatalf("Failed to open log file: %s", err)
		}
		defer f.Close()
		logger = slog.New(slog.NewTextHandler(f, &slog.HandlerOptions{
			AddSource: true,
		}))
	}

	slog.SetDefault(logger)

	<-ctx.Done()

	slog.Info("Shutting down server")
}
