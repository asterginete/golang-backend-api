package model

type Book struct {
	ID              int    `json:"id"`
	Title           string `json:"title"`
	Author          string `json:"author"`
	Genre           string `json:"genre"`
	PublicationDate string `json:"publication_date"` // Using string for simplicity, but consider time.Time
}
