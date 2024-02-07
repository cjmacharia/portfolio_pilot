package database

import (
	"context"
	"github.com/cjmacharia/portfolioPilot/internal/domain"
	"github.com/jmoiron/sqlx"
	"time"
)

func (db RepositoryDB) AddStock(s *domain.Stock) (*domain.Stock, error) {
	sql := "insert into stock (name, price, symbol, exchange)VALUES ($1, $2, $3, $4) RETURNING name, price, symbol, exchange"
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	row := db.Client.QueryRowContext(ctx, sql, s.Name, s.Price, s.Symbol, s.Exchange)
	var savedStock domain.Stock
	err := row.Scan(&savedStock.Name, &savedStock.Price, &savedStock.Symbol, &savedStock.Exchange)
	if err != nil {
		return nil, err
	}
	return &savedStock, nil
}

func (db RepositoryDB) GetStockByID(stockID int) (*domain.Stock, error) {
	sql := "SELECT price, name, symbol, exchange from stock WHERE stock_id= $1"
	row := db.Client.QueryRow(sql, stockID)
	var stock domain.Stock
	err := row.Scan(&stock.Price, &stock.Name, &stock.Symbol, &stock.Exchange)
	if err != nil {
		return nil, err
	}
	return &stock, nil
}

func (db RepositoryDB) GetStocks() ([]domain.Stock, error) {
	sql := "SELECT * from stock"
	rows, err := db.Client.Query(sql)
	if err != nil {
		return nil, err
	}
	var stocks []domain.Stock
	for rows.Next() {
		var s domain.Stock
		err = rows.Scan(&s.StockID, &s.Name, &s.Price, &s.Symbol, &s.Exchange, &s.LastUpdated)
		if err != nil {
			return nil, err
		}
		stocks = append(stocks, s)
	}

	return stocks, nil
}

func (db RepositoryDB) UpdateStock(symbol string, price float64) error {
	query := "UPDATE stock SET price = $2 where symbol = $1"
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	_, err := db.Client.QueryContext(ctx, query, symbol, price)
	if err != nil {
		return err
	}
	return nil
}

func NewStockRepositoryDB(dbClient *sqlx.DB) RepositoryDB {
	return RepositoryDB{Client: dbClient}
}
