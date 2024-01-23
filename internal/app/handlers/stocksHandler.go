package handlers

import (
	"encoding/json"
	"fmt"
	"github.com/cjmacharia/portfolioPilot/internal/app/service"
	"github.com/cjmacharia/portfolioPilot/internal/domain"
	"net/http"
)

type StockHandler struct {
	service service.StockService
}

func (sh *StockHandler) PostStocksHandler(w http.ResponseWriter, r *http.Request) {
	s := &domain.Stock{}
	err := json.NewDecoder(r.Body).Decode(s)
	if err != nil {
		fmt.Println("err", err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
	stock, err := sh.service.PostStock(s)
	fmt.Println(err)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(stock)
}
