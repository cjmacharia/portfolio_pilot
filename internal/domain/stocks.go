package domain

import "time"

type Stock struct {
	StockID     string
	Name        string
	Symbol      string
	Exchange    string
	Price       float64
	LastUpdated time.Time
}

type StockRepository interface {
	AddStock(stock *Stock) (*Stock, error)
}
