package main

import (
	"log"
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/asterginete/golang-backend-api/internal/handler"
	"github.com/asterginete/golang-backend-api/internal/middleware"
)

func main() {
	// Initialize Gin
	r := gin.Default()

	// Middleware
	r.Use(cors.Default()) // Handle CORS
	r.Use(middleware.Logging()) // Custom logging middleware
	r.Use(middleware.RateLimiting()) // Custom rate limiting middleware

	// Routes
	r.POST("/login", handler.Login)

	// Protected routes
	authorized := r.Group("/")
	authorized.Use(middleware.Authenticate())
	{
		authorized.POST("/books", handler.CreateBook)
		authorized.GET("/books", handler.ListBooks)
		// ... other book-related endpoints ...

		authorized.GET("/users", handler.ListUsers)
		// ... other user-related endpoints ...
	}

	// Start the server
	port := ":8080"
	log.Printf("Starting server on %s", port)
	if err := r.Run(port); err != nil && err != http.ErrServerClosed {
		log.Fatalf("Failed to start server: %v", err)
	}
}
