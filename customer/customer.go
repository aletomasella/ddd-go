package customer

import (
	"errors"

	"github.com/aletomasella/ddd-go/entities"
)

var (
	ErrorInvalidName = errors.New("invalid user name provided")
)

type Customer struct {
	user         *entities.User
	products     []*entities.Item
	transactions []entities.Transaction
}

// Factory function to create a new Customer
func NewCustomer(name string, lastname string) (Customer, error) {
	if name == "" || lastname == "" {
		return Customer{}, ErrorInvalidName
	}

	user := &entities.User{
		Id:       1,
		Name:     name,
		LastName: lastname,
	}

	return Customer{
		user:         user,
		products:     make([]*entities.Item, 0),
		transactions: make([]entities.Transaction, 0),
	}, nil
}
