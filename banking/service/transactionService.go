package service

import (
	"time"

	"github.com/neoyewchuan/RestDevGo/banking/domain"
	"github.com/neoyewchuan/RestDevGo/banking/dto"
	"github.com/neoyewchuan/RestDevGo/banking/errs"
)

type TransactionService interface {
	NewTransaction(dto.NewTransactionRequest) (*dto.NewTransactionResponse, *errs.AppError)
}

type DefaultTransactionService struct {
	repo domain.TransactionRepository
}

func (s DefaultTransactionService) NewTransaction(nt dto.NewTransactionRequest) (*dto.NewTransactionResponse, *errs.AppError) {
	err := nt.Validate()
	if err != nil {
		return nil, err
	}
	trx := domain.Transaction{
		TransactionID:   "",
		CustomerID:      nt.CustomerID,
		AccountID:       nt.AccountID,
		TransactionDate: time.Now().Format("2006-01-02 15:04:05"),
		TransactionType: nt.TransactionType,
		Amount:          nt.Amount,
	}
	var newTransaction *domain.Transaction
	switch trx.TransactionType {
	case "deposit":
		newTransaction, err = s.repo.Deposit(trx)
	case "withdrawal":
		newTransaction, err = s.repo.Withdrawal(trx)
	}
	if err != nil {
		return nil, err
	}
	response := newTransaction.ToDto()
	return &response, nil
}

func NewTransactionService(repository domain.TransactionRepository) DefaultTransactionService {
	return DefaultTransactionService{repository}
}
