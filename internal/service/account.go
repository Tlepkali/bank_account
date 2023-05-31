package service

import "bank_account/internal/models"

type AccountService struct {
	repo models.AccountRepository
}

func NewAccountService(repo models.AccountRepository) *AccountService {
	return &AccountService{repo: repo}
}

func (s *AccountService) CreateAccount(account *models.Account) error {
	account.GenerateAccountNumber()
	return s.repo.CreateAccount(account)
}

func (s *AccountService) GetAccountByID(id int64) (*models.Account, error) {
	return s.repo.GetAccountByID(id)
}

func (s *AccountService) GetAccountByNumber(accountNumber string) (*models.Account, error) {
	return s.repo.GetAccountByNumber(accountNumber)
}

func (s *AccountService) GetAllAccounts(limit, offset int) ([]*models.Account, error) {
	return s.repo.GetAllAccounts(limit, offset)
}

func (s *AccountService) UpdateAccount(account *models.Account) error {
	return s.repo.UpdateAccount(account)
}

func (s *AccountService) DeleteAccount(id string) error {
	return s.repo.DeleteAccount(id)
}
