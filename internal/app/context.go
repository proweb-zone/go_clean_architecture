package app

type ProductRepo interface {
	getProductById()
}

type Post struct{}

func (p *Post) NewContext() {

}
