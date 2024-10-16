package webappsrv

import (
	"errors"
	"net/http"
	"os"
	"time"

	"github.com/0x30c4/ghostbin/internal/services"
	"github.com/0x30c4/ghostbin/pkg/webappsrv/handlers"
	"golang.org/x/exp/slog"
)

func newServer() *http.Server {
	port := "8080"
	if p := os.Getenv("PORT"); p != "" {
		port = p
	}
  host := "localhost"
	if h := os.Getenv("HOST"); h != "" {
    host = h
	}

	return &http.Server{
		Addr:              host + ":" + port,
		ReadTimeout:       5 * time.Second,
		WriteTimeout:      5 * time.Second,
		ReadHeaderTimeout: 2 * time.Second,
	}
}

func SetupRouter(pasteSrv *services.PasteService, fileSrv *services.FileService, logger *slog.Logger) http.Handler {

  gbinHandlerSrv := handlers.NewHttpHandler(pasteSrv, fileSrv, logger)
  gbinHandler := gbinHandlerSrv.SetupRoutesWithLogging()

  return gbinHandler
}

func RunServer(pasteSrv *services.PasteService, fileSrv *services.FileService, logger *slog.Logger) {

	router := SetupRouter(pasteSrv, fileSrv, logger)
	if router == nil {
		return
	}

	server := newServer()
	server.Handler = router

	if err := server.ListenAndServe(); err != nil {
		if !errors.Is(err, http.ErrServerClosed) {
			logger.Error("Failed to run server", slog.String("err", err.Error()))
		}
	}
}
