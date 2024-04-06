package handlers

import (
	"net/http"

	"github.com/Anuolu-2020/Expense-Calculator-App/pkg"
)

type User struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func (h Handler) Auth(w http.ResponseWriter, r *http.Request) {
	Json := &pkg.Json{}

	var user User

	decodedUser, err := Json.Decode(r.Body, &user)

	if err != nil {
		http.Error(w, "Error occurred while decoding json", http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusOK)

	response, err := Json.Encode(decodedUser)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}

	w.Write(response)
}
