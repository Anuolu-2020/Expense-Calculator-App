package handlers

import (
	"net/http"

	"github.com/alexedwards/scs/v2"
	"gorm.io/gorm"
)

type Handler struct {
	DB      *gorm.DB
	Session *scs.SessionManager
}

func New(DB *gorm.DB, Session *scs.SessionManager) Handler {
	return Handler{DB, Session}
}

func (h Handler) Home(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "templates/index.html")
}
