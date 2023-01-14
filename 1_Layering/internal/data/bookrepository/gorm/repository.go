package gorm_bookrepository

import (
	"context"

	"github.com/kyzykyky/softwarearch/bookservice/internal/data/config/gorm/sqlite"
	"github.com/kyzykyky/softwarearch/bookservice/pkg/domain"
)

func (repo bookrepo) GetBook(ctx context.Context, id int) (domain.Book, error) {
	var book Book
	if err := repo.db.WithContext(ctx).First(&book, id).Error; err != nil {
		return domain.Book{}, sqlite.ConvertError(err)
	}
	return bookToModel(book), nil
}

func (repo bookrepo) GetBooks(ctx context.Context, count, offset int) ([]domain.Book, error) {
	var books []Book
	if err := repo.db.WithContext(ctx).Limit(count).Offset(offset).Find(&books).Error; err != nil {
		return nil, sqlite.ConvertError(err)
	}
	mbooks := make([]domain.Book, len(books))
	for i, book := range books {
		mbooks[i] = bookToModel(book)
	}
	return mbooks, nil
}

func (repo bookrepo) AddBook(ctx context.Context, book domain.Book) (domain.Book, error) {
	entity := bookToEntity(book)

	tx := repo.db.WithContext(ctx).Begin()
	if err := tx.Create(&entity).Error; err != nil {
		tx.Rollback()
		return domain.Book{}, sqlite.ConvertError(err)
	}
	if tx.Commit().Error != nil {
		return domain.Book{}, sqlite.ConvertError(tx.Commit().Error)
	}
	return bookToModel(entity), nil
}

func (repo bookrepo) UpdateBook(ctx context.Context, book domain.Book) (domain.Book, error) {
	entity := bookToEntity(book)

	tx := repo.db.WithContext(ctx).Begin()
	if err := tx.Model(&Book{}).Where(book.Id).Updates(entity).Error; err != nil {
		tx.Rollback()
		return domain.Book{}, sqlite.ConvertError(err)
	}
	if tx.Commit().Error != nil {
		return domain.Book{}, sqlite.ConvertError(tx.Commit().Error)
	}
	return bookToModel(entity), nil
}

func (repo bookrepo) DeleteBook(ctx context.Context, id int) error {
	tx := repo.db.WithContext(ctx).Begin()

	// var book Book
	// if err := tx.First(&book, "isbn = ?", isbn).Error; err != nil {
	// 	return sqlite.ConvertError(err)
	// }

	if err := tx.Where(id).Delete(&Book{}).Error; err != nil {
		tx.Rollback()
		return sqlite.ConvertError(err)
	}
	if tx.Commit().Error != nil {
		return sqlite.ConvertError(tx.Commit().Error)
	}
	return nil
}
