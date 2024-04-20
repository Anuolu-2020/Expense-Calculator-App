package handlers

import (
	"net/http"

	"github.com/alexedwards/scs/v2"
	"gorm.io/gorm"

	"github.com/Anuolu-2020/Expense-Calculator-App/pkg"
)

type Handler struct {
	DB      *gorm.DB
	Session *scs.SessionManager
}

func New(DB *gorm.DB, Session *scs.SessionManager) Handler {
	return Handler{DB, Session}
}

func (h *Handler) SignIn(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "templates/signIn.html")
}

func (h *Handler) SignUp(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "templates/signUp.html")
}

func (h *Handler) NotFound(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.ServeFile(w, r, "templates/notFound.html")
	}
	http.Redirect(w, r, "/dashboard", http.StatusFound)
}

func (h *Handler) Dashboard(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	// Get user's username
	username := ctx.Value("profileName")

	type templateData struct{ Username any }

	data := templateData{Username: username}

	pkg.SendTemplate(w, "dashboard.html", data)
}
