package handler

import (
	"net/http"
	"strconv"

	"bank_account/internal/models"
	"bank_account/pkg/validator"

	"github.com/go-chi/chi"
)

func (h *Handler) CreateAccount(w http.ResponseWriter, r *http.Request) {
	var input struct {
		Owner   string  `json:"owner"`
		Balance float64 `json:"balance"`
	}

	if err := h.readJSON(w, r, &input); err != nil {
		h.error(w, r, http.StatusBadRequest, err)
		return
	}

	account := &models.Account{
		Owner:   input.Owner,
		Balance: input.Balance,
	}

	v := validator.New()

	if err := account.ValidateAccount(v); err != nil {
		h.respond(w, r, http.StatusUnprocessableEntity, v.Errors)
		return
	}

	err := h.service.AccountService.CreateAccount(account)
	if err != nil {
		h.error(w, r, http.StatusInternalServerError, err)
		return
	}

	headers := make(http.Header)
	headers.Set("Location", "/accounts/"+account.AccountNumber)

	resp := map[string]*models.Account{
		"account_number": account,
	}

	h.respond(w, r, http.StatusCreated, resp, headers)
}

func (h *Handler) GetAccountByID(w http.ResponseWriter, r *http.Request) {
	paramId := chi.URLParam(r, "id")

	id, err := strconv.ParseInt(paramId, 10, 64)
	if err != nil {
		h.error(w, r, http.StatusBadRequest, err)
		return
	}

	account, err := h.service.AccountService.GetAccountByID(id)
	if err != nil {
		h.error(w, r, http.StatusInternalServerError, err)
		return
	}

	h.respond(w, r, http.StatusOK, account)
}

func (h *Handler) GetAccountByNumber(w http.ResponseWriter, r *http.Request) {
	paramNumber := chi.URLParam(r, "number")

	account, err := h.service.AccountService.GetAccountByNumber(paramNumber)
	if err != nil {
		h.error(w, r, http.StatusInternalServerError, err)
		return
	}

	h.respond(w, r, http.StatusOK, account)
}

func (h *Handler) GetAllAccounts(w http.ResponseWriter, r *http.Request) {
	paramPage := chi.URLParam(r, "page")

	page, err := strconv.Atoi(paramPage)
	if err != nil {
		h.error(w, r, http.StatusBadRequest, err)
		return
	}

	paramPageSize := chi.URLParam(r, "page_size")

	pageSize, err := strconv.Atoi(paramPageSize)
	if err != nil {
		h.error(w, r, http.StatusBadRequest, err)
		return
	}

	accounts, err := h.service.AccountService.GetAllAccounts(pageSize, page)
	if err != nil {
		h.error(w, r, http.StatusInternalServerError, err)
		return
	}

	metadata := map[string]int{
		"page":      page,
		"page_size": pageSize,
	}

	resp := map[string]interface{}{
		"metadata": metadata,
		"accounts": accounts,
	}

	h.respond(w, r, http.StatusOK, resp)
}

func (h *Handler) UpdateAccount(w http.ResponseWriter, r *http.Request) {
	paramNumber := chi.URLParam(r, "number")

	account, err := h.service.AccountService.GetAccountByNumber(paramNumber)
	if err != nil {
		h.error(w, r, http.StatusInternalServerError, err)
		return
	}

	var input struct {
		Owner   string  `json:"owner"`
		Balance float64 `json:"balance"`
	}

	if err := h.readJSON(w, r, &input); err != nil {
		h.error(w, r, http.StatusBadRequest, err)
		return
	}

	v := validator.New()

	if err := account.ValidateAccount(v); err != nil {
		h.respond(w, r, http.StatusUnprocessableEntity, v.Errors)
		return
	}

	account.Owner = input.Owner
	account.Balance = input.Balance

	err = h.service.AccountService.UpdateAccount(account)
	if err != nil {
		h.error(w, r, http.StatusInternalServerError, err)
		return
	}

	h.respond(w, r, http.StatusOK, account)
}

func (h *Handler) DeleteAccount(w http.ResponseWriter, r *http.Request) {
	paramNumber := chi.URLParam(r, "number")

	err := h.service.AccountService.DeleteAccount(paramNumber)
	if err != nil {
		h.error(w, r, http.StatusInternalServerError, err)
		return
	}

	h.respond(w, r, http.StatusOK, nil)
}
