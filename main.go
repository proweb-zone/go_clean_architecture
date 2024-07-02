package main

import (
	"clean/architector/internal/adapters/productrepo/postgresql"
	"clean/architector/internal/app"
	"clean/architector/internal/core/product"
)

type IproductUseCase interface {
	Run(productId int)
}

func main() {

var c *app.Context = app.NewContext()
c.StartWebServer()

	var productId int = 2
	var iProductUseCase IproductUseCase = product.NewViewProductUseCase(postgresql.NewProductRepo(productId))
	iProductUseCase.Run(productId)
}
