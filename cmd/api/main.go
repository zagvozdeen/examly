package main

import (
	"context"
	"github.com/den4ik117/examly/api"
	"github.com/den4ik117/examly/internal/db"
	"github.com/den4ik117/examly/internal/env"
	"github.com/den4ik117/examly/internal/store"
	"github.com/joho/godotenv"
	"github.com/rs/zerolog"
	"os"
)

func main() {
	logger := zerolog.New(zerolog.ConsoleWriter{
		Out:        os.Stderr,
		TimeFormat: "02.01.2006 15:04:05",
	}).Level(zerolog.TraceLevel).With().Caller().Timestamp().Logger()

	err := godotenv.Load()
	if err != nil {
		logger.Panic().Err(err).Msg("Failed to load .env file")
	}

	cfg := api.Config{
		AppEnv:    env.GetString("APP_ENV", "development"),
		AppURL:    env.GetString("APP_URL", "127.0.0.1:8080"),
		DBAddr:    env.GetString("DB_ADDR", "postgres://root:root@127.0.0.1:5432/examly?sslmode=disable"),
		SecretKey: env.GetString("APP_KEY", ""),
	}

	conn, err := db.New(context.Background(), cfg.DBAddr)
	if err != nil {
		logger.Panic().Err(err).Msg("Failed to connect to PostgreSQL")
	}
	logger.Info().Msg("Successfully connected to PostgreSQL")

	storage := store.NewStorage(conn)

	app := api.NewApplication(logger, cfg, storage)

	r := app.Mount()

	logger.Fatal().Err(app.Run(r))
}
