package postgresql

type ProductRepo struct {
	ProductId int
}

type Employee struct {
	Id   int
	Name string
	Desc string
}

func NewProductRepo(productId int) *ProductRepo {
	return &ProductRepo{ProductId: productId}
}

func GetProductById(productId int) *Employee {

	employees := map[string]Employee{
		"Лампа": {
			Id:   1,
			Name: "Лампа",
			Desc: "Desc 1",
		},
		"Кабель": {
			Id:   2,
			Name: "Кабель",
			Desc: "Desc 2",
		},
		"Садовая тележка": {
			Id:   3,
			Name: "Садовая тележка",
			Desc: "Desc 3",
		},
		"Фоторамка": {
			Id:   4,
			Name: "Фоторамка",
			Desc: "Desc 4",
		},
	}

	newEmploye := new(Employee)

	for _, employee := range employees {

		if productId == employee.Id {
			(*newEmploye) = employee
			break
		}

	}

	return newEmploye

}
