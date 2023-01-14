package bookrepository

import (
	"context"

	"github.com/kyzykyky/softwarearch/bookservice/pkg/domain"
)

type Initializer interface {
	Init() (BookRepository, error)
}

type BookRepository interface {
	GetBook(ctx context.Context, id int) (domain.Book, error)
	GetBooks(ctx context.Context, count, offset int) ([]domain.Book, error)
	AddBook(ctx context.Context, book domain.Book) (domain.Book, error)
	UpdateBook(ctx context.Context, book domain.Book) (domain.Book, error)
	DeleteBook(ctx context.Context, id int) error
}
