package adapter

import "clean/architector/internal/domain/entitie"

// interface for useCase and repository
type IproductRepo interface {
	GetProductById(productId int) entitie.Employee
}
