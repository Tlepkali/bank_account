package models

import "bank_account/pkg/validator"

func (a *Account) ValidateAccount(v *validator.Validator) map[string]string {
	if a.Owner == "" {
		v.AddError("owner", "must be provided")
	}

	if a.Balance < 0 {
		v.AddError("balance", "must be greater than zero")
	}

	if !v.Valid() {
		return v.Errors
	}

	return nil
}
