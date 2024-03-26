package middleware

import (
	"net/http"
)

type MiddleWare func(http.Handler) http.Handler

type WrappedWriter struct {
  http.ResponseWriter
  StatusCode int
}

func (w *WrappedWriter) WriteHeader(statusCode int) {
  w.ResponseWriter.WriteHeader(statusCode)
  w.StatusCode = statusCode
}

func CreateStack(xs ...MiddleWare) MiddleWare {
  return func(next http.Handler) http.Handler {
    for i := len(xs) - 1; i >= 0; i-- {
      x := xs[i]
      next = x(next)
    }
    return next
  }
}
