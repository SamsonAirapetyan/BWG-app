package repository

import (
	"github.com/SamsonAirapetyan/BWG-test"
	"github.com/jmoiron/sqlx"
)

type Transactions interface {
	AddSum(user BWG_test.Request) error
	TakeOff(user BWG_test.Request) error
	GetAll() ([]BWG_test.Answer, error)
}

type Repository struct {
	Transactions
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Transactions: NewTransactionPostgres(db),
	}
}
