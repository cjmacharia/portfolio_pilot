package database

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/cjmacharia/portfolioPilot/internal/domain"
	"github.com/cjmacharia/portfolioPilot/internal/utils"
	"github.com/jmoiron/sqlx"
	"log"
	"time"
)

func (db RepositoryDB) SignUp(u *domain.User) (int, error) {
	var exists bool
	err := db.Client.QueryRow("SELECT EXISTS(SELECT 1 FROM users where email= $1 )", u.Email).Scan(&exists)
	if err != nil {
		log.Println("Error checking existence:", err)
		return 0, err
	}
	if exists {
		return 0, utils.ErrDuplicateEmail
	}
	query := "INSERT into users(email, name, password_hash, wallet_balance)VALUES ($1, $2, $3, $4) RETURNING user_id"
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	row := db.Client.QueryRowContext(ctx, query, u.Email, u.Name, u.Password, u.WalletBalance)
	var userID int
	err = row.Scan(&userID)
	if ctx.Err() == context.DeadlineExceeded {
		return 0, context.DeadlineExceeded
	}
	if err != nil {
		return 0, err
	}
	return userID, nil
}

func (db RepositoryDB) GetUserByEmail(email string) (domain.User, error) {
	query := "SELECT * from Users where email=$1"
	row := db.Client.QueryRow(query, email)
	var user domain.User

	err := row.Scan(&user.UserID, &user.Name, &user.Password, &user.Email, &user.WalletBalance)
	if err != nil {
		if err == sql.ErrNoRows {
			fmt.Println(err, "k")
			return domain.User{}, utils.ErrUserNotFound
		}
		return domain.User{}, err
	}
	fmt.Println(user, "k")
	return user, nil
}

func (db RepositoryDB) GetUserWalletBalance(userID int) (float64, error) {
	query := "SELECT wallet_balance from Users where user_id=$1"
	row := db.Client.QueryRow(query, userID)
	var walletBalance float64
	err := row.Scan(&walletBalance)
	if err != nil {
		if err == sql.ErrNoRows {
			return walletBalance, utils.ErrUserNotFound
		}
		return walletBalance, err
	}
	return walletBalance, nil

}
func (db RepositoryDB) UpdateUserWalletBalance(userID int, newBalance float64) error {
	query := "UPDATE users SET wallet_balance = $1 WHERE user_id = $2"
	_, err := db.Client.Exec(query, newBalance, userID)
	if err != nil {
		return err
	}
	return nil
}

func NewUserRepositoryDB(dbClient *sqlx.DB) RepositoryDB {
	return RepositoryDB{Client: dbClient}
}
