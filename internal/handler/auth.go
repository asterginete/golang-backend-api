package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/asterginete/golang-backend-api/internal/model"
	"github.com/asterginete/golang-backend-api/internal/repository"
	"github.com/asterginete/golang-backend-api/internal/util"
)

// Login handles user authentication and returns a JWT token.
func Login(c *gin.Context) {
	var credentials struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}

	if err := c.BindJSON(&credentials); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Bad request"})
		return
	}

	user, err := repository.GetUserByUsername(credentials.Username)
	if err != nil || !util.CheckPassword(credentials.Password, user.Password) {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid username or password"})
		return
	}

	token, err := util.GenerateToken(user.Username)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": token})
}

// Register handles user registration.
func Register(c *gin.Context) {
	var newUser model.User
	if err := c.BindJSON(&newUser); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Bad request"})
		return
	}

	hashedPassword, err := util.HashPassword(newUser.Password)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
		return
	}
	newUser.Password = hashedPassword

	createdUser, err := repository.AddUser(newUser)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
		return
	}

	c.JSON(http.StatusCreated, createdUser)
}

// ... Potentially add other handlers related to authentication, like password reset, token refresh, etc. ...
