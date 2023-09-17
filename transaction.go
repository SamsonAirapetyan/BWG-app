package BWG_test

type Request struct {
	Currency   string  `json:"currency" db:"currency"`
	Sum        float64 `json:"sum" db:"sum" binding:"required"`
	Wallet_num uint64  `json:"wallet_num" db:"wallet_num"`
}

type Answer struct {
	Wallet_num uint64  `json:"wallet_num"`
	Usdt       float64 `json:"usdt"`
	Rub        float64 `json:"rus"`
	Eur        float64 `json:"eur"`
}
