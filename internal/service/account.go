package service

import "bank_account/internal/models"

type AccountService struct {
	repo models.AccountRepository
}

func NewAccountService(repo models.AccountRepository) *AccountService {
	return &AccountService{repo: repo}
}

func (s *AccountService) CreateAccount(account *models.Account) error {
	return s.repo.CreateAccount(account)
}

func (s *AccountService) GetAccount(id string) (*models.Account, error) {
	return s.repo.GetAccount(id)
}

func (s *AccountService) UpdateAccount(account *models.Account) error {
	return s.repo.UpdateAccount(account)
}

func (s *AccountService) DeleteAccount(id string) error {
	return s.repo.DeleteAccount(id)
}
