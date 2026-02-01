package service

import (
	"errors"
	"kasir-api/model"
	"kasir-api/repositories"
	"kasir-api/utils"
)

type AuthService struct {
	repo   *repositories.AuthRepository
	secret string
}

func NewAuthService(repo *repositories.AuthRepository, secret string) *AuthService {
	return &AuthService{repo: repo, secret: secret}
}

func (srvc *AuthService) Register(name, email, password string) (*model.User, error) {
	hashed, err := utils.HashPassword(password)
	if err != nil {
		return nil, err
	}

	user := model.User{
		Name:     name,
		Email:    email,
		Password: hashed,
	}

	err = srvc.repo.CreateUser(&user)
	return &user, err
}

func (srvc *AuthService) Login(email, password string) (*model.User, string, error) {
	user, err := srvc.repo.GetUserByEmail(email)
	if err != nil {
		return nil, "", errors.New("invalid credentials")
	}

	err = utils.CheckPassword(user.Password, password)
	if err != nil {
		return nil, "", errors.New("invalid credentials")
	}

	token, err := utils.GenerateToken(user.ID, srvc.secret)
	if err != nil {
		return nil, "", err
	}

	return user, token, nil
}
