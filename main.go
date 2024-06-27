package main

import (
	"clean/architector/internal/adapters/productrepo/postgresql"
	"fmt"
)

type ProductRepoInterface interface {
	GetProductById(productId int)
}

func main() {
	newProductRepo := postgresql.NewProductRepo(1)

	get1(newProductRepo)

}

func get1(p ProductRepoInterface) {
fmt.Println(p.GetProductById(1))
}
