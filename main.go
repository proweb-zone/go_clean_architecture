package main

import (
	"fmt"

	"clean/architector/internal/adapters/productrepo/postgresql"
)

func main() {
	test := postgresql.GetProductById(1)
	fmt.Println(test.Name)
	fmt.Println("main go success")
}
