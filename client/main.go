package main

import (
	"embed"
	"encoding/json"
	"log"
	"net/http"
)

func main() {
	//go:generate npm run build

	var static embed.FS

	router := http.NewServeMux()

	router.Handle("/static", http.FileServer(http.FS(static)))

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
