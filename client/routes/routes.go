package routes

import (
	"net/http"

	"github.com/Anuolu-2020/Expense-Calculator-App/handlers"
)

type SetupRoute struct {
	mux *http.ServeMux
}

func (r *SetupRoute) New(newMux *http.ServeMux) {
	r.mux = newMux
}

func (r SetupRoute) InitializeRoutes(handler handlers.Handler) {
	// main routes
	r.mux.HandleFunc("GET /index", handler.Home)
	r.mux.HandleFunc("GET /welcome", handler.Welcome)
	r.mux.HandleFunc("POST /signup", handler.SignUp)

	apiRoutes := http.NewServeMux()

	// Api Rooutes
	apiRoutes.HandleFunc("POST /signup", handler.SignUp)

	r.mux.Handle("/api/", http.StripPrefix("/api", apiRoutes))
}
