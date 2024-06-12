package handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/Anuolu-2020/Expense-Calculator-App/pkg"
)

type summaryReport struct {
	TotalIncome  json.Number `json:"totalIncome"`
	TotalExpense json.Number `json:"totalExpense"`
	NetIncome    json.Number `json:"netIncome"`
}

func GetSummary(userId string) (*http.Response, error) {
	resp, err := http.Get(
		fmt.Sprintf(
			"https://expense-calculator-api-j642.onrender.com/api/v1/summary/%s", userId,
		),
	)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (h Handler) GetReportsSummary(w http.ResponseWriter, r *http.Request) {
	var response summaryReport

	userId := r.PathValue("userId")

	resp, err := GetSummary(userId)
	if err != nil {
		log.Printf("An Error Occurred: %v", err)
		pkg.SendErrorResponse(w, "An Error Occurred", http.StatusInternalServerError)
		return
	}

	defer resp.Body.Close()

	if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
		log.Printf("An Error Ocurred: %v", err)
		pkg.SendErrorResponse(w, "An Error Occurred", http.StatusInternalServerError)
		return
	}

	pkg.SendTemplate(w, "summary.html", response)
}
