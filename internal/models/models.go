package models

import "time"

type Account struct {
	ID            int64     `json:"-" db:"id"`
	AccountNumber string    `json:"account_number" db:"account_number"`
	Owner         string    `json:"owner" db:"owner"`
	Balance       float64   `json:"balance" db:"balance"`
	CreatedAt     time.Time `json:"created_at" db:"created_at"`
	UpdatedAt     time.Time `json:"updated_at" db:"updated_at"`
}

type AccountService interface {
	CreateAccount(account *Account) error
	GetAccountByID(id int64) (*Account, error)
	GetAccountByNumber(accountNumber string) (*Account, error)
	GetAllAccounts(limit, offset int) ([]*Account, error)
	UpdateAccount(account *Account) error
	DeleteAccount(accountNumber string) error
}

type AccountRepository interface {
	CreateAccount(account *Account) error
	GetAccountByID(id int64) (*Account, error)
	GetAccountByNumber(accountNumber string) (*Account, error)
	GetAllAccounts(limit, offset int) ([]*Account, error)
	UpdateAccount(account *Account) error
	DeleteAccount(accountNumber string) error
}
