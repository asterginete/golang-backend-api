package middleware

import (
	"net/http"
	"strings"

	"github.com/asterginete/golang-backend-apiinternal/util"
	"github.com/gin-gonic/gin"
)

// Authenticate is a middleware function to check for a valid JWT token.
func Authenticate() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Get token from the 'Authorization' header
		tokenString := c.GetHeader("Authorization")

		// Check if the token is present
		if tokenString == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization header not provided"})
			c.Abort()
			return
		}

		// Typically, Authorization header might be in format: Bearer <token>
		// So, we split by space and get the second part
		splitted := strings.Split(tokenString, " ")
		if len(splitted) != 2 {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid authorization header format"})
			c.Abort()
			return
		}
		tokenString = splitted[1]

		// Validate the token
		username, err := util.ValidateToken(tokenString)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
			c.Abort()
			return
		}

		// Set the username in the context so that subsequent handlers can use it
		c.Set("username", username)

		c.Next()
	}
}
