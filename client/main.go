package main

import (
	"embed"
	"encoding/json"
	"log"
	"net/http"
)

//go:embed static
var static embed.FS

func main() {
	//go:generate npm run build

	router := http.NewServeMux()

	router.Handle("/static/", http.FileServer(http.FS(static)))

	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "templates/index.html")
	})

	router.HandleFunc("GET /home", func(w http.ResponseWriter, r *http.Request) {
		message := "welcome to expense calculator home route"

		responseMessage, err := json.Marshal(message)

		if err != nil {
			log.Fatal(err)
		}

		w.Write([]byte(responseMessage))
	})

	server := &http.Server{
		Addr:    ":8080",
		Handler: router,
	}

	log.Printf("Server listening on port %v", server.Addr)

	server.ListenAndServe()
}
