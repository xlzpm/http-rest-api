package db_book

import (
	"database/sql"

	"github.com/xlzpm/internal/author/model"
	"github.com/xlzpm/internal/book"
)

type Book struct {
	ID      string         `json:"id"`
	Name    string         `json:"name"`
	Age     sql.NullInt32  `json:"age"`
	Authors []model.Author `json:"authors"`
}

func (m *Book) ToDomain() book.Book {
	b := book.Book{
		ID:   m.ID,
		Name: m.Name,
	}

	if m.Age.Valid {
		b.Age = int(m.Age.Int32)
	}

	return b
}
