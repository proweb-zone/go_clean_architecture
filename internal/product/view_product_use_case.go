package product

import (
	"fmt"
)

type ViewProductUseCase struct {
	productRepoInterface ProductRepoInterface
}

func NewViewProductUseCase(productRepoInterface ProductRepoInterface) *ViewProductUseCase {
return &ViewProductUseCase{productRepoInterface: productRepoInterface}
}

func (u ViewProductUseCase) Run(productId int) {
var product = u.productRepoInterface.GetProductById(productId)

fmt.Println(product)

}
