package service

import "github.com/kyzykyky/softwarearch/svcreg/pkg/domain"

var Stock map[int]domain.Stock = map[int]domain.Stock{
	1: {ID: 1, Quantity: 100},
	2: {ID: 2, Quantity: 200},
	3: {ID: 3, Quantity: 300},
	4: {ID: 4, Quantity: 400},
	5: {ID: 5, Quantity: 500},
}
