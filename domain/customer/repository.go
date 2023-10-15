package customer

import "errors"

var (
	ErrorCustomerNotFound = errors.New("customer not found")
	ErrorFailedToSave     = errors.New("failed to save customer")
	ErrorFailedToUpdate   = errors.New("failed to update customer")
	ErrorFailedToDelete   = errors.New("failed to delete customer")
)

type CustomerRepository interface {
	FindById(id int) (Customer, error)
	Save(customer Customer) error
	Update(customer Customer) error
	Delete(customer Customer) error
}
