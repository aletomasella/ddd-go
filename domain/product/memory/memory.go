package memory

import (
	"sync"

	"github.com/aletomasella/ddd-go/domain/product"
	"github.com/aletomasella/ddd-go/entities"
)

type MemoryProductRepository struct {
	products map[int]entities.Product
	sync.Mutex
}

func NewMemoryProductRepository() *MemoryProductRepository {
	return &MemoryProductRepository{
		products: make(map[int]entities.Product),
	}
}

func (r *MemoryProductRepository) FindAll() ([]entities.Product, error) {
	r.Lock()
	defer r.Unlock()

	var productList []entities.Product
	for _, product := range r.products {
		productList = append(productList, product)
	}

	return productList, nil
}

func (r *MemoryProductRepository) FindById(id int) (entities.Product, error) {
	r.Lock()
	defer r.Unlock()

	p, ok := r.products[id]
	if !ok {
		return entities.Product{}, product.ErrorProductNotFound
	}

	return p, nil
}

func (r *MemoryProductRepository) Save(p entities.Product) error {
	r.Lock()
	defer r.Unlock()

	if _, ok := r.products[p.Id]; ok {
		return product.ErrorFailedToSave
	}

	r.products[p.Id] = p

	return nil
}

func (r *MemoryProductRepository) Update(p entities.Product) error {
	r.Lock()
	defer r.Unlock()

	if _, ok := r.products[p.Id]; !ok {
		return product.ErrorFailedToUpdate
	}

	r.products[p.Id] = p

	return nil
}

func (r *MemoryProductRepository) Delete(p entities.Product) error {
	r.Lock()
	defer r.Unlock()

	if _, ok := r.products[p.Id]; !ok {
		return product.ErrorFailedToDelete
	}

	delete(r.products, p.Id)

	return nil
}
