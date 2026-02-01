package service

import (
	"kasir-api/model"
	"kasir-api/repositories"
)

type UserService struct {
	repo *repositories.UserRepository
}

// struct method untuk struct: UserService
func NewUserService(repo *repositories.UserRepository) *UserService {
	return &UserService{repo: repo}
}

// === Services functions ===

func (srvc *UserService) GetAll() ([]model.User, error) {
	return srvc.repo.GetAll()
}

func (srvc *UserService) GetByID(id int) (*model.User, error) {
	return srvc.repo.GetByID(id)
}

func (srvc *UserService) Update(user *model.User) error {
	return srvc.repo.Update(user)
}

func (srvc *UserService) Delete(id int) error {
	return srvc.repo.Delete(id)
}
