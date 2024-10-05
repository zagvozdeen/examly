package main

import (
	"context"
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
	}).With().Timestamp().Logger()

	err := godotenv.Load()
	if err != nil {
		logger.Fatal().Err(err)
	}

	cfg := config{
		AppEnv:    env.GetString("APP_ENV", "development"),
		AppURL:    env.GetString("APP_URL", "127.0.0.1:8080"),
		DBAddr:    env.GetString("DB_ADDR", "postgres://root:root@127.0.0.1:5432/examly?sslmode=disable"),
		SecretKey: env.GetString("APP_KEY", ""),
	}

	conn, err := db.New(context.Background(), cfg.DBAddr)
	if err != nil {
		logger.Fatal().Err(err)
	}
	logger.Info().Msg("Successfully connected to PostgreSQL")

	storage := store.NewStorage(conn)

	initValidator()

	app := &application{
		log:    logger,
		config: cfg,
		store:  storage,
	}

	r := app.mount()

	logger.Fatal().Err(app.run(r))
}
