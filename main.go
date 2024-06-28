package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"

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

	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello World!"))
	})
	http.ListenAndServe("localhost:3000", r)
}
