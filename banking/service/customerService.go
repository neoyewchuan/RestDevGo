package service

import (
	"github.com/neoyewchuan/RestDevGo/banking/domain"
	"github.com/neoyewchuan/RestDevGo/banking/errs"
)

type CustomerService interface {
	GetAllCustomer() ([]domain.Customer, error)
	GetCustomer(string) (*domain.Customer, error)
}

type DefaultCustomerService struct {
	repo domain.CustomerRepository
}

func (dcs DefaultCustomerService) GetAllCustomer() ([]domain.Customer, *errs.AppError) {
	return dcs.repo.FindAll()
}

func (dcs DefaultCustomerService) GetCustomer(id string) (*domain.Customer, *errs.AppError) {
	return dcs.repo.ByID(id)
}

func NewCustomerService(repository domain.CustomerRepository) DefaultCustomerService {
	return DefaultCustomerService{repository}
}
