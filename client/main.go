package main

import (
	"embed"
	"log"
	"net/http"

	"github.com/Anuolu-2020/Expense-Calculator-App/controllers"
)

//go:embed static
var static embed.FS

func main() {
	//go:generate npm run build

	router := http.NewServeMux()

	router.Handle("/static/", http.FileServer(http.FS(static)))

	router.HandleFunc("/", controllers.Home)

	router.HandleFunc("GET /welcome", controllers.Welcome)

	server := &http.Server{
		Addr:    ":8080",
		Handler: router,
	}

	log.Printf("Server listening on port %v", server.Addr)

	server.ListenAndServe()
}
