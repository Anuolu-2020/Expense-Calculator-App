package handlers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/alexedwards/scs/v2"
	"gorm.io/gorm"

	"github.com/Anuolu-2020/Expense-Calculator-App/pkg"
)

type Handler struct {
	DB      *gorm.DB
	Session *scs.SessionManager
}

type Results struct {
	ID        string `json:"id"`
	Source    string `json:"source"`
	Amount    int    `json:"amount"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
	UserId    string `json:"user_id"`
	Type      string `json:"type"`
}

type Response struct {
	Results []Results `json:"results"`
}

type TemplateData struct {
	UserId   string
	Username string
	Photo    string
	Results  []Results
}

func New(DB *gorm.DB, Session *scs.SessionManager) Handler {
	return Handler{DB, Session}
}

func (h *Handler) Auth(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "templates/oauthPage.html")
}

func (h *Handler) NotFound(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.ServeFile(w, r, "templates/notFound.html")
	}
	http.Redirect(w, r, "/dashboard", http.StatusFound)
}

func (h Handler) Dashboard(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	// Get user's username and photo
	sessionData := ctx.Value("sessionData").(string)

	userData, err := pkg.DecodeSessionData(sessionData)
	if err != nil {
		log.Printf("Error occured while decoding session: %v", err)
		http.Error(w, "An error occurred", http.StatusInternalServerError)
		return
	}

	resp, err := http.Get(
		"https://expense-calculator-api-j642.onrender.com/api/v1/report/" + userData.UserId,
	)
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

	data := TemplateData{
		UserId:   userData.UserId,
		Username: userData.Username,
		Photo:    userData.Photo,
		Results:  response.Results,
	}

	pkg.SendTemplate(w, "dashboard.html", data)
}
