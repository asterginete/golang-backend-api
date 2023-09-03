package repository

import (
	"github.com/asterginete/golang-backend-api/internal/model"
	"github.com/jinzhu/gorm"
)

// Assuming a global DB connection for simplicity. In a real-world scenario, you'd inject the DB connection.
var db *gorm.DB

// AddBook inserts a new book into the database.
func AddBook(book model.Book) (model.Book, error) {
	if err := db.Create(&book).Error; err != nil {
		return model.Book{}, err
	}
	return book, nil
}

// GetBooks retrieves all books from the database.
func GetBooks() []model.Book {
	var books []model.Book
	db.Find(&books)
	return books
}

// GetBookByID retrieves a single book by its ID.
func GetBookByID(id int) (model.Book, error) {
	var book model.Book
	if err := db.First(&book, id).Error; err != nil {
		return model.Book{}, err
	}
	return book, nil
}

// UpdateBook updates a book's details in the database.
func UpdateBook(id int, updatedBook model.Book) (model.Book, error) {
	var book model.Book
	if err := db.First(&book, id).Error; err != nil {
		return model.Book{}, err
	}

	// Update fields
	book.Title = updatedBook.Title
	book.Author = updatedBook.Author
	book.Genre = updatedBook.Genre
	book.PublicationDate = updatedBook.PublicationDate
	book.CoverImageURL = updatedBook.CoverImageURL

	if err := db.Save(&book).Error; err != nil {
		return model.Book{}, err
	}
	return book, nil
}

// DeleteBook removes a book from the database.
func DeleteBook(id int) error {
	var book model.Book
	if err := db.First(&book, id).Error; err != nil {
		return err
	}
	return db.Delete(&book).Error
}

// ... Potentially add other database operations related to books, like filtering, sorting, etc. ...
