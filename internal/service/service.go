package service

import (
	"github.com/SamsonAirapetyan/BWG-test"
	"github.com/SamsonAirapetyan/BWG-test/internal/repository"
)

type Transactions interface {
	AddSum(user BWG_test.Request) error
	TakeOff(user BWG_test.Request) error
	GetAll() ([]BWG_test.Answer, error)
}

type Service struct {
	Transactions
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Transactions: NewTransactionService(repos),
	}
}
