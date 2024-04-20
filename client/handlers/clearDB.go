package handlers

import (
	"log"
	"net/http"

	"gorm.io/gorm"

	"github.com/Anuolu-2020/Expense-Calculator-App/models"
)

func (h Handler) ClearDB(w http.ResponseWriter, r *http.Request) {
	if result := h.DB.Session(&gorm.Session{AllowGlobalUpdate: true}).Delete(models.User{}); result.Error != nil {
		log.Printf("Error clearing db: %v", result.Error)
		http.Error(
			w,
			"Error clearing db",
			http.StatusInternalServerError,
		)
		return
	}

	w.Write([]byte("Users cleared successfully"))
}
