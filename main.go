package main
import "sort"

import (
	"github.com/gin-gonic/gin"
	"strconv"
	"sync"
)

type Book struct {
	ID              int    `json:"id"`
	Title           string `json:"title"`
	Author          string `json:"author"`
	Genre           string `json:"genre"`
	PublicationDate string `json:"publication_date"` // Using string for simplicity, but consider time.Time
}

var (
	books   = []Book{}
	bookID  = 1
	bookMux sync.Mutex
)

func main() {
	r := gin.Default()

	// CRUD for books
	r.POST("/books", createBook)
	r.GET("/books", listBooks)
	r.GET("/books/:id", getBook)
	r.PUT("/books/:id", updateBook)
	r.DELETE("/books/:id", deleteBook)

	// Additional endpoints
	r.GET("/books/author/:name", getBooksByAuthor)
	r.GET("/books/search/:query", searchBooks)
	r.GET("/books/genre/:genre", filterByGenre)
	r.GET("/books/sort/date", sortByDate)
	r.GET("/books/sort/title", sortByTitle)
	r.GET("/books/year/:year", filterByYear)

	r.Run(":8080")
}

func createBook(c *gin.Context) {
	var b Book
	if err := c.BindJSON(&b); err != nil {
		c.JSON(400, gin.H{"error": "Bad request"})
		return
	}

	bookMux.Lock()
	b.ID = bookID
	bookID++
	books = append(books, b)
	bookMux.Unlock()

	c.JSON(201, b)
}

func listBooks(c *gin.Context) {
	c.JSON(200, books)
}

func getBook(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(400, gin.H{"error": "Invalid ID"})
		return
	}

	for _, b := range books {
		if b.ID == id {
			c.JSON(200, b)
			return
		}
	}

	c.JSON(404, gin.H{"error": "Book not found"})
}

func updateBook(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(400, gin.H{"error": "Invalid ID"})
		return
	}

	var updatedBook Book
	if err := c.BindJSON(&updatedBook); err != nil {
		c.JSON(400, gin.H{"error": "Bad request"})
		return
	}

	bookMux.Lock()
	for i, b := range books {
		if b.ID == id {
			books[i] = updatedBook
			books[i].ID = id
			bookMux.Unlock()
			c.JSON(200, books[i])
			return
		}
	}
	bookMux.Unlock()

	c.JSON(404, gin.H{"error": "Book not found"})
}

func deleteBook(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(400, gin.H{"error": "Invalid ID"})
		return
	}

	bookMux.Lock()
	for i, b := range books {
		if b.ID == id {
			books = append(books[:i], books[i+1:]...)
			bookMux.Unlock()
			c.JSON(200, gin.H{"message": "Book deleted"})
			return
		}
	}
	bookMux.Unlock()

	c.JSON(404, gin.H{"error": "Book not found"})
}

func getBooksByAuthor(c *gin.Context) {
	author := c.Param("name")
	var authorBooks []Book

	for _, b := range books {
		if b.Author == author {
			authorBooks = append(authorBooks, b)
		}
	}

	c.JSON(200, authorBooks)
}

func searchBooks(c *gin.Context) {
	query := c.Param("query")
	var foundBooks []Book

	for _, b := range books {
		if contains(b.Title, query) || contains(b.Author, query) {
			foundBooks = append(foundBooks, b)
		}
	}

	c.JSON(200, foundBooks)
}

func filterByGenre(c *gin.Context) {
	genre := c.Param("genre")
	var genreBooks []Book

	for _, b := range books {
		if b.Genre == genre {
			genreBooks = append(genreBooks, b)
		}
	}

	c.JSON(200, genreBooks)
}

func sortByDate(c *gin.Context) {
	sortedBooks := make([]Book, len(books))
	copy(sortedBooks, books)

	sort.Slice(sortedBooks, func(i, j int) bool {
		return sortedBooks[i].PublicationDate < sortedBooks[j].PublicationDate
	})

	c.JSON(200, sortedBooks)
}

func sortByTitle(c *gin.Context) {
	sortedBooks := make([]Book, len(books))
	copy(sortedBooks, books)

	sort.Slice(sortedBooks, func(i, j int) bool {
		return sortedBooks[i].Title < sortedBooks[j].Title
	})

	c.JSON(200, sortedBooks)
}

func filterByYear(c *gin.Context) {
	year := c.Param("year")
	var yearBooks []Book

	for _, b := range books {
		if b.PublicationDate[:4] == year { // Assuming YYYY-MM-DD format
			yearBooks = append(yearBooks, b)
		}
	}

	c.JSON(200, yearBooks)
}

func contains(source, target string) bool {
	return strconv.Contains(source, target)
}
