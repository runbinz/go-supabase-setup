// cmd/server/main.go

// Package main is the entry point for the server application.
// It initializes the database, sets up middleware, and starts the server.
// This is where the application begins execution and sets up the necessary components for handling requests.

package main

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/runbinz/dashboard/handlers"
	"github.com/runbinz/dashboard/middleware"
	"github.com/runbinz/dashboard/supabase"
)

// init function loads environment variables from a .env file.
// It is automatically called before the main function.
// This is useful for setting up configuration values that the application will use.

func init() {
	_ = godotenv.Load() // Loads .env file into os.Getenv
}

// main function is the entry point of the application.
// It initializes the Supabase database and sets up the Gin router.
// The server listens on a specified port and handles incoming requests.
// This function orchestrates the setup of the server and its routes, making it ready to handle client requests.

func main() {
	// Initialize Database
	if err := supabase.Init(); err != nil {
		log.Fatalf("failed to init Supabase DB: %v", err)
	}

	// Load JWT auth middlware
	authMiddleware := middleware.AuthMiddleware()

	// Start router
	router := gin.Default()
	router.Use(gin.Logger(), gin.Recovery())

	// Public test route
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "pong"})
	})

	// Authenticated routes
	api := router.Group("/api")
	api.Use(authMiddleware)

	api.GET("/get-holdings", handlers.GetHoldings)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	log.Printf("Server running on http://localhost:%s", port)
	router.Run(":" + port)
}

// router.GET defines a public route that responds with "pong" to a GET request.
// This is a simple health check endpoint to verify that the server is running.

// api.GET defines an authenticated route that requires a valid JWT.
// It uses the GetHoldings handler to fetch user holdings.
// This route is protected by the AuthMiddleware, ensuring that only authenticated users can access it.
