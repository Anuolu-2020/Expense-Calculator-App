package handlers

import (
	"encoding/json"
	"log"
	"net/http"
)

func (h Handler) Welcome(w http.ResponseWriter, r *http.Request) {
	message := "welcome to expense calculator welcome route"

	responseMessage, err := json.Marshal(message)
	if err != nil {
		log.Fatal(err)
	}

	w.Header().Set("Content-Type", "application/json")

	w.Write([]byte(responseMessage))
}
