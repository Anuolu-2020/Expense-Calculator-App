package middleware

import (
	"log"
	"net/http"
	"time"
)

type Middleware func(http.Handler) http.Handler

func MiddlewareStack(xs ...Middleware) Middleware {
	return func(next http.Handler) http.Handler {
		for i := len(xs) - 1; i >= 0; i-- {
			x := xs[i]
			next = x(next)
		}

		return next
	}
}

type wrappedWriter struct {
	http.ResponseWriter
	statusCode int
}

func (w *wrappedWriter) WriteHeader(statusCode int) {
	w.ResponseWriter.WriteHeader(statusCode)
	w.statusCode = statusCode
}

func statusColor(wrapped *wrappedWriter) string {
	// statusColors
	Red := "\033[31m"
	Green := "\033[32m"
	Yellow := "\033[33m"
	Blue := "\033[34m"
	Cyan := "\033[36m"

	if wrapped.statusCode > 500 {
		return Red
	} else if wrapped.statusCode >= 400 {
		return Yellow
	} else if wrapped.statusCode >= 300 {
		return Blue
	} else if wrapped.statusCode >= 200 {
		return Green
	} else {
		return Cyan
	}
}

func methodColor(r *http.Request) string {
	var methodString string

	switch r.Method {
	case "POST":
		methodString = "\033[33m[POST]\033[0m"
	case "PUT":
		methodString = "\033[34m[PUT]\033[0m"
	case "DELETE":
		methodString = "\033[31m[DELETE]\033[0m"
	default:
		methodString = "\033[32m[GET]\033[0m"
	}

	return methodString
}

func LoggerMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		wrapped := &wrappedWriter{
			ResponseWriter: w,
			statusCode:     http.StatusOK,
		}

		next.ServeHTTP(wrapped, r)

		color := statusColor(wrapped)
		Reset := "\033[0m" // Reset Color

		log.Printf(
			"%s %s %s%v%s %v - %s",
			methodColor(r),
			r.URL.Path,
			color,
			wrapped.statusCode,
			Reset,
			r.ContentLength,
			time.Since(start),
		)
	})
}
