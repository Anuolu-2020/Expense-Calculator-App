package handlers

import (
	"context"
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

// Google endpoint handler
func (h Handler) ApiGoogle(w http.ResponseWriter, r *http.Request) {
	oauthState := pkg.GenerateStateOauthCookie(w)

	url := pkg.OauthConfig.GoogleLoginConfig.AuthCodeURL(oauthState)

	if r.Header.Get("Hx-Request") != "" {
		w.Header().Set("Hx-Redirect", url)
		w.Write([]byte("Redirect to dashboard"))
	}

	http.Redirect(w, r, url, http.StatusMovedPermanently)
}

// Google Calback handler
func (h Handler) ApiGoogleCallback(w http.ResponseWriter, r *http.Request) {
	oauthState, _ := r.Cookie("oauthstate")

	failureRedirect := ""

	if r.FormValue("state") != oauthState.Value {
		log.Println("[AUTH]: Oauth states do not match")
		http.Redirect(w, r, failureRedirect, http.StatusTemporaryRedirect)
		return
	}

	oauthCode := r.FormValue("code")

	// Get response from google
	resp := getGoogleResponse(w, r, oauthCode)

	// Read user data
	defer resp.Body.Close()

	var userData struct {
		Email      string `json:"email"`
		Username   string `json:"name"`
		Password   string `json:"id"`
		ProfilePic string `json:"picture"`
	}

	// Read response into struct
	if err := json.NewDecoder(resp.Body).Decode(&userData); err != nil {
		log.Println("failed to decode response: %w", err)
		http.Redirect(w, r, failureRedirect, http.StatusTemporaryRedirect)
		return
	}

	// Check if user exists
	var userFound models.User
	result := h.DB.Where("email = ?", userData.Email).First(&userFound)

	// If a user is found Sign in user
	if result.RowsAffected > 0 {
		if err := bcrypt.CompareHashAndPassword([]byte(userFound.Password), []byte(userData.Password)); err != nil {
			log.Printf("Error comparing password: %v", err)
			pkg.SendErrorResponse(w, "Username or password not correct", http.StatusUnauthorized)
			return
		}

		// Renew session token
		h.Session.RenewToken(r.Context())

		// Encode session data
		buf, err := pkg.EncodeSessionData(
			userFound.ID.String(),
			userFound.Username,
			userFound.ProfilePic,
		)
		if err != nil {
			log.Printf("Error encoding session: %v", err)
			http.Error(w, "An error occcured", http.StatusInternalServerError)
			return
		}

		// Store Session data
		h.Session.Put(r.Context(), "userSession", buf.String())

		// if r.Header.Get("Hx-Request") != "" {
		// 	w.Header().Set("Hx-Redirect", "/dashboard")
		//
		// 	w.Write([]byte("Redirect to dashboard"))
		//
		// 	return
		// }
		//
		http.Redirect(w, r, "/dashboard", http.StatusMovedPermanently)

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

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(userData.Password), 10)
	if err != nil {
		pkg.SendErrorResponse(w, "An error occurred", http.StatusInternalServerError)
		log.Println(err)
		return
	}

	newUser := &models.User{
		Email:      userData.Email,
		Username:   userData.Username,
		Password:   string(hashedPassword),
		ProfilePic: userData.ProfilePic,
	}

	// Create new user
	if result := h.DB.Create(&newUser); result.Error != nil {
		pkg.SendErrorResponse(
			w,
			"An Error Occurred While Creating User",
			http.StatusInternalServerError,
		)
		fmt.Println(result.Error)
		return
	}

	buf, err := pkg.EncodeSessionData(newUser.ID.String(), newUser.Username, newUser.ProfilePic)
	if err != nil {
		log.Printf("Error encoding session: %v", err)
		http.Error(w, "An error occcured", http.StatusInternalServerError)
		return
	}

	h.Session.Put(
		r.Context(),
		"userSession",
		buf.String(),
	)

	// if r.Header.Get("Hx-Request") != "" {
	// 	w.Header().Set("Hx-Redirect", "/dashboard")
	// 	w.Write([]byte("Redirect to dashboard"))
	// 	return
	// }
	//
	http.Redirect(w, r, "/dashboard", http.StatusMovedPermanently)
}

// Get google response function
func getGoogleResponse(
	w http.ResponseWriter,
	r *http.Request,
	oauthCode string,
) *http.Response {
	failureRedirect := ""

	googlecon := pkg.InitGoogle()

	// Get token from google
	token, err := googlecon.Exchange(context.Background(), oauthCode)
	if err != nil {
		log.Println("[FAIL]:", "Code-Token Exchange Failed")
		http.Redirect(w, r, failureRedirect, http.StatusTemporaryRedirect)
		return nil
	}

	resp, err := http.Get(
		"https://www.googleapis.com/oauth2/v2/userinfo?access_token=" + token.AccessToken,
	)
	if err != nil {
		log.Println("User Data Fetch Failed")
		http.Redirect(w, r, failureRedirect, http.StatusTemporaryRedirect)
		return nil
	}

	return resp
}
