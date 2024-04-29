package v1

import (
	"fmt"
	"github.com/go-chi/chi/v5"
	"log"
	"net/http"
)

type Handler struct {
}

func NewHandler() *Handler {
	return &Handler{}
}

func (h *Handler) helloHandler() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if _, err := fmt.Fprintf(w, "Hello, World!"); err != nil {
			log.Printf("Error writing response: %v", err)
		}
	})
}

func (h *Handler) Routes() *chi.Mux {
	r := chi.NewRouter()
	r.Get("/hello", h.helloHandler().ServeHTTP)
	return r
}
