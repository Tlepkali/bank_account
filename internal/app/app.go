package app

import (
	"bank_account/config"
	"bank_account/internal/handler"
	"bank_account/internal/repository"
	"bank_account/internal/service"
	"bank_account/pkg/client/postgres"
)

func Run() {
	config := config.GetConfig()

	db, err := postgres.OpenDB(config.DSN)
	if err != nil {
		panic(err)
	}

	repo := repository.NewRepository(db)

	service := service.NewService(repo)

	handler := handler.NewHandler(service)

	handler.InitRoutes()
}
