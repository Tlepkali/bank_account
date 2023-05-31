package app

import (
	"log"
	"net/http"

	"bank_account/config"
	"bank_account/internal/handler"
	"bank_account/internal/repository"
	"bank_account/internal/service"
	"bank_account/pkg/client/postgres"
)

func Run() {
	config := config.GetConfig()
	log.Println("Configurated")

	db, err := postgres.OpenDB(config.DSN)
	if err != nil {
		panic(err)
	}
	log.Println("Connected to DB")

	repo := repository.NewRepository(db)

	service := service.NewService(repo)

	handler := handler.NewHandler(service)

	log.Println("Server is starting on port", config.Port)
	log.Fatal(http.ListenAndServe(config.Port, handler.InitRoutes()))
}
