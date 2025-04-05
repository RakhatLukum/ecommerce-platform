package service

import (
	"order-service/model"
	"order-service/repository"
)

type OrderService struct {
	repo *repository.MemoryRepo
}

func NewOrderService(repo *repository.MemoryRepo) *OrderService {
	return &OrderService{repo}
}

func (s *OrderService) Create(o model.Order) {
	s.repo.Create(o)
}

func (s *OrderService) Get(id string) (model.Order, bool) {
	return s.repo.Get(id)
}

func (s *OrderService) Update(id string, o model.Order) {
	s.repo.Update(id, o)
}

func (s *OrderService) List() []model.Order {
	return s.repo.List()
}
