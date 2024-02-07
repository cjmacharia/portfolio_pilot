package service

import (
	"github.com/cjmacharia/portfolioPilot/internal/domain"
)

type Transactions interface {
	NewTransaction(transaction *domain.Transaction) (*domain.Transaction, error)
	GetStockQuantityService(userId int, stockId int) (int64, error)
	GetUserStockService(userId int) ([]*domain.StockPortfolio, error)
	GetUserTransactions(userId int) ([]domain.UserTransactionsResponse, error)
}
type DefaultTransactionService struct {
	repo domain.TransactionRepository
}

func (ts DefaultTransactionService) NewTransaction(t *domain.Transaction) (*domain.Transaction, error) {
	createdTransaction, err := ts.repo.CreateTransaction(t)
	if err != nil {
		return nil, err
	}
	return createdTransaction, nil
}

func (ts DefaultTransactionService) GetUserTransactions(userId int) ([]domain.UserTransactionsResponse, error) {
	transactions, err := ts.repo.GetUserTransactions(userId)
	if err != nil {
		return nil, err
	}
	return transactions, nil
}

func (ts DefaultTransactionService) GetStockQuantityService(userId int, stockId int) (int64, error) {
	amount, err := ts.repo.GetStockQuantity(userId, stockId)
	if err != nil {
		return 0, err
	}
	return amount, nil
}

func (ts DefaultTransactionService) GetUserStockService(userId int) ([]*domain.StockPortfolio, error) {
	stock, err := ts.repo.GetUserStocks(userId)
	if err != nil {
		return nil, err
	}
	return stock, nil
}

func NewTransactionService(repository domain.TransactionRepository) DefaultTransactionService {
	return DefaultTransactionService{repository}
}
