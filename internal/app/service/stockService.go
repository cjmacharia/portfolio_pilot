package service

import "github.com/cjmacharia/portfolioPilot/internal/domain"

type StockService interface {
	PostStock(stock *domain.Stock) (*domain.Stock, error)
	GetStockByID(id int) (*domain.Stock, error)
}
type DefaultStockService struct {
	repo domain.StockRepository
}

func (s DefaultStockService) PostStock(stock *domain.Stock) (*domain.Stock, error) {
	savedStock, err := s.repo.AddStock(stock)
	if err != nil {
		return nil, err
	}
	return savedStock, nil
}
func (s DefaultStockService) GetStockByID(id int) (*domain.Stock, error) {
	stock, err := s.repo.GetStockByID(id)
	if err != nil {
		return nil, err
	}
	return stock, nil
}

func NewStockService(repository domain.StockRepository) DefaultStockService {
	return DefaultStockService{repository}
}
