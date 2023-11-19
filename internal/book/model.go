package book

import "github.com/xlzpm/internal/author"

type Book struct {
	ID      string          `json:"id"`
	Name    string          `json:"name"`
	Age     int             `json:"age"`
	Authors []author.Author `json:"authors"`
}
