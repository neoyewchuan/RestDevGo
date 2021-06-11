package service

import (
	"github.com/neoyewchuan/RestDevGo/banking/domain"
	"github.com/neoyewchuan/RestDevGo/banking/errs"
)

type CustomerService interface {
	GetAllCustomer(string) ([]domain.Customer, *errs.AppError)
	GetCustomer(string) (*domain.Customer, *errs.AppError)
}

type DefaultCustomerService struct {
	repo domain.CustomerRepository
}

func (dcs DefaultCustomerService) GetAllCustomer(status string) ([]domain.Customer, *errs.AppError) {
	switch status {
	case "active", "1":
		status = "1"
	case "inactive", "0":
		status = "0"
	default:
		status = ""
	}
	return dcs.repo.FindAll(status)
}

func (dcs DefaultCustomerService) GetCustomer(id string) (*domain.Customer, *errs.AppError) {
	return dcs.repo.ByID(id)
}

func NewCustomerService(repository domain.CustomerRepository) DefaultCustomerService {
	return DefaultCustomerService{repository}
}
