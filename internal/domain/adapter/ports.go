package adapter

import "clean/architector/internal/domain/entitie"

type IproductRepo interface {
	GetProductById(productId int) entitie.Employee
}
