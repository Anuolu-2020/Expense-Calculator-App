package pkg

import (
	"encoding/json"
	"log"
	"net/http"
)

type Response struct {
	Status  string      `json:"status"`
	Message string      `json:"message"`
	Payload interface{} `json:"payload"`
}

type ResponseError struct {
	Status  string `json:"status"`
	Message string `json:"message"`
}

func SendResponse(w http.ResponseWriter, msg string, body interface{}, HttpCode int) {
	res := Response{Status: "success", Message: msg, Payload: body}

	response, err := json.Marshal(res)
	if err != nil {
		log.Print("Error encoding json: %w", err)
		http.Error(w, "An Error Occurred", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(HttpCode)
	w.Write(response)
}

func SendErrorResponse(w http.ResponseWriter, msg string, HttpCode int) {
	res := ResponseError{Status: "fail", Message: msg}

	response, err := json.Marshal(res)
	if err != nil {
		log.Print("Error encoding json: %w", err)
		http.Error(w, "An Error Occurred", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(HttpCode)
	w.Write(response)
}
