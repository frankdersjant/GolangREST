package Domain

type CustomerRepository interface {
	FindAll() ([]Customer, error)
	FindById(id string) (*Customer, *AppError)
}
