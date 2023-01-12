package sqlite_bookrepository

import (
	"github.com/kyzykyky/softwarearch/bookservice/internal/data/config/sqlite"
	"github.com/kyzykyky/softwarearch/bookservice/pkg/domain"
)

func (repo bookrepo) GetBook(isbn string) (domain.Book, error) {
	var book Book
	if err := repo.db.First(&book, "isbn = ?", isbn).Error; err != nil {
		return domain.Book{}, sqlite.ConvertError(err)
	}
	return bookToModel(book), nil
}

func (repo bookrepo) GetBooks() ([]domain.Book, error) {
	var books []Book
	if err := repo.db.Find(&books).Error; err != nil {
		return nil, sqlite.ConvertError(err)
	}
	mbooks := make([]domain.Book, len(books))
	for i, book := range books {
		mbooks[i] = bookToModel(book)
	}
	return mbooks, nil
}

func (repo bookrepo) AddBook(book domain.Book) (domain.Book, error) {
	entity := bookToEntity(book)

	tx := repo.db.Begin()
	if err := tx.Create(&entity).Error; err != nil {
		tx.Rollback()
		return domain.Book{}, sqlite.ConvertError(err)
	}
	return bookToModel(entity), sqlite.ConvertError(tx.Commit().Error)
}

func (repo bookrepo) UpdateBook(isbn string, book domain.Book) (domain.Book, error) {
	entity := bookToEntity(book)

	tx := repo.db.Begin()
	if err := tx.Model(&Book{}).Where("isbn = ?", isbn).Updates(entity).Error; err != nil {
		tx.Rollback()
		return domain.Book{}, sqlite.ConvertError(err)
	}
	return bookToModel(entity), sqlite.ConvertError(tx.Commit().Error)
}

func (repo bookrepo) DeleteBook(isbn string) error {
	tx := repo.db.Begin()

	// var book Book
	// if err := tx.First(&book, "isbn = ?", isbn).Error; err != nil {
	// 	return sqlite.ConvertError(err)
	// }

	if err := tx.Where("isbn = ?", isbn).Delete(&Book{}).Error; err != nil {
		tx.Rollback()
		return sqlite.ConvertError(err)
	}
	return sqlite.ConvertError(tx.Commit().Error)
}
