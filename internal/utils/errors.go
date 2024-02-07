package utils

import "errors"

var (
	ErrMissingStockID      = errors.New("missing stock_id parameter")
	ErrInvalidStockID      = errors.New("invalid stock_id parameter")
	ErrFetchStockPrice     = errors.New("failed to fetch stock price")
	ErrFetchStockById      = errors.New("failed to fetch stock")
	ErrInvalidPayload      = errors.New("invalid request payload")
	ErrCreateTransaction   = errors.New("failed to create transaction")
	ErrCreateStock         = errors.New("failed to create a stock")
	ErrHashPassword        = errors.New("failed to hash the password")
	ErrDatabaseRequest     = errors.New("failed to fetch data")
	ErrTokenGenerate       = errors.New("failed to generate token")
	ErrAuthenticate        = errors.New("failed to authenticate user")
	ErrStartServer         = errors.New("server not started on 8080")
	ErrUserNotFound        = errors.New("user not found")
	ErrDuplicateEmail      = errors.New("user already exists")
	ErrInsufficientFunds   = errors.New("you don not have enough money in your wallet")
	ErrInsufficientStock   = errors.New("you don not have enough stock to sell")
	ErrFetchStockByUserID  = errors.New("failed to fetch stock using the user ID")
	ErrInternalServerError = errors.New("something went wrong")
	ErrTokenSignature      = errors.New("please check your token and try again")
	ErrFetchAllStocks      = errors.New("failed to fetch available stock")
)
