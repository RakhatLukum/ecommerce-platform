package repository

import (
	"order-service/model"
	"sync"
)

type MemoryRepo struct {
	mu     sync.Mutex
	orders map[string]model.Order
}

func NewMemoryRepo() *MemoryRepo {
	return &MemoryRepo{
		orders: make(map[string]model.Order),
	}
}

func (r *MemoryRepo) Create(o model.Order) {
	r.mu.Lock()
	defer r.mu.Unlock()
	r.orders[o.ID] = o
}

func (r *MemoryRepo) Get(id string) (model.Order, bool) {
	r.mu.Lock()
	defer r.mu.Unlock()
	o, ok := r.orders[id]
	return o, ok
}

func (r *MemoryRepo) Update(id string, o model.Order) {
	r.mu.Lock()
	defer r.mu.Unlock()
	r.orders[id] = o
}

func (r *MemoryRepo) List() []model.Order {
	r.mu.Lock()
	defer r.mu.Unlock()
	var list []model.Order
	for _, o := range r.orders {
		list = append(list, o)
	}
	return list
}
