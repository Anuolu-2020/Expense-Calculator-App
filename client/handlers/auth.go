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

	decodedUser := Json.Decode(w, r.Body, &user)

	w.WriteHeader(http.StatusOK)

	response := Json.Encode(w, decodedUser)

	w.Write(response)
}
