package CustomerServices

import (
	"example/rest/Domain"
)

type CustomerService interface {
	GetAll() []Domain.Customer
}
