package repository

import (
	"fmt"
	"github.com/SamsonAirapetyan/BWG-test"
	"github.com/jmoiron/sqlx"
	"log"
)

const (
	Status_neg = "Error"
	Status_pos = "Success"
)

type TransactionPostgres struct {
	db *sqlx.DB
}

func NewTransactionPostgres(db *sqlx.DB) *TransactionPostgres {
	return &TransactionPostgres{db: db}
}

func (r *TransactionPostgres) AddSum(user BWG_test.Request) error {
	var id int
	id, err := r.CreateTransaction(user.Wallet_num, user.Currency, user.Sum)
	Add_money := fmt.Sprintf("UPDATE %s SET %s = %s + %f WHERE wallet_num = %d ", WalletTable, user.Currency, user.Currency, user.Sum, user.Wallet_num)
	_, err = r.db.Exec(Add_money)
	if err != nil {
		r.UpdateStatus(Status_neg, id)
		return err
	}
	r.UpdateStatus(Status_pos, id)
	return err

}

func (r *TransactionPostgres) TakeOff(user BWG_test.Request) error {
	var id int
	id, err := r.CreateTransaction(user.Wallet_num, user.Currency, user.Sum)
	var sum float64
	Get_sum := fmt.Sprintf("SELECT %s FROM %s WHERE wallet_num = %d", user.Currency, WalletTable, user.Wallet_num)
	if err := r.db.Get(&sum, Get_sum); err != nil {
		return err
	}
	if sum-user.Sum < 0 {
		r.UpdateStatus(Status_neg, id)
		return nil
	}
	Take_money := fmt.Sprintf("UPDATE %s SET %s = %s - %f WHERE wallet_num = %d ", WalletTable, user.Currency, user.Currency, user.Sum, user.Wallet_num)
	_, err = r.db.Exec(Take_money)
	if err != nil {
		r.UpdateStatus(Status_neg, id)
		return err
	}
	r.UpdateStatus(Status_pos, id)
	return err
}

func (r *TransactionPostgres) GetAll() ([]BWG_test.Answer, error) {
	var list []BWG_test.Answer

	query := fmt.Sprintf("SELECT t2.wallet_num ,t2.usdt, t2.rub, t2.eur FROM %s t2 JOIN "+
		"(SELECT wallet_num, status, ROW_NUMBER() OVER (PARTITION BY wallet_num ORDER BY id DESC) "+
		"AS rn FROM %s) t1 ON t1.wallet_num = t2.wallet_num WHERE t1.rn = 1 AND t1.status != 'Error'", WalletTable, TransacitonTable)
	err := r.db.Select(&list, query)
	return list, err
}

func (r TransactionPostgres) UpdateStatus(status string, id int) error {
	SetStatus := fmt.Sprintf("UPDATE %s SET status = $1 WHERE id = $2", TransacitonTable)
	_, err := r.db.Exec(SetStatus, status, id)
	return err
}

func (r TransactionPostgres) CreateTransaction(wallet_num uint64, currency string, sum float64) (int, error) {
	var id int
	CreateTransactionQuery := fmt.Sprintf("INSERT INTO %s (wallet_num, currency, sum) VALUES ($1, $2, $3) RETURNING id", TransacitonTable)
	row := r.db.QueryRow(CreateTransactionQuery, wallet_num, currency, sum)
	err := row.Scan(&id)
	if err != nil {
		log.Println(err.Error())
	}
	return id, err
}
