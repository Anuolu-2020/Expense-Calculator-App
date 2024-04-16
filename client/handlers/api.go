package handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"

	"github.com/Anuolu-2020/Expense-Calculator-App/models"
	"github.com/Anuolu-2020/Expense-Calculator-App/pkg"
)

type User struct {
	Username string `json:"username" validate:"required,min=5,max=15"`
	Password string `json:"password" validate:"required"`
}

// Sign up handler
func (h Handler) SignUp(w http.ResponseWriter, r *http.Request) {
	var user User

	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, "Invalid Json Object", http.StatusBadRequest)
		log.Println(err.Error())
		return
	}

	// Validate User Inputs
	if ok, errors := pkg.ValidateInputs(user); !ok {
		pkg.ValidationError(w, http.StatusUnprocessableEntity, errors)
		return
	}

	var userFound models.User
	result := h.DB.Where("username = ?", user.Username).First(&userFound)

	// If a user is found
	if result.RowsAffected > 0 {
		pkg.SendErrorResponse(w, "User Already Exists", http.StatusBadRequest)
		return
	}

	// If its a db error
	if result.Error != gorm.ErrRecordNotFound {
		pkg.SendErrorResponse(
			w,
			"An Error Occurred",
			http.StatusInternalServerError,
		)
		log.Println(result.Error)
		return
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), 10)
	if err != nil {
		pkg.SendErrorResponse(w, "An error occurred", http.StatusInternalServerError)
		log.Println(err)
		return
	}

	newUser := &models.User{Username: user.Username, Password: string(hashedPassword)}

	if result := h.DB.Create(&newUser); result.Error != nil {
		pkg.SendErrorResponse(
			w,
			"An Error Occurred While Creating User",
			http.StatusInternalServerError,
		)
		fmt.Println(result.Error)
		return
	}

	pkg.SendResponse(w, "User created successfully", newUser, http.StatusCreated)
}
