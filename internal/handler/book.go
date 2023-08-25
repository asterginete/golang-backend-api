package handler

import (
	"github.com/gin-gonic/gin"
	"your_project_path/internal/model"
	"your_project_path/internal/repository"
	"strconv"
)

func CreateBook(c *gin.Context) {
	var b model.Book
	if err := c.BindJSON(&b); err != nil {
		c.JSON(400, gin.H{"error": "Bad request"})
		return
	}

	createdBook, err := repository.AddBook(b)
	if err != nil {
		c.JSON(500, gin.H{"error": "Internal Server Error"})
		return
	}

	c.JSON(201, createdBook)
}

func ListBooks(c *gin.Context) {
	books := repository.GetBooks()
	c.JSON(200, books)
}

// ... Add other handlers for GetBook, UpdateBook, DeleteBook, etc.
