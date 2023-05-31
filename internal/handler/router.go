package handler

import (
	"net/http"

	"github.com/go-chi/chi"
)

func (h *Handler) InitRoutes() http.Handler {
	router := chi.NewRouter()

	router.Post("/account", h.CreateAccount)
	router.Get("/account/{id}", h.GetAccountByID)
	router.Get("/account/{accountNumber}", h.GetAccountByNumber)
	router.Get("/accounts", h.GetAllAccounts)
	router.Put("/account/{accountNumber}", h.UpdateAccount)
	router.Delete("/account/{accountNumber}", h.DeleteAccount)

	return router
}
