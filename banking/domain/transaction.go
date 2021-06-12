package domain

import (
	"github.com/neoyewchuan/RestDevGo/banking/dto"
	"github.com/neoyewchuan/RestDevGo/banking/errs"
)

type Transaction struct {
	CustomerID      string  `json:"customer-id" xml:"customer-id db:"customer_id"`
	TransactionID   string  `json:"transaction-id" xml:"transaction-id" db:"transaction_id"`
	AccountID       string  `json:"account-id" xml:"account-id" db:"account_id"`
	TransactionDate string  `json:"transaction-date" xml:"transaction-date" db:"transaction_date"`
	TransactionType string  `json:"transaction-type" xml:"transaction-type" db:"transaction_type"`
	Amount          float64 `json:"amount" xml:"amount" db:"amount"`
}

func (txn Transaction) ToDto() dto.NewTransactionResponse {

	return dto.NewTransactionResponse{
		TransactionID: txn.TransactionID,
		LegalBalance:  txn.Amount,
		AvailBalance:  txn.Amount - 1,
	}
}

type TransactionRepository interface {
	Deposit(Transaction) (*Transaction, *errs.AppError)
	Withdrawal(Transaction) (*Transaction, *errs.AppError)
}
