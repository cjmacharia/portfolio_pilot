package database

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/cjmacharia/portfolioPilot/internal/domain"
	"github.com/jmoiron/sqlx"
	"log"
	"time"
)

func (db RepositoryDB) CreateTransaction(t *domain.Transaction) (*domain.Transaction, error) {
	s := "insert into transactions(transaction_type, stock_id, user_id, total_amount, quantity) VALUES($1, $2, $3, $4, $5)RETURNING transaction_type, total_amount, quantity"
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	query := `SELECT * FROM transactions WHERE transaction_type NOT IN ('DEPOSIT', 'BUY', 'SELL', 'WITHDRAW');`
	rows, err := db.Client.Query(query)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	for rows.Next() {
		var constraintCatalog, constraintSchema, constraintName, tableCatalog, tableName string
		var checkClause sql.NullString

		err := rows.Scan(&constraintCatalog, &constraintSchema, &constraintName, &tableCatalog, &tableName, &checkClause)
		if err != nil {
			log.Fatal(err)
		}

		fmt.Printf("Constraint Name: %s\n", constraintName)
		fmt.Printf("Table Name: %s\n", tableName)
		fmt.Printf("Check Clause: %s\n", checkClause.String)
	}
	if err := rows.Err(); err != nil {
		log.Fatal(err)
	}
	row := db.Client.QueryRowContext(ctx, s, t.TransactionType, t.StockID, t.UserID, t.TotalAmount, t.Quantity)
	var savedTransaction domain.Transaction
	err = row.Scan(&savedTransaction.TransactionType, &savedTransaction.TotalAmount, &savedTransaction.Quantity)
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

func (db RepositoryDB) GetUserTransactions(userID int) ([]domain.UserTransactionsResponse, error) {
	query := `SELECT transactions.transaction_type, transactions.total_amount, transactions.quantity, transactions.transaction_date,
stock.name from transactions LEFT JOIN stock on transactions.stock_id = stock.stock_id WHERE transactions.user_id = $1
AND (transactions.transaction_type IN ('DEPOSIT', 'WITHDRAW') OR stock.stock_id IS NOT NULL);`
	rows, err := db.Client.Query(query, userID)
	if err != nil {
		fmt.Println("An error occurred GetUserTransactions", err)
	}
	var transactions []domain.UserTransactionsResponse
	for rows.Next() {
		var t domain.UserTransactionsResponse
		err = rows.Scan(&t.TransactionType, &t.TotalAmount, &t.Quantity, &t.TransactionDate, &t.StockName)
		if err != nil {
			fmt.Println("An error occurred scanning", err)
		}
		transactions = append(transactions, t)
	}
	return transactions, nil
}

func NewTransactionRepositoryDB(dbClient *sqlx.DB) RepositoryDB {
	return RepositoryDB{Client: dbClient}
}
