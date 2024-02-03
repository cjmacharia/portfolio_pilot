package handlers

import (
	"encoding/json"
	"github.com/cjmacharia/portfolioPilot/internal/app/service"
	"github.com/cjmacharia/portfolioPilot/internal/domain"
	"github.com/cjmacharia/portfolioPilot/internal/dto"
	"github.com/cjmacharia/portfolioPilot/internal/utils"
	"net/http"
)

type StockHandler struct {
	stockService service.StockService
}

func (sh *StockHandler) PostStocksHandler(w http.ResponseWriter, r *http.Request) {
	s := &domain.Stock{}
	err := json.NewDecoder(r.Body).Decode(s)
	if err != nil {
		http.Error(w, utils.ErrInvalidPayload.Error(), http.StatusBadRequest)
		return
	}
	_, err = sh.stockService.PostStock(s)
	if err != nil {
		http.Error(w, utils.ErrInvalidStockID.Error(), http.StatusInternalServerError)
		return
	}
	response := dto.StockResponse{
		Name:     s.Name,
		Symbol:   s.Name,
		Exchange: s.Exchange,
		Price:    s.Price,
	}
	json.NewEncoder(w).Encode(response)
}
