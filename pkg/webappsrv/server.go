package webappsrv

import (
	"errors"
	"net/http"
	"os"
	"time"

	"github.com/0x30c4/ghostbin/internal/services"
	"github.com/0x30c4/ghostbin/pkg/webappsrv/handlers"
	"log/slog"
)

// newServer initializes and returns a pointer to an http.Server using environment
// variables for HOST and PORT if they are set. It sets sensible timeouts and builds
// the server address from the resolved host and port values.
func newServer() *http.Server {
	port := "8080"
	if p := os.Getenv("PORT"); p != "" {
		port = p
	}

	host, err := os.Hostname()
	h := os.Getenv("HOST")
	if h != "" && err == nil {
		host = h
	}

	return &http.Server{
		Addr:              host + ":" + port,
		ReadTimeout:       5 * time.Second,
		WriteTimeout:      5 * time.Second,
		ReadHeaderTimeout: 2 * time.Second,
	}
}

// SetupRouter configures and returns an HTTP handler with routing and logging.
// It initializes a new HTTP handler using the provided paste and file services along
// with the given logger, and sets up the routes with logging enabled.
func SetupRouter(pasteSrv *services.PasteService, fileSrv *services.FileService, logger *slog.Logger) http.Handler {

	gbinHandlerSrv := handlers.NewHttpHandler(pasteSrv, fileSrv, logger)
	gbinHandler := gbinHandlerSrv.SetupRoutesWithLogging()

	return gbinHandler
}

// RunServer initializes and runs the HTTP server. It sets up the router with the
// provided services and logger, assigns it to the server handler, and starts the server.
// Logs an error if the server fails to start for reasons other than a graceful shutdown.
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
