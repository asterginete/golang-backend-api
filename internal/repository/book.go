package repository

import (
	"errors"
	"sync"
	"golang-backend-api/internal/model"
)

var (
	books   = []model.Book{}
	bookID  = 1
	bookMux sync.Mutex
)

func AddBook(b model.Book) (model.Book, error) {
	bookMux.Lock()
	defer bookMux.Unlock()

	b.ID = bookID
	bookID++
	books = append(books, b)

	return b, nil
}

func GetBooks() []model.Book {
	return books
}
