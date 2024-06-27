package product

type ProductRepo interface {
	GetProductById(productId string) (Product, error)
}
