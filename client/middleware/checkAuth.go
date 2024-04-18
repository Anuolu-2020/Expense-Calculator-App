package middleware

import (
	"log"
	"net/http"

	"github.com/alexedwards/scs/v2"
)

func CheckAuth(next http.HandlerFunc, m *scs.SessionManager) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		cookie := m.GetString(r.Context(), "username")
		log.Printf("This is the cookie: %v", cookie)
		if !m.Exists(r.Context(), "username") {
			http.Redirect(w, r, "/api/signin", http.StatusPermanentRedirect)
			return
		}

		next.ServeHTTP(w, r)
	})
}
