package book

import "context"

type Repository interface {
	Create(ctx context.Context, book *Book) error
	FindAll(ctx context.Context) (b []Book, err error)
	FindOne(ctx context.Context, id string) (Book, error)
	Update(ctx context.Context, author Book) error
	Delete(ctx context.Context, id string) error
}
