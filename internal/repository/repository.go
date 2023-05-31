package repository

import (
	"database/sql"

	"bank_account/internal/models"
)

type Repository struct {
	AccountRepo models.AccountRepository
}

func NewRepository(db *sql.DB) *Repository {
	return &Repository{
		AccountRepo: NewAccountRepo(db),
	}
}
