package dto

import (
	"strings"

	"github.com/neoyewchuan/RestDevGo/banking/errs"
)

type NewTransactionRequest struct {
	CustomerID      string  `json:"customer-id" xml:"customer-id" db:"customer_id"`
	AccountID       string  `json:"account-id" xml:"account-id" db:"account_id"`
	TransactionDate string  `json:"transaction-date" xml:"transaction-date" db:"transaction_date"`
	TransactionType string  `json:"transaction-type" xml:"transaction-type" db:"transaction_type"`
	Amount          float64 `json:"amount" xml:"amount" db:"amount"`
}

func (ntr *NewTransactionRequest) Validate() *errs.AppError {
	if strings.ToLower(ntr.TransactionType) != "withdrawal" && strings.ToLower(ntr.TransactionType) != "deposit" {
		return errs.NewValidationError("Transaction type not supported (deposit or withdrawal only).")
	}
	if ntr.Amount <= 0 {
		return errs.NewValidationError(ntr.TransactionType + " amount cannot be $0.00.")
	}

	return nil
}
