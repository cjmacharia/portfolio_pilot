package domain

type TransactionType string

const (
	Deposit TransactionType = "DEPOSIT"
	Buy     TransactionType = "BUY"
	Sell    TransactionType = "SELL"
)

type Transaction struct {
	TransactionID   int             `json:"_"`
	UserID          int             `json:"user_id"`
	StockID         interface{}     `json:"_"`
	TransactionType TransactionType `json:"transaction_type"`
	Quantity        int64           `json:"quantity"`
	TotalAmount     float64         `json:"total_amount"`
}

type TransactionRepository interface {
	CreateTransaction(transaction *Transaction) (*Transaction, error)
	GetStockQuantity(userId int, stockId int) (int64, error)
	GetUserStocks(userID int) ([]*StockPortfolio, error)
}
