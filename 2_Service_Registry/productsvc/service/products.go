package service

import (
	"fmt"

	"github.com/kyzykyky/softwarearch/svcreg/pkg/domain"
)

var Products map[int]domain.Product = map[int]domain.Product{
	1: {ID: 1, Name: "Butter", Price: 100},
	2: {ID: 2, Name: "Milk", Price: 200},
	3: {ID: 3, Name: "Bread", Price: 300},
	4: {ID: 4, Name: "Cheese", Price: 400},
	5: {ID: 5, Name: "Eggs", Price: 500},
}

var ErrProductNotFound = fmt.Errorf("product not found")

func (s *Service) GetProduct(id int) (domain.Product, error) {
	product, ok := Products[id]
	if !ok {
		return domain.Product{}, ErrProductNotFound
	}
	stock, err := s.GetStock(id)
	if err != nil {
		return domain.Product{}, err
	}
	product.Available = stock.Quantity > 0
	return product, nil
}
