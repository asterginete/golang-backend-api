package handler

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/asterginete/golang-backend-api/internal/model"
	"github.com/asterginete/golang-backend-api/internal/repository"
	"github.com/asterginete/golang-backend-api/internal/util"
)

// GetUserProfile retrieves the profile of a specific user.
func GetUserProfile(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}

	user, err := repository.GetUserByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	// Hide sensitive information
	user.Password = ""

	c.JSON(http.StatusOK, user)
}

// UpdateUserProfile updates the details of a specific user.
func UpdateUserProfile(c *gin.Context) {
	var updatedUser model.User
	if err := c.BindJSON(&updatedUser); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Bad request"})
		return
	}

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}

	// If the password is being updated, hash the new password
	if updatedUser.Password != "" {
		hashedPassword, err := util.HashPassword(updatedUser.Password)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
			return
		}
		updatedUser.Password = hashedPassword
	}

	user, err := repository.UpdateUser(id, updatedUser)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	// Hide sensitive information
	user.Password = ""

	c.JSON(http.StatusOK, user)
}

// ListUsers retrieves a list of all users.
func ListUsers(c *gin.Context) {
	users := repository.GetUsers()

	// Hide sensitive information
	for i := range users {
		users[i].Password = ""
	}

	c.JSON(http.StatusOK, users)
}

// ... Potentially add other handlers related to user operations, like deleting a user, etc. ...
