package logger

import (
	"log"
	"net/http"
	"os"
	"time"

	"github.com/0x30c4/ghostbin/internal/middleware"
)

func LogRequest(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
    start := time.Now()
    wrapped := &middleware.WrappedWriter{
      ResponseWriter: w,
      StatusCode: http.StatusOK,
    }
		handler.ServeHTTP(wrapped, r)
		log.Printf("%s [%d] [%s] [%s] %s\n", r.RemoteAddr, wrapped.StatusCode, r.Method, time.Since(start), r.URL.Path)
	})
}

func LoggerInit(logfile string, logflag int) {
	if logfile != "" {
		lf, err := os.OpenFile(logfile, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0640)

		if err != nil {
			log.Fatal("OpenLogfile: os.OpenFile:", err)
		}

		log.SetOutput(lf)
	}

	log.SetFlags(logflag)
}
