package repository

import (
	"context"
	"database/sql"
	"errors"
	"time"

	"bank_account/internal/models"
)

type AccountRepo struct {
	db *sql.DB
}

func NewAccountRepo(db *sql.DB) *AccountRepo {
	return &AccountRepo{db: db}
}

var ErrNotFound = errors.New("account not found")

func (r *AccountRepo) CreateAccount(account *models.Account) error {
	query := `INSERT INTO accounts (account_number, owner, balance)
	VALUES ($1, $2, $3) 
	RETURNING account_number, created_at, updated_at`

	args := []interface{}{
		account.AccountNumber,
		account.Owner,
		account.Balance,
	}

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	return r.db.QueryRowContext(ctx, query, args...).Scan(
		&account.AccountNumber,
		&account.CreatedAt,
		&account.UpdatedAt,
	)
}

func (r *AccountRepo) GetAccountByID(id int64) (*models.Account, error) {
	query := `SELECT id, account_number, owner, balance, created_at, updated_at
	FROM accounts 
	WHERE id = $1`

	var account models.Account

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	err := r.db.QueryRowContext(ctx, query, id).Scan(
		&account.ID,
		&account.AccountNumber,
		&account.Owner,
		&account.Balance,
		&account.CreatedAt,
		&account.UpdatedAt,
	)
	if err != nil {
		return nil, err
	}

	return &account, nil
}

func (r *AccountRepo) GetAccountByNumber(accountNumber string) (*models.Account, error) {
	query := `SELECT id, account_number, owner, balance, created_at, updated_at
	FROM accounts 
	WHERE account_number = $1`

	var account models.Account

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	err := r.db.QueryRowContext(ctx, query, accountNumber).Scan(
		&account.ID,
		&account.AccountNumber,
		&account.Owner,
		&account.Balance,
		&account.CreatedAt,
		&account.UpdatedAt,
	)
	if err != nil {
		return nil, err
	}

	return &account, nil
}

func (r *AccountRepo) GetAllAccounts(limit, offset int) ([]*models.Account, error) {
	query := `SELECT id, account_number, owner, balance, created_at, updated_at
	FROM accounts 
	LIMIT $1 OFFSET $2`

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	rows, err := r.db.QueryContext(ctx, query, limit, offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var accounts []*models.Account

	for rows.Next() {
		var account models.Account

		err := rows.Scan(
			&account.ID,
			&account.AccountNumber,
			&account.Owner,
			&account.Balance,
			&account.CreatedAt,
			&account.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}

		accounts = append(accounts, &account)
	}

	return accounts, nil
}

func (r *AccountRepo) UpdateAccount(account *models.Account) error {
	query := `UPDATE accounts 
	SET owner = $1, balance = $2, updated_at = NOW()
	WHERE account_number = $3 AND updated_at = $4
	RETURNING updated_at`

	args := []interface{}{
		account.Owner,
		account.Balance,
		account.AccountNumber,
		account.UpdatedAt,
	}

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	err := r.db.QueryRowContext(ctx, query, args...).Scan(&account.UpdatedAt)
	if err != nil {
		return err
	}

	return nil
}

func (r *AccountRepo) DeleteAccount(accountNumber string) error {
	query := `DELETE FROM accounts WHERE account_number = $1`

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	res, err := r.db.ExecContext(ctx, query, accountNumber)
	if err != nil {
		return err
	}

	rowsAffected, err := res.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected != 1 {
		return ErrNotFound
	}

	return nil
}
