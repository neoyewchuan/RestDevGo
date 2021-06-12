package service

import (
	"time"

	"github.com/neoyewchuan/RestDevGo/banking/domain"
	"github.com/neoyewchuan/RestDevGo/banking/dto"
	"github.com/neoyewchuan/RestDevGo/banking/errs"
)

type AccountService interface {
	NewAccount(dto.NewAccountRequest) (*dto.NewAccountResponse, *errs.AppError)
}

type DefaultAccountService struct {
	repo domain.AccountRepository
}

func (s DefaultAccountService) NewAccount(na dto.NewAccountRequest) (*dto.NewAccountResponse, *errs.AppError) {
	err := na.Validate()
	if err != nil {
		return nil, err
	}
	acct := domain.Account{
		AccountID:   "",
		CustomerID:  na.CustomerID,
		OpeningDate: time.Now().Format("2006-01-02 15:04:05"),
		AccountType: na.AccountType,
		Amount:      na.Amount,
		Status:      "1",
	}
	newAccount, err := s.repo.Save(acct)
	if err != nil {
		return nil, err
	}
	response := newAccount.ToDto()
	return &response, nil
}

func NewAccountService(repository domain.AccountRepository) DefaultAccountService {
	return DefaultAccountService{repository}
}
