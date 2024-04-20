package middleware

import (
	"context"
	"log"
	"net/http"
	"time"

	"github.com/alexedwards/scs/v2"
)

func clearCookie(w http.ResponseWriter) {
	cookie := &http.Cookie{
		Name:    "session",
		Value:   "",         // Clear the cookie value
		MaxAge:  -1,         // Set MaxAge to negative value to expire immediately
		Expires: time.Now(), // Set Expires to a time in the past
		Path:    "/",        // Specify the path of the cookie you want to delete
	}
	http.SetCookie(w, cookie)
}

func CheckAuth(next http.HandlerFunc, m *scs.SessionManager) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		cookie := m.GetString(r.Context(), "userSession")

		if exists := m.Exists(r.Context(), "userSession"); !exists {
			log.Printf("Session exists?:%v", exists)
			m.Destroy(r.Context()) //  Delete session
			clearCookie(w)         // Clear cookie
			http.Redirect(w, r, "/auth", http.StatusFound)
			return
		}

		// Pass username to next handler e.g dashboard
		ctx := context.WithValue(r.Context(), "sessionData", cookie)

		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
