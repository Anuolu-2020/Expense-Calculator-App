package handlers

import (
	"log"
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

func (h *Handler) Auth(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "templates/oauthPage.html")
}

func (h *Handler) NotFound(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.ServeFile(w, r, "templates/notFound.html")
	}
	http.Redirect(w, r, "/dashboard", http.StatusFound)
}

func (h *Handler) Dashboard(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	// Get user's username and photo
	sessionData := ctx.Value("sessionData").(string)

	data, err := pkg.DecodeSessionData(sessionData)
	if err != nil {
		log.Printf("Error occured while decoding session: %v", err)
		http.Error(w, "An error occurred", http.StatusInternalServerError)
		return
	}

	pkg.SendTemplate(w, "dashboard.html", data)
}
