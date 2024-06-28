package product

type IproductRepo interface {
	GetProductById(productId int) Employee
}
