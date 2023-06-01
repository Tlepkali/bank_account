package handler

import (
	"net/http"

	"github.com/go-chi/chi"

	_ "bank_account/docs"

	httpSwagger "github.com/swaggo/http-swagger"
)

func (h *Handler) InitRoutes() http.Handler {
	router := chi.NewRouter()

	router.Post("/account", h.CreateAccount)
	router.Get("/account/{accountNumber}", h.GetAccountByNumber)
	router.Get("/accounts", h.GetAllAccounts)
	router.Put("/account/{accountNumber}", h.UpdateAccount)
	router.Delete("/account/{accountNumber}", h.DeleteAccount)

	router.Get("/swagger/*", httpSwagger.Handler(
		httpSwagger.URL("/swagger/doc.json"), // The url pointing to API definition
	))

	return router
}
