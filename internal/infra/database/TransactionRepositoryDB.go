package database

import (
	"context"
	"fmt"
	"github.com/cjmacharia/portfolioPilot/internal/domain"
	"github.com/jmoiron/sqlx"
	"log"
	"time"
)

func (db RepositoryDB) CreateTransaction(t *domain.Transaction) (*domain.Transaction, error) {
	sql := "insert into transactions(transaction_type, stock_id, user_id, total_amount, quantity) VALUES($1, $2, $3, $4, $5)RETURNING transaction_type, total_amount, quantity"
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	row := db.Client.QueryRowContext(ctx, sql, t.TransactionType, t.StockID, t.UserID, t.TotalAmount, t.Quantity)
	var savedTransaction domain.Transaction
	err := row.Scan(&savedTransaction.TransactionType, &savedTransaction.TotalAmount, &savedTransaction.Quantity)
	if err != nil {
		fmt.Println("An error occurred", err)
	}
	return &savedTransaction, nil
}

func (db RepositoryDB) GetStockQuantity(userId int, stockId int) (int64, error) {
	query := `
	SELECT COALESCE(SUM(CASE WHEN transaction_type = 'SELL' THEN -quantity ELSE quantity END), 0) AS net_quantity
	FROM transactions
	WHERE user_id = $2 AND stock_id = $1
	 AND (transaction_type = 'BUY' OR transaction_type = 'SELL');
	`
	var netQuantity int64
	err := db.Client.QueryRow(query, stockId, userId).Scan(&netQuantity)
	if err != nil {
		fmt.Println("An error occurred totalSellQuantity", err)
	}

	return netQuantity, nil
}

func (db RepositoryDB) GetUserStocks(userID int) ([]*domain.StockPortfolio, error) {
	query := `SELECT stock_id, COALESCE(SUM(CASE WHEN transaction_type = 'SELL' THEN -quantity ELSE quantity END), 0)
AS net_quantity FROM transactions WHERE user_id =  $1 AND (transaction_type = 'BUY' OR transaction_type = 'SELL')
GROUP BY stock_id`
	rows, err := db.Client.Query(query, userID)
	if err != nil {
		fmt.Println("An error occurred GetUserStocks", err)
	}
	defer rows.Close()
	var portfolios []*domain.StockPortfolio
	for rows.Next() {
		var stockPortfolio domain.StockPortfolio
		if err := rows.Scan(&stockPortfolio.StockID, &stockPortfolio.NetQuantity); err != nil {
			fmt.Println(err, "hold my tea")
			log.Fatal(err)
		}
		portfolios = append(portfolios, &stockPortfolio)
	}
	return portfolios, err
}

func NewTransactionRepositoryDB(dbClient *sqlx.DB) RepositoryDB {
	return RepositoryDB{Client: dbClient}
}
