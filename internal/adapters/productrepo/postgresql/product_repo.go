package postgresql

import "clean/architector/internal/core/product"

type ProductRepo struct {
	ProductId int
}

func NewProductRepo(productId int) *ProductRepo {
	return &ProductRepo{ProductId: productId}
}

func (p *ProductRepo) GetProductById(productId int) product.Employee {
	employees := map[string]product.Employee{
		"Лампа": {
			Id:   1,
			Name: "Лампа",
			Desc: "Desc 1",
		},
		"Кабель": {
			Id:   2,
			Name: "Кабель",
			Desc: "Desc 2",
		},
		"Садовая тележка": {
			Id:   3,
			Name: "Садовая тележка",
			Desc: "Desc 3",
		},
		"Фоторамка": {
			Id:   4,
			Name: "Фоторамка",
			Desc: "Desc 4",
		},
	}

	var newEmploye = product.Employee{}

	for _, employee := range employees {

		if productId == employee.Id {
			newEmploye = employee
			break
		}

	}

	return newEmploye

}
