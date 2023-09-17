package service

import (
	"github.com/SamsonAirapetyan/BWG-test"
	"github.com/SamsonAirapetyan/BWG-test/internal/repository"
)

type TransactionService struct {
	repo repository.Transactions
}

func NewTransactionService(repo repository.Transactions) *TransactionService {
	return &TransactionService{repo: repo}
}

func (s *TransactionService) AddSum(user BWG_test.Request) error {
	return s.repo.AddSum(user)
}

func (s *TransactionService) TakeOff(user BWG_test.Request) error {
	return s.repo.TakeOff(user)
}

func (s *TransactionService) GetAll() ([]BWG_test.Answer, error) {
	return s.repo.GetAll()
}
