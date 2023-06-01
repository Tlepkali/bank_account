package repository

import (
	"context"
	"database/sql"
	"errors"
	"strings"
	"time"

	"bank_account/internal/models"
)

const requestTimeout = 3 * time.Second

type AccountRepo struct {
	db *sql.DB
}

func NewAccountRepo(db *sql.DB) *AccountRepo {
	return &AccountRepo{db: db}
}

func (r *AccountRepo) CreateAccount(account *models.Account) (string, error) {
	query := `INSERT INTO accounts (account_number, owner, balance)
	VALUES ($1, $2, $3) 
	RETURNING account_number, created_at, updated_at`

	args := []interface{}{
		account.AccountNumber,
		account.Owner,
		account.Balance,
	}

	ctx, cancel := context.WithTimeout(context.Background(), requestTimeout)
	defer cancel()

	err := r.db.QueryRowContext(ctx, query, args...).Scan(
		&account.AccountNumber,
		&account.CreatedAt,
		&account.UpdatedAt,
	)

	switch {
	case errors.Is(err, context.DeadlineExceeded):
		return "", models.ErrTimeout
	case err != nil && strings.Contains(err.Error(), "pq: duplicate key value violates unique constraint"):
		return "", models.ErrDuplicateAccount
	case err != nil:
		return "", err
	default:
		return account.AccountNumber, nil
	}
}

// func (r *AccountRepo) GetAccountByID(id int64) (*models.Account, error) {
// 	query := `SELECT id, account_number, owner, balance, created_at, updated_at
// 	FROM accounts
// 	WHERE id = $1`

// 	var account models.Account

// 	ctx, cancel := context.WithTimeout(context.Background(), requestTimeout)
// 	defer cancel()

// 	err := r.db.QueryRowContext(ctx, query, id).Scan(
// 		&account.ID,
// 		&account.AccountNumber,
// 		&account.Owner,
// 		&account.Balance,
// 		&account.CreatedAt,
// 		&account.UpdatedAt,
// 	)

// 	switch {
// 	case errors.Is(err, sql.ErrNoRows):
// 		return nil, models.ErrNotFound
// 	case errors.Is(err, context.DeadlineExceeded):
// 		return nil, models.ErrTimeout
// 	case err != nil:
// 		return nil, err
// 	default:
// 		return &account, nil
// 	}
// }

func (r *AccountRepo) GetAccountByNumber(accountNumber string) (*models.Account, error) {
	query := `SELECT id, account_number, owner, balance, created_at, updated_at
	FROM accounts 
	WHERE account_number = $1`

	var account models.Account

	ctx, cancel := context.WithTimeout(context.Background(), requestTimeout)
	defer cancel()

	err := r.db.QueryRowContext(ctx, query, accountNumber).Scan(
		&account.ID,
		&account.AccountNumber,
		&account.Owner,
		&account.Balance,
		&account.CreatedAt,
		&account.UpdatedAt,
	)

	switch {
	case errors.Is(err, sql.ErrNoRows):
		return nil, models.ErrNotFound
	case errors.Is(err, context.DeadlineExceeded):
		return nil, models.ErrTimeout
	case err != nil:
		return nil, err
	default:
		return &account, nil
	}
}

func (r *AccountRepo) GetAllAccounts() ([]*models.Account, error) {
	query := `SELECT id, account_number, owner, balance, created_at, updated_at
	FROM accounts`

	ctx, cancel := context.WithTimeout(context.Background(), requestTimeout)
	defer cancel()

	rows, err := r.db.QueryContext(ctx, query)
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

	if err := rows.Err(); err != nil {
		return nil, err
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

	ctx, cancel := context.WithTimeout(context.Background(), requestTimeout)
	defer cancel()

	err := r.db.QueryRowContext(ctx, query, args...).Scan(&account.UpdatedAt)

	switch {
	case errors.Is(err, sql.ErrNoRows):
		return models.ErrNotFound
	case errors.Is(err, context.DeadlineExceeded):
		return models.ErrTimeout
	case err != nil:
		return err
	default:
		return nil
	}
}

func (r *AccountRepo) DeleteAccount(accountNumber string) error {
	query := `DELETE FROM accounts WHERE account_number = $1`

	ctx, cancel := context.WithTimeout(context.Background(), requestTimeout)
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
		return models.ErrNotFound
	}

	return nil
}
