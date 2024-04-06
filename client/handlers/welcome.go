package handlers

import (
	"encoding/json"
	"log"
	"net/http"
)

func Welcome(w http.ResponseWriter, r *http.Request) {
	message := "welcome to expense calculator welcome route"

	responseMessage, err := json.Marshal(message)

	if err != nil {
		log.Fatal(err)
	}

	w.Write([]byte(responseMessage))
}
