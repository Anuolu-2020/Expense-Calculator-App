package handlers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/Anuolu-2020/Expense-Calculator-App/pkg"
)

type Report struct {
	Type   string      `json:"type"`
	Source string      `json:"source"`
	Amount json.Number `json:"amount"`
}

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

// Send a request to the backend api to create a report
func (h Handler) CreateUserReport(w http.ResponseWriter, r *http.Request) {
	// Get user id
	userId := r.PathValue("userId")

	var userRequest Report
	err := json.NewDecoder(r.Body).Decode(&userRequest)
	if err != nil {
		log.Printf("Error decoding json: %v", err)
		pkg.ServeErrorPage(w, r)
		return
	}

	// Check report type
	if userRequest.Type != "income" && userRequest.Type != "expense" {
		fmt.Println(w, "Type must be income or expense")
		return
	}

	// Check if amount is negative
	// Assuming userRequest.Amount is of type json.Number
	if amountInt, err := strconv.Atoi(userRequest.Amount.String()); err == nil && amountInt <= 0 {
		fmt.Println(w, "Amount must not be negative")
	}

	// Convert request body to json
	reqBody, err := json.Marshal(&userRequest)
	if err != nil {
		log.Printf("Error encoding json: %v", err)
		pkg.ServeErrorPage(w, r)
		return
	}

	resp, err := http.Post(
		"https://expense-calculator-api-j642.onrender.com/api/v1/report/"+userRequest.Type+"/"+userId,
		"application/json",
		bytes.NewBuffer(reqBody),
	)
	if err != nil {
		log.Printf("Error sending api request:%v", err)
		pkg.ServeErrorPage(w, r)
		return
	}

	defer resp.Body.Close()

	var userData Results

	json.NewDecoder(resp.Body).Decode(&userData)

	pkg.SendTemplate(w, "newReport.html", userData)
}
