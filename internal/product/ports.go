package product

type Employee struct {
	Id   int
	Name string
	Desc string
}

type ProductRepoInterface interface {
	GetProductById(productId int) Employee
}
