package service

import (
	"kasir-api/model"
	"kasir-api/repositories"
)

type ProductService struct {
	repo *repositories.ProductRepository
}

func NewProductService(repo *repositories.ProductRepository) *ProductService {
	return &ProductService{repo: repo}
}

func (s *ProductService) GetAll(name string) ([]model.Product, error) {
	return s.repo.GetAll(name)
}

func (s *ProductService) Create(data *model.Product) error {
	return s.repo.Create(data)
}

func (s *ProductService) GetByID(id int) (*model.Product, error) {
	return s.repo.GetByID(id)
}

func (s *ProductService) GetByIDWithCategory(id int) (*model.ProductWithCategory, error) {
	return s.repo.GetByIDWithCategory(id)
}

func (s *ProductService) Update(product *model.Product) error {
	return s.repo.Update(product)
}

func (s *ProductService) Delete(id int) error {
	return s.repo.Delete(id)
}
