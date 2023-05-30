package repository

import (
	"database/sql"

	"bank_account/internal/models"
)

type AccountRepo struct {
	db *sql.DB
}

func NewAccountRepo(db *sql.DB) *AccountRepo {
	return &AccountRepo{db: db}
}

func (r *AccountRepo) GetAccount(id int) (*models.Account, error) {
	var account models.Account
	err := r.db.QueryRow("SELECT * FROM accounts WHERE id = $1", id).Scan(&account.ID, &account.Name, &account.Balance)
	if err != nil {
		return nil, err
	}
	return &account, nil
}

func (r *AccountRepo) UpdateAccount(account *models.Account) error {
	_, err := r.db.Exec("UPDATE accounts SET name = $1, balance = $2 WHERE id = $3", account.Name, account.Balance, account.ID)
	if err != nil {
		return err
	}
	return nil
}

func (r *AccountRepo) CreateAccount(account *models.Account) error {
	_, err := r.db.Exec("INSERT INTO accounts (name, balance) VALUES ($1, $2)", account.Name, account.Balance)
	if err != nil {
		return err
	}
	return nil
}

func (r *AccountRepo) DeleteAccount(id int) error {
	_, err := r.db.Exec("DELETE FROM accounts WHERE id = $1", id)
	if err != nil {
		return err
	}
	return nil
}
