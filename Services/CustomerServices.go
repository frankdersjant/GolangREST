package CustomerServices

import (
	"example/rest/Domain"
)

type CustomerService interface {
	GetAllCustomer() ([]Domain.Customer, error)
}
