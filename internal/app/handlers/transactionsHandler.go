package handlers

import (
	"encoding/json"
	"fmt"
	"github.com/cjmacharia/portfolioPilot/internal/app/service"
	"github.com/cjmacharia/portfolioPilot/internal/domain"
	"github.com/cjmacharia/portfolioPilot/internal/dto"
	"github.com/cjmacharia/portfolioPilot/internal/utils"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

type TransactionsHandler struct {
	transactionService service.Transactions
	stockService       service.StockService
	userService        service.UserService
}

func (th *TransactionsHandler) HandleTransaction(w http.ResponseWriter, r *http.Request) {
	stockId, err := getID(w, r)
	if err != nil {
		http.Error(w, utils.ErrInvalidStockID.Error(), http.StatusBadRequest)
		return
	}
	stock, err := th.stockService.GetStockByID(stockId)
	if err != nil {
		http.Error(w, utils.ErrFetchStockPrice.Error(), http.StatusInternalServerError)
		return
	}
	t := &domain.Transaction{}
	err = json.NewDecoder(r.Body).Decode(&t)
	if err != nil {
		http.Error(w, utils.ErrInvalidPayload.Error(), http.StatusBadRequest)
		return
	}
	amount := utils.CalculateAmount(stock.Price, t.Quantity)

	currentWalletBalance, err := th.userService.GetUserWalletBalanceService(t.UserID)
	if err != nil {
		http.Error(w, utils.ErrCreateTransaction.Error(), http.StatusInternalServerError)
		return
	}
	var newBalance float64
	switch t.TransactionType {
	case "BUY":
		if amount > currentWalletBalance {
			http.Error(w, utils.ErrInsufficientFunds.Error(), http.StatusForbidden)
			return
		}
		newBalance = currentWalletBalance - amount
		t.TotalAmount = amount
		t.StockID = stockId
	case "SELL":
		var netQuantity int64
		netQuantity, err = th.transactionService.GetStockQuantityService(t.UserID, stockId)
		fmt.Println(netQuantity, t.Quantity, stockId, t.UserID, "pppppopopo")
		if err != nil {
			http.Error(w, utils.ErrCreateTransaction.Error(), http.StatusInternalServerError)
			return
		}
		if t.Quantity > netQuantity {
			http.Error(w, utils.ErrInsufficientStock.Error(), http.StatusForbidden)
			return
		}
		amount = utils.CalculateAmount(stock.Price, t.Quantity)
		newBalance = currentWalletBalance + amount
		t.TotalAmount = amount
		t.StockID = stockId

	case "DEPOSIT":
		newBalance = currentWalletBalance + t.TotalAmount
		t.StockID = nil

	case "WITHDRAW":

		newBalance = currentWalletBalance - t.TotalAmount
		t.StockID = nil
	}

	_, err = th.transactionService.NewTransaction(t)
	if err != nil {
		http.Error(w, utils.ErrCreateTransaction.Error(), http.StatusInternalServerError)
		return
	}
	err = th.userService.UpdateUserWalletBalanceService(t.UserID, newBalance)
	if err != nil {
		http.Error(w, utils.ErrInternalServerError.Error(), http.StatusInternalServerError)
		return
	}
	response := dto.TransactionResponse{
		TransactionType: t.TransactionType,
		Quantity:        t.Quantity,
		TotalAmount:     t.TotalAmount,
	}

	json.NewEncoder(w).Encode(response)
}

func getID(w http.ResponseWriter, r *http.Request) (int, error) {
	vars := mux.Vars(r)
	id, ok := vars["id"]
	if !ok {
		http.Error(w, utils.ErrInvalidStockID.Error(), http.StatusBadRequest)
	}
	urlId, err := strconv.Atoi(id)
	if err != nil {
		http.Error(w, utils.ErrMissingStockID.Error(), http.StatusBadRequest)
	}
	return urlId, nil
}

func (th *TransactionsHandler) GetUserTransactions(w http.ResponseWriter, r *http.Request) {
	id, err := getID(w, r)
	if err != nil {
		http.Error(w, utils.ErrFetchStockById.Error(), http.StatusInternalServerError)

	}
	t, err := th.transactionService.GetUserTransactions(id)
	if err != nil {
		http.Error(w, utils.ErrInternalServerError.Error(), http.StatusInternalServerError)

	}
	json.NewEncoder(w).Encode(t)
}

func (th *TransactionsHandler) GetUserPortfolio(w http.ResponseWriter, r *http.Request) {
	id, err := getID(w, r)
	var p []domain.Portfolio
	stocks, err := th.transactionService.GetUserStockService(id)
	if err != nil {
		http.Error(w, utils.ErrFetchStockById.Error(), http.StatusInternalServerError)

	}
	for _, stock := range stocks {
		s, err := th.stockService.GetStockByID(stock.StockID)
		if err != nil {
			http.Error(w, utils.ErrFetchStockByUserID.Error(), http.StatusInternalServerError)
		}
		portfolio := domain.Portfolio{
			Stock:       *s,
			NetQuantity: stock.NetQuantity,
		}
		p = append(p, portfolio)
	}
	json.NewEncoder(w).Encode(p)
}
