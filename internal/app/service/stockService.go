package service

import "github.com/cjmacharia/portfolioPilot/internal/domain"

type StockService interface {
	PostStock(stock *domain.Stock) (*domain.Stock, error)
}
type DefaultCustomerService struct {
	repo domain.StockRepository
}

func (s DefaultCustomerService) PostStock(stock *domain.Stock) (*domain.Stock, error) {
	savedStock, err := s.repo.AddStock(stock)
	if err != nil {
		return nil, err
	}
	return savedStock, nil
}

func NewStockService(repository domain.StockRepository) DefaultCustomerService {
	return DefaultCustomerService{repository}
}
