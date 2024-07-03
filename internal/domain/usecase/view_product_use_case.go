package usecase

import (
	"clean/architector/internal/domain/adapter"
	"fmt"
)

type ViewProductUseCase struct {
	iProductRepo adapter.IproductRepo
}

func NewViewProductUseCase(iProductRepo adapter.IproductRepo) *ViewProductUseCase {
	return &ViewProductUseCase{iProductRepo: iProductRepo}
}

func (u ViewProductUseCase) Run(productId int) {
	var product = u.iProductRepo.GetProductById(productId)
	fmt.Println(product)
}
