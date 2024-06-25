package app

import "fmt"

type ProductRepo interface {
	getProductById()
}

type Post struct {}

func (p *Post) NewContext() {
	
}
