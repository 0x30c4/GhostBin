package middleware

import (
	"log/slog"
	"net/http"
	"time"
)

func LogRequest(handler http.Handler, logger *slog.Logger) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		// Wrap the ResponseWriter to capture the status code
		wrapped := &WrappedWriter{
			ResponseWriter: w,
			StatusCode:     http.StatusOK,
		}

		// Serve the request
		handler.ServeHTTP(wrapped, r)

		// Log the details using slog
		logger.Info("request completed",
			"remote_addr", r.RemoteAddr,
			"status", wrapped.StatusCode,
			"method", r.Method,
			"duration", time.Since(start).String(),
			"path", r.URL.Path,
		)
	})
}
