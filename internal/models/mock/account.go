package mock

import (
	"bank_account/internal/models"
)

var mockAccount = &models.Account{
	Owner:         "test",
	AccountNumber: "1234567890",
	Balance:       100,
}

type AccountModel struct{}

func (m *AccountModel) CreateAccount(account *models.Account) (string, error) {
	return "1234567890", nil
}

func (m *AccountModel) GetAccountByNumber(accountNumber string) (*models.Account, error) {
	switch accountNumber {
	case "1234567890":
		return mockAccount, nil
	default:
		return nil, models.ErrNotFound
	}
}

func (m *AccountModel) GetAllAccounts() ([]*models.Account, error) {
	return []*models.Account{mockAccount}, nil
}

func (m *AccountModel) UpdateAccount(account *models.Account) error {
	return nil
}

func (m *AccountModel) DeleteAccount(accountNumber string) error {
	switch accountNumber {
	case "1234567890":
		return nil
	default:
		return models.ErrNotFound
	}
}
