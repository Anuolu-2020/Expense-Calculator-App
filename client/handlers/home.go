package handlers

import (
	"net/http"

	"gorm.io/gorm"
)

type Handler struct {
	DB *gorm.DB
}

func New(DB *gorm.DB) Handler {
	return Handler{DB}
}

func (h Handler) Home(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "templates/index.html")
}
