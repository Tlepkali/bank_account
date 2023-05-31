package service

import (
	"bank_account/internal/models"
	"bank_account/internal/repository"
)

type Service struct {
	AccountService models.AccountService
}

func NewService(repo *repository.Repository) *Service {
	return &Service{
		AccountService: NewAccountService(repo.AccountRepo),
	}
}
