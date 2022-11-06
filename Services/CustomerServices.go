package Services

import (
	"example/rest/Domain"
)

type CustomerService interface {
	GetAllCustomer() ([]Domain.Customer, error)
	FindCustomerById(id string) (*Domain.Customer, *Domain.AppError)
}

//"constructor" like function
//whereby we pass in the repo (interface) as a dependency
type DefaultCustomerService struct {
	repo Domain.CustomerRepository
}

// receiver function -attaches it as a method to a class
func (s DefaultCustomerService) GetAllCustomer() ([]Domain.Customer, error) {

	//Once again we talk to the interface
	return s.repo.FindAll()
}

func (s DefaultCustomerService) FindCustomerById(id string) (*Domain.Customer, *Domain.AppError) {

	//Once again we talk to the interface
	return s.repo.FindById(id)
}

// Helper function to instantiate customer service
func NewCustomerService(repo Domain.CustomerRepository) DefaultCustomerService {
	return DefaultCustomerService{repo}
}
