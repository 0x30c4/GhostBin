package handlers

import (
	"net/http"
	"github.com/0x30c4/ghostbin/internal/services"
	"github.com/0x30c4/ghostbin/pkg/webappsrv/middleware"
	"golang.org/x/exp/slog"
)

type HttpHandlers struct {
  fileSrv     *services.FileService
  pasteSrv    *services.PasteService
  logger      *slog.Logger
}

func NewHttpHandler(pasteSrv *services.PasteService, fileSrv *services.FileService, logger *slog.Logger) *HttpHandlers {
  return &HttpHandlers{
    pasteSrv: pasteSrv,
    fileSrv: fileSrv,
    logger: logger,
  }
}


func (h *HttpHandlers) SetupRoutesWithLogging() http.Handler {
  mux := &http.ServeMux{}

  mux.HandleFunc("GET /", h.Index)
  mux.HandleFunc("GET /{id}", h.GetPaste)
  mux.HandleFunc("POST /", h.NewPaste)
  // mux.HandleFunc("PUT /{id}", h.EditPaste)
  mux.HandleFunc("DELETE /{id}", h.DeletePaste)

	return middleware.LogRequest(mux, h.logger)
}
