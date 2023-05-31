package handler

import (
	"net/http"

	"github.com/go-chi/chi"
)

func (h *Handler) InitRoutes() http.Handler {
	router := chi.NewRouter()

	router.Post("/account", h.CreateAccount)
	router.Get("/account/{id}", h.GetAccount)
	router.Put("/account/{id}", h.UpdateAccount)
	router.Delete("/account/{id}", h.DeleteAccount)

	return router
}
