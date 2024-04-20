package middleware

import (
	"context"
	"net/http"

	"github.com/alexedwards/scs/v2"
)

func CheckAuth(next http.HandlerFunc, m *scs.SessionManager) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		cookie := m.GetString(r.Context(), "username")

		if !m.Exists(r.Context(), "username") {
			http.Redirect(w, r, "/sign-in", http.StatusFound)
			return
		}

		// Pass username to next handler e.g dashboard
		ctx := context.WithValue(r.Context(), "profileName", cookie)

		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
