package bookrepository

import (
	"github.com/kyzykyky/softwarearch/bookservice/pkg/domain"
)

type Initializer interface {
	Init() (BookRepository, error)
}

type BookRepository interface {
	GetBook(isbn string) (domain.Book, error)
	GetBooks() ([]domain.Book, error)
	AddBook(book domain.Book) (domain.Book, error)
	UpdateBook(isbn string, book domain.Book) (domain.Book, error)
	DeleteBook(isbn string) error
}
