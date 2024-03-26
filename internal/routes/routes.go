package routes

import (
	"net/http"

	"github.com/0x30c4/ghostbin/internal/handlers"
)

func Router() http.Handler {
  mux := &http.ServeMux{}

  mux.HandleFunc("GET /", handlers.Index)
  mux.HandleFunc("POST /", handlers.NewPaste)
  mux.HandleFunc("GET /{id}", handlers.GetPaste)
  mux.HandleFunc("PUT /{id}", handlers.EditPaste)
  mux.HandleFunc("DELETE /{id}", handlers.DeletePaste)

  return mux
}
