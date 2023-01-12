package service

import (
	"github.com/kyzykyky/softwarearch/bookservice/internal/data/bookrepository"
	"github.com/kyzykyky/softwarearch/bookservice/internal/domain"
)

// Initial configuration
type Service struct {
	BookDAO bookrepository.BookRepository
}

// Logic owner
type service struct {
	Domain  domain.IDomain
	BookDAO bookrepository.BookRepository
}

func NewService(serv Service) service {
	return service{
		BookDAO: serv.BookDAO,
		Domain: domain.Domain{
			BookDAO: serv.BookDAO,
		},
	}
}
