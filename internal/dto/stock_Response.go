package dto

type StockResponse struct {
	Name     string  `json:"name"`
	Symbol   string  `json:"symbol"`
	Exchange string  `json:"exchange"`
	Price    float64 `json:"price"`
}
