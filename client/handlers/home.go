package handlers

import "net/http"

type Handler struct {
}

func (h Handler) Home(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "templates/index.html")
}
