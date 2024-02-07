package domain

import (
	"time"
)

type TransactionType string

const (
	Deposit  TransactionType = "DEPOSIT"
	Buy      TransactionType = "BUY"
	Sell     TransactionType = "SELL"
	Withdraw TransactionType = "WITHDRAW"
)

type Transaction struct {
	TransactionID   int             `json:"_"`
	UserID          int             `json:"user_id"`
	StockID         interface{}     `json:"_"`
	TransactionType TransactionType `json:"transaction_type"`
	Quantity        int64           `json:"quantity"`
	TotalAmount     float64         `json:"total_amount"`
}
type UserTransactionsResponse struct {
	TransactionType TransactionType `json:"transaction_type"`
	Quantity        int64           `json:"quantity"`
	TotalAmount     float64         `json:"total_amount"`
	StockName       interface{}     `json:"stock_name"`
	TransactionDate time.Time       `json:"transaction_date"`
}
type TransactionRepository interface {
	CreateTransaction(transaction *Transaction) (*Transaction, error)
	GetStockQuantity(userId int, stockId int) (int64, error)
	GetUserStocks(userID int) ([]*StockPortfolio, error)
	GetUserTransactions(userID int) ([]UserTransactionsResponse, error)
}
