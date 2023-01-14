package service

import (
	"context"

	"github.com/kyzykyky/softwarearch/bookservice/pkg/domain"
)

func (s service) CreateBook(ctx context.Context, book domain.Book) (domain.Book, error) {
	book, err := s.BookDAO.AddBook(ctx, book)
	if err != nil {
		return domain.Book{}, err
	}
	return book, nil
}

func (s service) GetBook(ctx context.Context, id int) (domain.Book, error) {
	book, err := s.BookDAO.GetBook(ctx, id)
	if err != nil {
		return domain.Book{}, err
	}
	return book, nil
}

func (s service) GetBooks(ctx context.Context, count, offset int) ([]domain.Book, error) {
	books, err := s.BookDAO.GetBooks(ctx, count, offset)
	if err != nil {
		return nil, err
	}
	return books, nil
}

func (s service) UpdateBook(ctx context.Context, book domain.Book) (domain.Book, error) {
	book, err := s.BookDAO.UpdateBook(ctx, book)
	if err != nil {
		return domain.Book{}, err
	}
	return book, nil
}

func (s service) DeleteBook(ctx context.Context, id int) error {
	err := s.BookDAO.DeleteBook(ctx, id)
	if err != nil {
		return err
	}
	return nil
}
