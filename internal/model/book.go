package model

import "time"

// Book represents the structure of a book in our application.
type Book struct {
	ID              int       `json:"id" gorm:"primaryKey;autoIncrement"`
	Title           string    `json:"title" gorm:"type:varchar(255);not null"`
	Author          string    `json:"author" gorm:"type:varchar(255);not null"`
	Genre           string    `json:"genre" gorm:"type:varchar(100)"`
	PublicationDate time.Time `json:"publication_date" gorm:"type:date"`
	CoverImageURL   string    `json:"cover_image_url" gorm:"type:varchar(255)"`
	CreatedBy       int       `json:"created_by" gorm:"type:int;not null"`
	CreatedAt       time.Time `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt       time.Time `json:"updated_at" gorm:"autoUpdateTime"`
}

// BeforeSave is a GORM hook that can be used to set or modify fields before saving the book.
func (b *Book) BeforeSave() error {
	// Any logic you want to run before saving a book
	// For example, you might want to validate the book's data or set default values.
	return nil
}

// AfterSave is a GORM hook that can be used to perform actions after saving the book.
func (b *Book) AfterSave() error {
	// Any logic you want to run after saving a book
	// For example, you might want to log the creation of a new book.
	return nil
}

// ... Potentially add other methods or hooks related to the Book model ...
