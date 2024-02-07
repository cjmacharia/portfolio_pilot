package domain

import "time"

type Stock struct {
	StockID     int       `json:"_"`
	Name        string    `json:"name"`
	Symbol      string    `json:"symbol"`
	Exchange    string    `json:"exchange"`
	Price       float64   `json:"price"`
	LastUpdated time.Time `json:"_"`
}

type StockPortfolio struct {
	StockID     int `json:"_"`
	NetQuantity int `json:"net_quantity"`
}

type StockRepository interface {
	AddStock(stock *Stock) (*Stock, error)
	GetStockByID(stockID int) (*Stock, error)
	GetStocks() ([]Stock, error)
	UpdateStock(symbol string, price float64) error
}
