package service

import (
	"bank_account/internal/models"
)

type AccountService struct {
	repo models.AccountRepository
}

func NewAccountService(repo models.AccountRepository) *AccountService {
	return &AccountService{repo: repo}
}

func (s *AccountService) CreateAccount(dto *models.CreateAccountDTO) (string, error) {
	account := &models.Account{
		Owner:   dto.Owner,
		Balance: dto.Balance,
	}
	account.GenerateAccountNumber()

	return s.repo.CreateAccount(account)
}

func (s *AccountService) GetAccountByNumber(accountNumber string) (*models.Account, error) {
	return s.repo.GetAccountByNumber(accountNumber)
}

func (s *AccountService) GetAllAccounts() ([]*models.Account, error) {
	return s.repo.GetAllAccounts()
}

func (s *AccountService) UpdateAccount(account *models.Account) error {
	return s.repo.UpdateAccount(account)
}

func (s *AccountService) DeleteAccount(id string) error {
	return s.repo.DeleteAccount(id)
}
