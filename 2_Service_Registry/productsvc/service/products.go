package service

import "github.com/kyzykyky/softwarearch/svcreg/pkg/domain"

var Products map[int]domain.Product = map[int]domain.Product{
	1: {ID: 1, Name: "Butter", Price: 100},
	2: {ID: 2, Name: "Milk", Price: 200},
	3: {ID: 3, Name: "Bread", Price: 300},
	4: {ID: 4, Name: "Cheese", Price: 400},
	5: {ID: 5, Name: "Eggs", Price: 500},
}
