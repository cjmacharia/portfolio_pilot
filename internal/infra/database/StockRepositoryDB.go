package database

import (
	"context"
	"fmt"
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
		fmt.Println("No rows returned after INSERT")
	}
	return &savedStock, nil
}

func NewStockRepositoryDB(dbClient *sqlx.DB) RepositoryDB {
	return RepositoryDB{Client: dbClient}
}
