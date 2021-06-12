package dto

import (
	"strings"

	"github.com/neoyewchuan/RestDevGo/banking/errs"
)

type NewAccountRequest struct {
	CustomerID  string `json:"customer-id" xml:"customer-id"`
	AccountType string `json:"account-type" xml:"account-type"`
	//OpeningDate string  `json:"opening-date" xml:"opening-date"`
	Amount float64 `json:"amount" xml:"amount"`
	//Status      string  `json:"status" xml:"status"`
}

func (nar *NewAccountRequest) Validate() *errs.AppError {
	if nar.Amount < 50 {
		return errs.NewValidationError("You need a minimum of S$50 to open a new account.")
	}
	if strings.ToLower(nar.AccountType) != "saving" && strings.ToLower(nar.AccountType) != "current" {
		return errs.NewValidationError("You can only open a saving or current account.")
	}
	return nil
}
