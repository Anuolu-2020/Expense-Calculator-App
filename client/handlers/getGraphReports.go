package handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func GetReports(userId string) (*Response, error) {
	resp, err := http.Get(
		"https://expense-calculator-api-j642.onrender.com/api/v1/report/" + userId)
	if err != nil {
		log.Printf("Error occured while fetching data: %v", err)
		return nil, err
	}

	defer resp.Body.Close()

	var response Response
	if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
		log.Printf("Error occured while decoding data: %v", err)
		// http.Error(w, "An error occurred", http.StatusInternalServerError)
		return nil, err
	}

	return &response, nil
}

func GetReportsType(reportType string, userId string) (*Response, error) {
	url := fmt.Sprintf(
		"https://expense-calculator-api-j642.onrender.com/api/v1/report/type/%s/%s",
		reportType,
		userId,
	)

	resp, err := http.Get(url)
	if err != nil {
		log.Printf("Error occured while fetching data: %v", err)
		return nil, err
	}

	defer resp.Body.Close()

	var response Response
	if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
		log.Printf("Error occured while decoding data: %v", err)
		// http.Error(w, "An error occurred", http.StatusInternalServerError)
		return nil, err
	}

	return &response, nil
}
