package repository

import "clean/architector/internal/domain/entitie"

type ProductRepo struct {
	ProductId int
}

func NewProductRepo(productId int) *ProductRepo {
	return &ProductRepo{ProductId: productId}
}

func (p *ProductRepo) GetProductById(productId int) entitie.Employee {
	employees := map[string]entitie.Employee{
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

	var newEmploye = entitie.Employee{}

	for _, employee := range employees {

		if productId == employee.Id {
			newEmploye = employee
			break
		}

	}

	return newEmploye

}
