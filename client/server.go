package main

import (
	"embed"
	"log"
	"net/http"

	"github.com/joho/godotenv"

	"github.com/Anuolu-2020/Expense-Calculator-App/db"
	"github.com/Anuolu-2020/Expense-Calculator-App/handlers"
	"github.com/Anuolu-2020/Expense-Calculator-App/middleware"
	"github.com/Anuolu-2020/Expense-Calculator-App/pkg"
	"github.com/Anuolu-2020/Expense-Calculator-App/routes"
)

//go:embed static
var static embed.FS

func main() {
	//go:generate npm run build

	// Load env files
	err := godotenv.Load()
	if err != nil {
		log.Println("Error loading env files")
		log.Fatal(err)
	}

	router := http.NewServeMux()

	// Serve static files
	router.Handle("/static/", http.FileServer(http.FS(static)))

	setUpRoutes := routes.SetupRoute{}

	// connect to db
	DB := db.Init()

	// Initailize Session
	sessionManager := middleware.InitSession()

	// Initialize Google Config
	pkg.InitGoogle()

	// Initialize Handlers
	h := handlers.New(DB, sessionManager)

	// Setup and initialize routes
	setUpRoutes.New(router)
	setUpRoutes.InitializeRoutes(h, sessionManager)

	// Middleware stacks
	stacks := middleware.MiddlewareStack(
		middleware.LoggerMiddleware,
		sessionManager.LoadAndSave,
	)

	server := &http.Server{
		Addr:    ":8080",
		Handler: stacks(router),
	}

	log.Printf("Server listening on port %v", server.Addr)

	server.ListenAndServe()
}
