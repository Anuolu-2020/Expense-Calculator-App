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
func (h Handler) ApiSignUp(w http.ResponseWriter, r *http.Request) {
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

	h.Session.Put(r.Context(), "username", newUser.Username)

	newUser.Password = "" // Exclude password

	if r.Header.Get("Hx-Request") != "" {
		w.Header().Set("Hx-Redirect", "/dashboard")
		w.Write([]byte("Redirect to dashboard"))

	}

	http.Redirect(w, r, "/dashboard", http.StatusMovedPermanently)
}

// Sign In route
func (h Handler) ApiSignIn(w http.ResponseWriter, r *http.Request) {
	var user User

	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		pkg.SendErrorResponse(
			w,
			"An Error Occurred while Decoding Json",
			http.StatusInternalServerError,
		)
		log.Print("An error occurred while encoding json %w", err)
		return
	}

	if ok, errors := pkg.ValidateInputs(user); !ok {
		pkg.ValidationError(w, http.StatusUnprocessableEntity, errors)
		return
	}

	var userFound models.User
	result := h.DB.Where("username = ?", user.Username).First(&userFound)

	if result.RowsAffected == 0 {
		pkg.SendErrorResponse(w, "User Not Found", http.StatusNotFound)
		return
	}

	if result.Error != nil {
		log.Print(result.Error)
		pkg.SendErrorResponse(w, "An error occurred", http.StatusInternalServerError)
		return
	}

	if err := bcrypt.CompareHashAndPassword([]byte(userFound.Password), []byte(user.Password)); err != nil {
		log.Printf("Error comparing password: %v", err)
		pkg.SendErrorResponse(w, "Username or password not correct", http.StatusUnauthorized)
		return
	}

	h.Session.RenewToken(r.Context())

	h.Session.Put(r.Context(), "username", user.Username)

	userFound.Password = ""

	if r.Header.Get("Hx-Request") != "" {
		w.Header().Set("Hx-Redirect", "/dashboard")
		w.Write([]byte("Redirect to dashboard"))

	}

	http.Redirect(w, r, "/dashboard", http.StatusMovedPermanently)
}
