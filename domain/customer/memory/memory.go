// In memory implementation of the customer repository.
package memory

import (
	"sync"

	"github.com/aletomasella/ddd-go/domain/customer"
)

type CustomerMemoryRepository struct {
	customers map[int]customer.Customer
	// This is for locking the map
	sync.Mutex
}

func NewCustomerMemoryRepository() *CustomerMemoryRepository {
	return &CustomerMemoryRepository{
		customers: make(map[int]customer.Customer),
	}
}

// Full implementation of the CustomerRepository interface

func (mr *CustomerMemoryRepository) FindById(id int) (customer.Customer, error) {
	// This is a read operation, so we need to lock the map
	mr.Lock()
	defer mr.Unlock()

	if c, ok := mr.customers[id]; ok {
		return c, nil
	}

	return customer.Customer{}, customer.ErrorCustomerNotFound
}

func (mr *CustomerMemoryRepository) Save(c customer.Customer) error {
	// This is a write operation, so we need to lock the map
	mr.Lock()
	defer mr.Unlock()

	// If the customer is already in the map, we return an error
	if _, ok := mr.customers[c.GetId()]; ok {
		return customer.ErrorFailedToSave
	}

	mr.customers[c.GetId()] = c

	return nil
}

func (mr *CustomerMemoryRepository) Update(c customer.Customer) error {
	// This is a write operation, so we need to lock the map
	mr.Lock()
	defer mr.Unlock()

	// Cant update a customer that is not in the map
	if _, ok := mr.customers[c.GetId()]; !ok {
		return customer.ErrorFailedToUpdate
	}

	mr.customers[c.GetId()] = c

	return nil
}

func (mr *CustomerMemoryRepository) Delete(c customer.Customer) error {
	// This is a write operation, so we need to lock the map
	mr.Lock()
	defer mr.Unlock()

	delete(mr.customers, c.GetId())

	return nil
}
