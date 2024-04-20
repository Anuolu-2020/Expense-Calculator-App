package middleware

import (
	"context"
	"log"
	"net/http"

	"github.com/alexedwards/scs/v2"
)

func CheckAuth(next http.HandlerFunc, m *scs.SessionManager) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		cookie := m.GetString(r.Context(), "userprofile")

		if !m.Exists(r.Context(), "userprofile") {
			log.Println("it doesn't exist")
			http.Redirect(w, r, "/auth", http.StatusFound)
			return
		}

		// Pass username to next handler e.g dashboard
		ctx := context.WithValue(r.Context(), "sessionData", cookie)

		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
