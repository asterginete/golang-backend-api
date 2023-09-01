package handler

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/asterginete/golang-backend-api/internal/model"
	"github.com/asterginete/golang-backend-api/internal/repository"
)

// CreateBook handles the creation of a new book.
func CreateBook(c *gin.Context) {
	var b model.Book
	if err := c.BindJSON(&b); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Bad request"})
		return
	}

	createdBook, err := repository.AddBook(b)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
		return
	}

	c.JSON(http.StatusCreated, createdBook)
}

// ListBooks handles the retrieval of all books.
func ListBooks(c *gin.Context) {
	books := repository.GetBooks()
	c.JSON(http.StatusOK, books)
}

// GetBook handles the retrieval of a single book by ID.
func GetBook(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid book ID"})
		return
	}

	book, err := repository.GetBookByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Book not found"})
		return
	}

	c.JSON(http.StatusOK, book)
}

// UpdateBook handles updating a book's details.
func UpdateBook(c *gin.Context) {
	var updatedBook model.Book
	if err := c.BindJSON(&updatedBook); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Bad request"})
		return
	}

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid book ID"})
		return
	}

	book, err := repository.UpdateBook(id, updatedBook)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Book not found"})
		return
	}

	c.JSON(http.StatusOK, book)
}

// DeleteBook handles the deletion of a book.
func DeleteBook(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid book ID"})
		return
	}

	if err := repository.DeleteBook(id); err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Book not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Book deleted successfully"})
}

// ... Add other handlers for filtering, sorting, etc. as per the specifications ...
