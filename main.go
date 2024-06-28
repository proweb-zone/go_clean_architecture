package main

import (
	"clean/architector/internal/adapters/productrepo/postgresql"
	"clean/architector/internal/core/product"
)

type ProductUseCaseInterface interface {
	Run(productId int)
}


func main() {
	var portProductRepo product.ProductRepoInterface = postgresql.NewProductRepo(1)
	var productUseCaseInterface ProductUseCaseInterface = product.NewViewProductUseCase(portProductRepo)
	productUseCaseInterface.Run(1)
}
