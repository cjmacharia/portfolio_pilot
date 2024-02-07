package service

import (
	"github.com/cjmacharia/portfolioPilot/internal/domain"
	"github.com/cjmacharia/portfolioPilot/internal/dto"
)

type StockService interface {
	PostStock(stock *domain.Stock) (*domain.Stock, error)
	GetStockByID(id int) (*domain.Stock, error)
	GetAllStocks() ([]dto.StockResponse, error)
	UpdateStockPrice(symbol string, price float64) error
}
type DefaultStockService struct {
	repo domain.StockRepository
}

func (s DefaultStockService) GetAllStocks() ([]dto.StockResponse, error) {
	stocks, err := s.repo.GetStocks()
	if err != nil {
		return nil, err
	}
	var stockResponses []dto.StockResponse
	for _, stock := range stocks {
		stockResponse := dto.StockResponse{
			Name:     stock.Name,
			Symbol:   stock.Symbol,
			Exchange: stock.Exchange,
			Price:    stock.Price,
		}
		stockResponses = append(stockResponses, stockResponse)
	}
	return stockResponses, nil
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

func (s DefaultStockService) UpdateStockPrice(symbol string, price float64) error {
	err := s.repo.UpdateStock(symbol, price)
	if err != nil {
		return err
	}
	return nil
}

func NewStockService(repository domain.StockRepository) DefaultStockService {
	return DefaultStockService{repository}
}
