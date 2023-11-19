package db_book

import (
	"context"

	"github.com/xlzpm/internal/book"
	"github.com/xlzpm/pkg/client/postgresql"
	"github.com/xlzpm/pkg/logging"
)

type repository struct {
	client postgresql.Client
	logger *logging.Logger
}

func (*repository) Create(ctx context.Context, book *book.Book) error {
	panic("unimplemented")
}

func (*repository) Delete(ctx context.Context, id string) error {
	panic("unimplemented")
}

func (r *repository) FindAll(ctx context.Context) (b []book.Book, err error) {
	q := `SELECT id, name, age FROM book`

	rows, err := r.client.Query(ctx, q)
	if err != nil {
		return nil, err
	}

	books := make([]book.Book, 0)

	for rows.Next() {
		var book Book

		err = rows.Scan(&book.ID, &book.Name, &book.Age)
		if err != nil {
			return nil, err
		}

		books = append(books, book.ToDomain())
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return books, nil
}

func (*repository) FindOne(ctx context.Context, id string) (book.Book, error) {
	panic("unimplemented")
}

func (*repository) Update(ctx context.Context, author book.Book) error {
	panic("unimplemented")
}

func NewRepository(client postgresql.Client, logger *logging.Logger) book.Repository {
	return &repository{
		client: client,
		logger: logger,
	}
}
