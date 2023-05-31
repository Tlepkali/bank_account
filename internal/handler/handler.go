package handler

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"

	"bank_account/internal/service"
)

type Handler struct {
	service *service.Service
}

func NewHandler(service *service.Service) *Handler {
	return &Handler{service: service}
}

func (h *Handler) readJSON(w http.ResponseWriter, r *http.Request, dst interface{}) error {
	maxBodyBytes := int64(1024 * 1024)
	r.Body = http.MaxBytesReader(w, r.Body, maxBodyBytes)

	dec := json.NewDecoder(r.Body)
	dec.DisallowUnknownFields()

	if err := dec.Decode(dst); err != nil {
		switch err.(type) {
		case *json.SyntaxError:
			return errors.New("invalid json")
		case *json.UnmarshalTypeError:
			return errors.New("invalid json type")
		default:
			return err
		}
	}

	if err := dec.Decode(&struct{}{}); err != io.EOF {
		return errors.New("json contains superfluous data")
	}

	return nil
}

func (h *Handler) error(w http.ResponseWriter, r *http.Request, code int, err error) {
	h.respond(w, r, code, map[string]string{"error": err.Error()})
}

func (h *Handler) respond(w http.ResponseWriter, r *http.Request, code int, data interface{}, headers ...http.Header) {
	w.WriteHeader(code)

	if data != nil {
		if err := json.NewEncoder(w).Encode(data); err != nil {
			h.error(w, r, http.StatusInternalServerError, err)
			return
		}
	}
}
