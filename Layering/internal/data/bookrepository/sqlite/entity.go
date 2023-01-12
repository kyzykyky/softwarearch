package sqlite_bookrepository

import (
	"github.com/kyzykyky/softwarearch/bookservice/pkg/domain"
	"gorm.io/gorm"
)

type Book struct {
	gorm.Model
	Id     uint   `gorm:"primary_key"`
	Isbn   string `gorm:"type:varchar(100);unique"`
	Title  string
	Author string
	Price  float32
}

func bookToEntity(book domain.Book) Book {
	return Book{
		Isbn:   book.Isbn,
		Title:  book.Title,
		Author: book.Author,
		Price:  book.Price,
	}
}

func bookToModel(book Book) domain.Book {
	return domain.Book{
		Isbn:   book.Isbn,
		Title:  book.Title,
		Author: book.Author,
		Price:  book.Price,
	}
}
