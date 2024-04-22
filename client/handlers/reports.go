package handlers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/Anuolu-2020/Expense-Calculator-App/pkg"
)

func (h Handler) GetUserReports(w http.ResponseWriter, r *http.Request) {
	// user id
	userId := r.PathValue("userId")

	resp, err := http.Get(
		"https://expense-calculator-api-j642.onrender.com/api/v1/report/" + userId)
	if err != nil {
		log.Printf("Error occured while fetching data: %v", err)
		http.Error(w, "An error occurred", http.StatusInternalServerError)
		return
	}

	defer resp.Body.Close()

	var response Response
	if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
		log.Printf("Error occured while decoding data: %v", err)
		http.Error(w, "An error occurred", http.StatusInternalServerError)
		return
	}

	// Send template
	pkg.SendTemplate(w, "reports.html", response)
}
