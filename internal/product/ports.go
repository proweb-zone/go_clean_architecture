package product

type ProductRepo interface {
	GetPostByPath(productId string) (Product, error)
}
