package domain

import (
	"context"
)

type BookService interface {
	GetBook(ctx context.Context, id int) (Book, error)
	GetBooks(ctx context.Context, count, offset int) ([]Book, error)
	CreateBook(ctx context.Context, book Book) (Book, error)
	UpdateBook(ctx context.Context, book Book) (Book, error)
	DeleteBook(ctx context.Context, id int) error
}
