package app

import (
	"log"

	"bank_account/config"
	"bank_account/internal/handler"
	"bank_account/internal/repository"
	"bank_account/internal/service"
	"bank_account/pkg/client/postgres"
)

func Run() {
	config := config.GetConfig()
	log.Println("Configs are parsed")

	db, err := postgres.OpenDB(config.DSN)
	if err != nil {
		log.Println("Error while connecting to DB:", err)
		return
	}
	defer db.Close()
	log.Println("Connected to DB")

	repo := repository.NewRepository(db)

	service := service.NewService(repo)

	handler := handler.NewHandler(service)

	log.Println("Server is starting on port", config.Port)

	if err := Serve(config, handler.InitRoutes()); err != nil {
		log.Println("Error while running http server:", err)
	}
}
