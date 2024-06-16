package main

import (
	"context"
	"embed"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gomodule/redigo/redis"
	"github.com/joho/godotenv"
	"github.com/sethvargo/go-limiter"
	"github.com/sethvargo/go-limiter/httplimit"
	"github.com/sethvargo/go-redisstore"

	"github.com/Anuolu-2020/Expense-Calculator-App/db"
	"github.com/Anuolu-2020/Expense-Calculator-App/handlers"
	"github.com/Anuolu-2020/Expense-Calculator-App/middleware"
	"github.com/Anuolu-2020/Expense-Calculator-App/pkg"
	"github.com/Anuolu-2020/Expense-Calculator-App/routes"
)

//go:embed static
var static embed.FS

// Connect to redis for rate limiter store
func ConnectRedis() limiter.Store {
	// ctx := context.Background()

	addr := os.Getenv("REDIS_ADDR")

	store, err := redisstore.New(&redisstore.Config{
		Tokens:   40,
		Interval: time.Minute,
		Dial: func() (redis.Conn, error) {
			return redis.DialURL(addr)
		},
	})
	if err != nil {
		log.Printf("[ERROR]: Error connecting to redis: %v", err)
		log.Fatal(err)
	}
	// defer store.Close(ctx)
	log.Print("Connected to Redis Successfully")

	return store
}

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

	// Connect to redis
	store := ConnectRedis()

	rateLimiterMiddleware, err := httplimit.NewMiddleware(store, httplimit.IPKeyFunc())
	if err != nil {
		log.Fatal(err)
	}

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

	// Rate limit all routes
	wrappedRouter := stacks(rateLimiterMiddleware.Handle(router))

	server := &http.Server{
		Addr:    ":8080",
		Handler: wrappedRouter,
	}

	go func() {
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			fmt.Printf("Could not listen on :8080: %v\n", err)
			log.Fatal(err)
		}
	}()

	log.Printf("Server listening on port %v", server.Addr)

	// Channel to listen for interrupt or terminate signals from the OS
	stopChan := make(chan os.Signal, 1)
	signal.Notify(stopChan, os.Interrupt, syscall.SIGTERM)

	// Block until we receive our signal
	<-stopChan
	log.Println("Shutting down server...")

	// Create a context with a timeout to allow the server to shutdown gracefully
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Shutdown the server
	if err := server.Shutdown(ctx); err != nil {
		log.Fatalf("Server forced to shutdown: %v\n", err)
	} else {
		log.Println("Server stopped gracefully")
	}
}
