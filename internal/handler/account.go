package handler

import (
	"fmt"
	"net/http"

	"bank_account/internal/models"
	"bank_account/pkg/validator"

	"github.com/go-chi/chi"
)

func (h *Handler) CreateAccount(w http.ResponseWriter, r *http.Request) {
	var account *models.CreateAccountDTO

	if err := h.readJSON(w, r, &account); err != nil {
		h.error(w, r, http.StatusBadRequest, err)
		return
	}

	v := validator.New()

	if err := account.Validate(v); err != nil {
		h.respond(w, r, http.StatusUnprocessableEntity, v.Errors)
		return
	}

	accountNumber, err := h.service.AccountService.CreateAccount(account)
	fmt.Println(accountNumber)
	if err != nil {
		switch err {
		case models.ErrDuplicateAccount:
			h.error(w, r, http.StatusConflict, err)
			return
		default:
			h.error(w, r, http.StatusInternalServerError, err)
			return
		}
	}

	headers := make(http.Header)
	headers.Set("Location", "/accounts/"+accountNumber)

	resp := map[string]string{
		"account_number": accountNumber,
	}

	h.respond(w, r, http.StatusCreated, resp, headers)
}

func (h *Handler) GetAccountByNumber(w http.ResponseWriter, r *http.Request) {
	paramNumber := chi.URLParam(r, "accountNumber")

	account, err := h.service.AccountService.GetAccountByNumber(paramNumber)
	if err != nil {
		switch err {
		case models.ErrNotFound:
			h.error(w, r, http.StatusNotFound, err)
		default:
			h.error(w, r, http.StatusInternalServerError, err)
		}
	}

	h.respond(w, r, http.StatusOK, account)
}

func (h *Handler) GetAllAccounts(w http.ResponseWriter, r *http.Request) {
	accounts, err := h.service.AccountService.GetAllAccounts()
	if err != nil {
		h.error(w, r, http.StatusInternalServerError, err)
		return
	}

	h.respond(w, r, http.StatusOK, accounts)
}

func (h *Handler) UpdateAccount(w http.ResponseWriter, r *http.Request) {
	paramNumber := chi.URLParam(r, "accountNumber")

	account, err := h.service.AccountService.GetAccountByNumber(paramNumber)
	if err != nil {
		h.error(w, r, http.StatusNotFound, err)
		return
	}

	var input *models.CreateAccountDTO

	if err := h.readJSON(w, r, &input); err != nil {
		h.error(w, r, http.StatusBadRequest, err)
		return
	}

	v := validator.New()

	if err := input.Validate(v); err != nil {
		h.respond(w, r, http.StatusUnprocessableEntity, v.Errors)
		return
	}

	account.Owner = input.Owner
	account.Balance = input.Balance

	err = h.service.AccountService.UpdateAccount(account)
	if err != nil {
		switch err {
		case models.ErrNotFound:
			h.error(w, r, http.StatusNotFound, err)
			return
		default:
			h.error(w, r, http.StatusInternalServerError, err)
			return
		}
	}

	h.respond(w, r, http.StatusOK, account)
}

func (h *Handler) DeleteAccount(w http.ResponseWriter, r *http.Request) {
	paramNumber := chi.URLParam(r, "accountNumber")

	err := h.service.AccountService.DeleteAccount(paramNumber)
	if err != nil {
		switch err {
		case models.ErrNotFound:
			h.error(w, r, http.StatusNotFound, err)
			return
		default:
			h.error(w, r, http.StatusInternalServerError, err)
			return
		}
	}

	resp := map[string]string{
		"message": "account deleted successfully",
	}

	h.respond(w, r, http.StatusOK, resp)
}
