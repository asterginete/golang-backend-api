package handler

import (
	"github.com/gin-gonic/gin"
	"golang-backend-api/internal/util"
)

func Login(c *gin.Context) {
	var credentials struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}

	if err := c.BindJSON(&credentials); err != nil {
		c.JSON(400, gin.H{"error": "Bad request"})
		return
	}

	// For simplicity, we're hardcoding a username and password.
	// In a real-world scenario, you'd fetch and verify this from a database.
	if credentials.Username == "admin" && credentials.Password == "password" {
		token, err := util.GenerateToken(credentials.Username)
		if err != nil {
			c.JSON(500, gin.H{"error": "Internal Server Error"})
			return
		}
		c.JSON(200, gin.H{"token": token})
	} else {
		c.JSON(401, gin.H{"error": "Authentication failed"})
	}
}
