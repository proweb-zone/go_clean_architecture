package product

import "fmt"

type ViewProductUseCase struct {
	productRepo ProductRepo
}

func NewViewProductUseCase(productRepo ProductRepo) *ViewProductUseCase {
return &ViewProductUseCase{productRepo: productRepo}
}

func (u *ViewProductUseCase) Run(productId int) {
product, err := u.productRepo.getProductById(productId)

if err != nil {
fmt.Println("есть ошибки получения данных")
}

fmt.Println(product)

}
