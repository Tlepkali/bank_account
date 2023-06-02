package handler

import (
	"net/http"

	"bank_account/internal/models"
	"bank_account/pkg/validator"

	"github.com/go-chi/chi"
)

// @Summary		Create a new account
// @Tags			account
// @Description	Create a new account with the input payload
// @Accept			json
// @Produce		json
// @Param			account	body		models.CreateAccountDTO	true	"Account"
// @Success		201		{object}	MessageResponse
// @Failure		400		{object}	ErrorResponse
// @Failure		422		{object}	ErrorResponse
// @Router			/account [post]
func (h *Handler) CreateAccount(w http.ResponseWriter, r *http.Request) {
	var account *models.CreateAccountDTO

	if err := h.readJSON(w, r, &account); err != nil {
		h.error(w, r, http.StatusBadRequest, ErrorResponse{Error: err.Error()})
		return
	}

	v := validator.New()

	if err := account.Validate(v); err != nil {
		h.respond(w, r, http.StatusUnprocessableEntity, v.Errors)
		return
	}

	accountNumber, err := h.service.AccountService.CreateAccount(account)
	if err != nil {
		switch err {
		case models.ErrDuplicateAccount:
			h.error(w, r, http.StatusConflict, ErrorResponse{Error: err.Error()})
			return
		default:
			h.error(w, r, http.StatusInternalServerError, ErrorResponse{Error: err.Error()})
			return
		}
	}

	headers := make(http.Header)
	headers.Set("Location", "/accounts/"+accountNumber)

	message := MessageResponse{Message: "Account created successfully, account number: " + accountNumber}

	h.respond(w, r, http.StatusCreated, message, headers)
}

// @Summary		Get an account by account number
// @Tags			account
// @Description	Get an account by account number
// @Produce		json
// @Success		200	{object}	models.Account
// @Failure		400	{object}	ErrorResponse
// @Failure		404	{object}	ErrorResponse
// @Router			/account/{accountNumber} [get]
func (h *Handler) GetAccountByNumber(w http.ResponseWriter, r *http.Request) {
	paramNumber := chi.URLParam(r, "accountNumber")

	account, err := h.service.AccountService.GetAccountByNumber(paramNumber)
	if err != nil {
		switch err {
		case models.ErrNotFound:
			h.error(w, r, http.StatusNotFound, ErrorResponse{Error: err.Error()})
			return
		default:
			h.error(w, r, http.StatusInternalServerError, ErrorResponse{Error: err.Error()})
			return
		}
	}

	h.respond(w, r, http.StatusOK, account)
}

// @Summary		Get all accounts
// @Tags			accounts
// @Description	Get all accounts
// @Produce		json
// @Success		200	{object}	[]models.Account
// @Failure		400	{object}	ErrorResponse
// @Failure		404	{object}	ErrorResponse
// @Router			/accounts [get]
func (h *Handler) GetAllAccounts(w http.ResponseWriter, r *http.Request) {
	accounts, err := h.service.AccountService.GetAllAccounts()
	if err != nil {
		switch err {
		case models.ErrNotFound:
			h.error(w, r, http.StatusNotFound, ErrorResponse{Error: err.Error()})
			return
		default:
			h.error(w, r, http.StatusInternalServerError, ErrorResponse{Error: err.Error()})
			return
		}
	}

	h.respond(w, r, http.StatusOK, accounts)
}

// @Summary		Update an account by account number
// @Tags			account
// @Description	Update an account by account number
// @Accept			json
// @Produce		json
// @Param			input	body		models.CreateAccountDTO	true	"Account"
// @Success		200		{object}	models.Account
// @Failure		400		{object}	ErrorResponse
// @Failure		404		{object}	ErrorResponse
// @Router			/account/{accountNumber} [put]
func (h *Handler) UpdateAccount(w http.ResponseWriter, r *http.Request) {
	paramNumber := chi.URLParam(r, "accountNumber")

	var input *models.CreateAccountDTO

	if err := h.readJSON(w, r, &input); err != nil {
		h.error(w, r, http.StatusBadRequest, ErrorResponse{Error: err.Error()})
		return
	}

	v := validator.New()

	if err := input.Validate(v); err != nil {
		h.respond(w, r, http.StatusUnprocessableEntity, v.Errors)
		return
	}

	account, err := h.service.AccountService.GetAccountByNumber(paramNumber)
	if err != nil {
		h.error(w, r, http.StatusNotFound, ErrorResponse{Error: err.Error()})
		return
	}

	account.Owner = input.Owner
	account.Balance = input.Balance

	err = h.service.AccountService.UpdateAccount(account)
	if err != nil {
		switch err {
		case models.ErrNotFound:
			h.error(w, r, http.StatusNotFound, ErrorResponse{Error: err.Error()})
			return
		default:
			h.error(w, r, http.StatusInternalServerError, ErrorResponse{Error: err.Error()})
			return
		}
	}

	h.respond(w, r, http.StatusOK, account)
}

// @Summary		Delete an account by account number
// @Tags			account
// @Description	Delete an account by account number
// @Produce		json
// @Param			accountNumber	path		string	true	"Account Number"
// @Success		200				{object}	MessageResponse
// @Failure		400				{object}	ErrorResponse
// @Failure		404				{object}	ErrorResponse
// @Router			/account/{accountNumber} [delete]
func (h *Handler) DeleteAccount(w http.ResponseWriter, r *http.Request) {
	paramNumber := chi.URLParam(r, "accountNumber")

	err := h.service.AccountService.DeleteAccount(paramNumber)
	if err != nil {
		switch err {
		case models.ErrNotFound:
			h.error(w, r, http.StatusNotFound, ErrorResponse{Error: err.Error()})
			return
		default:
			h.error(w, r, http.StatusInternalServerError, ErrorResponse{Error: err.Error()})
			return
		}
	}

	message := MessageResponse{Message: "Account deleted successfully"}

	h.respond(w, r, http.StatusOK, message)
}
