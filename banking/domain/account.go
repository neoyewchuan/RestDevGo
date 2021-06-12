package domain

import (
	"github.com/neoyewchuan/RestDevGo/banking/dto"
	"github.com/neoyewchuan/RestDevGo/banking/errs"
)

type Account struct {
	AccountID   string  `json:"account-id" xml:"account-id" db:"account_id"`
	CustomerID  string  `json:"customer-id" xml:"customer-id" db:"customer_id"`
	OpeningDate string  `json:"opening-date" xml:"opening-date" db:"opening_date"`
	AccountType string  `json:"account-type" xml:"account-type" db:"account_type"`
	Amount      float64 `json:"amount" xml:"amount" db:"amount"`
	Status      string  `json:"status" xml:"status" db:"status"`
}

func (a Account) ToDto() dto.NewAccountResponse {

	return dto.NewAccountResponse{
		AccountID: a.AccountID,
	}
}

type AccountRepository interface {
	Save(Account) (*Account, *errs.AppError)
}
