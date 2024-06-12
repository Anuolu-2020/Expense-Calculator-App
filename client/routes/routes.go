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
	r.mux.HandleFunc("/", middleware.CheckAuth(handler.NotFound, SessionManager))
	r.mux.HandleFunc("GET /auth", middleware.IsLoggedIn(handler.Auth, SessionManager))
	r.mux.HandleFunc("GET /welcome", middleware.CheckAuth(handler.Welcome, SessionManager))
	r.mux.HandleFunc("GET /dashboard", middleware.CheckAuth(handler.Dashboard, SessionManager))
	r.mux.HandleFunc("GET /reports-graph", middleware.CheckAuth(handler.Graph, SessionManager))
	r.mux.HandleFunc(
		"POST /reportschart/{userId}",
		middleware.CheckAuth(handler.ReportChart, SessionManager),
	)
	r.mux.HandleFunc(
		"POST /reportschart/summary/{userId}",
		middleware.CheckAuth(handler.ReportSummaryChart, SessionManager),
	)

	apiRoutes := http.NewServeMux()

	// Api Routes
	apiRoutes.HandleFunc("POST /google", handler.ApiGoogle)
	apiRoutes.HandleFunc("GET /google/callback", handler.ApiGoogleCallback)
	apiRoutes.HandleFunc("GET /cleardb", handler.ClearDB)
	apiRoutes.HandleFunc("GET /reports/{userId}", handler.GetUserReports)
	apiRoutes.HandleFunc(
		"GET /reports/summary/{userId}",
		middleware.CheckAuth(handler.GetReportsSummary, SessionManager),
	)

	apiRoutes.HandleFunc("POST /createUserReport/{userId}", handler.CreateUserReport)

	r.mux.Handle("/api/", http.StripPrefix("/api", apiRoutes))
}
