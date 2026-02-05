package service

import (
	"kasir-api/model"
	"kasir-api/repositories"
)

type TransactionService struct {
	repo *repositories.TransactionRepository
}

func NewTransactionService(repo *repositories.TransactionRepository) *TransactionService {
	return &TransactionService{repo: repo}
}

func (s *TransactionService) Checkout(items []model.CheckoutItem) (*model.Transaction, error) {
	return s.repo.CreateTransaction(items)
}

func (s *TransactionService) GetTodaySummary() (*model.SalesSummary, error) {
	return s.repo.GetTodaySummary()
}
