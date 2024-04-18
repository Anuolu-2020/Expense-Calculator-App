package routes

import (
	"net/http"

	"github.com/alexedwards/scs/v2"

	"github.com/Anuolu-2020/Expense-Calculator-App/handlers"
	"github.com/Anuolu-2020/Expense-Calculator-App/middleware"
)

type SetupRoute struct {
	mux *http.ServeMux
}

func (r *SetupRoute) New(newMux *http.ServeMux) {
	r.mux = newMux
}

func (r SetupRoute) InitializeRoutes(handler handlers.Handler, SessionManager *scs.SessionManager) {
	// main routes
	r.mux.HandleFunc("GET /index", handler.Home)
	r.mux.HandleFunc("GET /welcome", middleware.CheckAuth(handler.Welcome, SessionManager))

	apiRoutes := http.NewServeMux()

	// Api Routes
	apiRoutes.HandleFunc("POST /signup", handler.SignUp)
	apiRoutes.HandleFunc("POST /signin", handler.SignIn)

	r.mux.Handle("/api/", http.StripPrefix("/api", apiRoutes))
}
