package main

import (
	"embed"
	"log"
	"net/http"

	"github.com/joho/godotenv"

	"github.com/Anuolu-2020/Expense-Calculator-App/db"
	"github.com/Anuolu-2020/Expense-Calculator-App/handlers"
	"github.com/Anuolu-2020/Expense-Calculator-App/middleware"
	"github.com/Anuolu-2020/Expense-Calculator-App/routes"
)

//go:embed static
var static embed.FS

func main() {
	//go:generate npm run build

	// Load env files
	err := godotenv.Load()
	if err != nil {
		log.Fatal(err)
	}

	router := http.NewServeMux()

	// Serve static files
	router.Handle("/static/", http.FileServer(http.FS(static)))

	setUpRoutes := routes.SetupRoute{}

	// connect to db
	DB := db.Init()

	// Initialize Handlers
	h := handlers.New(DB)

	// Setup and initialize routes
	setUpRoutes.New(router)
	setUpRoutes.InitializeRoutes(h)

	server := &http.Server{
		Addr:    ":8080",
		Handler: middleware.LoggerMiddleware(router),
	}

	log.Printf("Server listening on port %v", server.Addr)

	server.ListenAndServe()
}
