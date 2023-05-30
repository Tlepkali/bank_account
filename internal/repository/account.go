package repository

import (
	"context"
	"database/sql"
	"time"

	"bank_account/internal/models"
)

type AccountRepo struct {
	db *sql.DB
}

func NewAccountRepo(db *sql.DB) *AccountRepo {
	return &AccountRepo{db: db}
}

func (r *AccountRepo) CreateAccount(account *models.Account) error {
	query := `INSERT INTO accounts (id, name, balance, created_at, version) VALUES ($1, $2, $3, $4, $5)`

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	_, err := r.db.ExecContext(ctx, query, account.ID, account.Owner, account.Balance, account.CreatedAt, account.Version)
	if err != nil {
		return err
	}

	return nil
}

func (r *AccountRepo) GetAccount(id string) (*models.Account, error) {
	query := `SELECT id, name, balance, created_at, version FROM accounts WHERE id = $1`

	var account models.Account

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	err := r.db.QueryRowContext(ctx, query, id).Scan(
		&account.ID,
		&account.Owner,
		&account.Balance,
		&account.CreatedAt,
		&account.Version,
	)
	if err != nil {
		return nil, err
	}

	return &account, nil
}

func (r *AccountRepo) UpdateAccount(account *models.Account) error {
	query := `UPDATE accounts SET name = $1, balance = $2, created_at = $3, version = version + 1 
	WHERE id = $5  AND version = $4
	RETURNING version`

	args := []interface{}{
		account.Owner,
		account.Balance,
		account.CreatedAt,
		account.Version,
		account.ID,
	}

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	err := r.db.QueryRowContext(ctx, query, args...).Scan(&account.Version)
	if err != nil {
		return err
	}

	return nil
}

func (r *AccountRepo) DeleteAccount(id int) error {
	query := `DELETE FROM accounts WHERE id = $1`

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	_, err := r.db.ExecContext(ctx, query, id)
	if err != nil {
		return err
	}

	return nil
}
