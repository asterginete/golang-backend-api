package middleware

import (
	"fmt"
	"time"

	"github.com/asterginete/golang-backend-apiinternal/repository"
	"github.com/gin-gonic/gin"
)

// Logging is a middleware function to log each request.
func Logging() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Start timer
		startTime := time.Now()

		// Process request
		c.Next()

		// Calculate request duration
		duration := time.Since(startTime)

		// Get status code and user (if available)
		statusCode := c.Writer.Status()
		username, _ := c.Get("username") // This assumes the Authenticate middleware sets the username

		// Log the request details
		logEntry := fmt.Sprintf("METHOD: %s | PATH: %s | STATUS: %d | DURATION: %s", c.Request.Method, c.Request.URL.Path, statusCode, duration)
		if username != nil {
			logEntry += fmt.Sprintf(" | USER: %s", username)
		}

		// Save the log to the database (or any other logging system)
		// For this example, we'll assume a simple function in the repository package
		repository.AddLog(logEntry)

		// Additionally, you can print the log to the console or write to a file
		fmt.Println(logEntry)
	}
}
