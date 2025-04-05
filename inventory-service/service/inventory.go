package service

import (
	"inventory-service/model"
	"inventory-service/repository"
)

type InventoryService struct {
	repo *repository.MemoryRepo
}

func NewInventoryService(repo *repository.MemoryRepo) *InventoryService {
	return &InventoryService{repo}
}

func (s *InventoryService) Create(p model.Product) {
	s.repo.Create(p)
}

func (s *InventoryService) Get(id string) (model.Product, bool) {
	return s.repo.Get(id)
}

func (s *InventoryService) Update(id string, p model.Product) {
	s.repo.Update(id, p)
}

func (s *InventoryService) Delete(id string) {
	s.repo.Delete(id)
}

func (s *InventoryService) List() []model.Product {
	return s.repo.List()
}
