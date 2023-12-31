package services

import (
	"github.com/aletomasella/ddd-go/domain/customer"
	customerMemory "github.com/aletomasella/ddd-go/domain/customer/memory"
	"github.com/aletomasella/ddd-go/domain/product"
	productMemory "github.com/aletomasella/ddd-go/domain/product/memory"
)

// A service is a type that provides access to the business logic of an application.
// OrderConfiguration is a function that configures a OrderService.

type OrderConfiguration func(os *OrderService) error

type OrderService struct {
	customers customer.CustomerRepository
	products  product.ProductRepository
}

// Factory function that creates a new OrderService. ... means that the function accepts a variable number of arguments.
func NewOrderService(cfgs ...OrderConfiguration) (*OrderService, error) {
	os := &OrderService{}

	// Loop and apply each configuration
	for _, cfg := range cfgs {
		if err := cfg(os); err != nil {
			return nil, err
		}
	}

	return os, nil
}

// WithCustomerRepository is a configuration function that sets the customer repository.

func WithCustomerRepository(cr customer.CustomerRepository) OrderConfiguration {

	// Return a OrderConfiguration that sets the customer repository.
	return func(os *OrderService) error {
		os.customers = cr
		return nil
	}
}

// WithMemoryCustomerRepository is a configuration function that sets the customer repository to an in-memory implementation.
func WithMemoryCustomerRepository() OrderConfiguration {
	return WithCustomerRepository(customerMemory.NewCustomerMemoryRepository())
}

// WithProductRepository is a configuration function that sets the product repository.
func WithProductRepository(pr product.ProductRepository) OrderConfiguration {
	return func(os *OrderService) error {
		os.products = pr
		return nil
	}
}

// WithMemoryProductRepository is a configuration function that sets the product repository to an in-memory implementation.
func WithMemoryProductRepository() OrderConfiguration {
	return WithProductRepository(productMemory.NewMemoryProductRepository())
}
