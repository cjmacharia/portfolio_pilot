package dto

import (
	"github.com/cjmacharia/portfolioPilot/internal/domain"
)

type TransactionResponse struct {
	TransactionType domain.TransactionType `json:"transaction_type"`
	Quantity        int64                  `json:"quantity"`
	TotalAmount     float64                `json:"total_amount"`
}
