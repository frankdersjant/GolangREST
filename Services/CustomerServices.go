package CustomerServices

import (
	"example/rest/Domain"
)

type CustomerService interface {
	GetAllCustomer() ([]Domain.Customer, error)
	FindCustomerById(id string) *Domain.Customer
}
