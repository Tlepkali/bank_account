package models

import "errors"

var (
	ErrNotFound         = errors.New("account not found")
	ErrDuplicateAccount = errors.New("account already exists")
	ErrTimeout          = errors.New("operation timeout")
)
