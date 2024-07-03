package main

import (
	"clean/architector/internal/app"
	"clean/architector/internal/domain/repository"
	"clean/architector/internal/domain/usecase"
)

type IproductUseCase interface {
	Run(productId int)
}

func main() {

	var c *app.Context = app.NewContext()
	c.StartWebServer()

	var productId int = 2
	var iProductUseCase IproductUseCase = usecase.NewViewProductUseCase(repository.NewProductRepo(productId))
	iProductUseCase.Run(productId)
}
