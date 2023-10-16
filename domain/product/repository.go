package product

import (
	"errors"

	"github.com/aletomasella/ddd-go/entities"
)

var (
	ErrorProductNotFound = errors.New("product not found")
	ErrorFailedToSave    = errors.New("failed to save product")
	ErrorFailedToUpdate  = errors.New("failed to update product")
	ErrorFailedToDelete  = errors.New("failed to delete product")
)

type ProductRepository interface {
	FindAll() ([]entities.Product, error)
	FindById(id int) (entities.Product, error)
	Save(product entities.Product) error
	Update(product entities.Product) error
	Delete(product entities.Product) error
}
