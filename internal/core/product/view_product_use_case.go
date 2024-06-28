package product

import (
	"fmt"
)

type ViewProductUseCase struct {
	iProductRepo IproductRepo
}

func NewViewProductUseCase(iProductRepo IproductRepo) *ViewProductUseCase {
	return &ViewProductUseCase{iProductRepo: iProductRepo}
}

func (u ViewProductUseCase) Run(productId int) {
	var product = u.iProductRepo.GetProductById(productId)

	fmt.Println(product)

}
