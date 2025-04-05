package repository

import (
	"inventory-service/model"
	"sync"
)

type MemoryRepo struct {
	mu       sync.Mutex
	products map[string]model.Product
}

func NewMemoryRepo() *MemoryRepo {
	return &MemoryRepo{
		products: make(map[string]model.Product),
	}
}

func (r *MemoryRepo) Create(p model.Product) {
	r.mu.Lock()
	defer r.mu.Unlock()
	r.products[p.ID] = p
}

func (r *MemoryRepo) Get(id string) (model.Product, bool) {
	r.mu.Lock()
	defer r.mu.Unlock()
	p, ok := r.products[id]
	return p, ok
}

func (r *MemoryRepo) Update(id string, p model.Product) {
	r.mu.Lock()
	defer r.mu.Unlock()
	r.products[id] = p
}

func (r *MemoryRepo) Delete(id string) {
	r.mu.Lock()
	defer r.mu.Unlock()
	delete(r.products, id)
}

func (r *MemoryRepo) List() []model.Product {
	r.mu.Lock()
	defer r.mu.Unlock()
	var list []model.Product
	for _, p := range r.products {
		list = append(list, p)
	}
	return list
}
