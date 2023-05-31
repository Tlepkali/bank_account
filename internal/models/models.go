package models

import "time"

type Account struct {
	ID        string    `json:"id"`
	Owner     string    `json:"name"`
	Balance   float64   `json:"balance"`
	CreatedAt time.Time `json:"created_at"`
	Version   int       `json:"version"`
}

type AccountService interface {
	CreateAccount(account *Account) error
	GetAccount(id string) (*Account, error)
	UpdateAccount(account *Account) error
	DeleteAccount(id string) error
}

type AccountRepository interface {
	CreateAccount(account *Account) error
	GetAccount(id string) (*Account, error)
	UpdateAccount(account *Account) error
	DeleteAccount(id string) error
}
