package pkg

import "net/http"

func ServeErrorPage(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "templates/errors.html")
}
