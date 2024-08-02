package main

import (
	"github.com/Den4ik117/examly/internal/handler"
	"github.com/Den4ik117/examly/internal/repository"
	"github.com/Den4ik117/examly/internal/service"
	"log"
	"net/http"
)

func main() {
	db, err := repository.NewPostgresDB()
	if err != nil {
		log.Fatalf("Failed to connect to db: %s", err)
	}

	repos := repository.NewRepository(db)
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)

	router := handlers.InitRoutes()

	err = http.ListenAndServe(":8080", router)
	if err != nil {
		log.Fatalf("Failed to start server: %s", err)
	}
}
