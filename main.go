package main

import (
	"clean/architector/internal/adapters/productrepo/postgresql"
	"clean/architector/internal/core/product"
)

type IproductUseCase interface {
	Run(productId int)
}

func main() {
	var productId int = 2
	var iProductUseCase IproductUseCase = product.NewViewProductUseCase(postgresql.NewProductRepo(productId))
	iProductUseCase.Run(productId)
}
