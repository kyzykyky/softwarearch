package service

import "github.com/kyzykyky/softwarearch/bookservice/pkg/domain"

func (s service) AddBook(book domain.Book) (domain.Book, error) {
	book, err := s.BookDAO.AddBook(book)
	if err != nil {
		return domain.Book{}, err
	}
	return book, nil
}

func (s service) GetBook(isbn string) (domain.Book, error) {
	book, err := s.BookDAO.GetBook(isbn)
	if err != nil {
		return domain.Book{}, err
	}
	return book, nil
}

func (s service) GetBooks() ([]domain.Book, error) {
	books, err := s.BookDAO.GetBooks()
	if err != nil {
		return nil, err
	}
	return books, nil
}

func (s service) UpdateBook(isbn string, book domain.Book) (domain.Book, error) {
	book, err := s.BookDAO.UpdateBook(isbn, book)
	if err != nil {
		return domain.Book{}, err
	}
	return book, nil
}

func (s service) DeleteBook(isbn string) error {
	err := s.BookDAO.DeleteBook(isbn)
	if err != nil {
		return err
	}
	return nil
}
