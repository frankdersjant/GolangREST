package CustomerServices

import (
	"example/rest/Domain"
)

type CustomerService interface {
	GetAllCustomer() ([]Domain.Customer, error)
	FindCustomerById(id string) (*Domain.Customer, *Domain.AppError)
}

type DefaultCustomerService struct {
	repo Domain.CustomerRepository
}

// receiver function -attaches it as a method to a class
func (s DefaultCustomerService) GetAllCustomer() ([]Domain.Customer, error) {
	return s.repo.FindAll()
}

func (s DefaultCustomerService) FindCustomerById(id string) (*Domain.Customer, *Domain.AppError) {
	return s.repo.FindById(id)
}
