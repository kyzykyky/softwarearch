package gorm_bookrepository

import (
	"github.com/kyzykyky/softwarearch/bookservice/pkg/domain"
	"gorm.io/gorm"
)

type Book struct {
	gorm.Model
	Id     uint   `gorm:"primary_key"`
	Isbn   string `gorm:"type:varchar(100)"`
	Title  string
	Author string
	Price  float32
}

func bookToEntity(book domain.Book) Book {
	return Book{
		Id:     uint(book.Id),
		Isbn:   book.Isbn,
		Title:  book.Title,
		Author: book.Author,
		Price:  book.Price,
	}
}

func bookToModel(book Book) domain.Book {
	return domain.Book{
		Id:     int(book.Id),
		Isbn:   book.Isbn,
		Title:  book.Title,
		Author: book.Author,
		Price:  book.Price,
	}
}
