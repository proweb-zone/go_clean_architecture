package main

import (
	"clean/architector/internal/adapters/productrepo/postgresql"
	"clean/architector/internal/core/product"
)

type IproductUseCase interface {
	Run(productId int)
}

func main() {
	var iProductUseCase IproductUseCase = product.NewViewProductUseCase(postgresql.NewProductRepo(1))
	iProductUseCase.Run(1)
}
