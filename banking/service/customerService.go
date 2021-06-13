package service

import (
	"github.com/neoyewchuan/RestDevGo/banking/domain"
	"github.com/neoyewchuan/RestDevGo/banking/dto"
	"github.com/neoyewchuan/RestDevGo/banking/errs"
)

type CustomerService interface {
	GetAllCustomers(string) ([]dto.CustomerResponse, *errs.AppError)
	GetCustomer(string) (*dto.CustomerResponse, *errs.AppError)
}

type DefaultCustomerService struct {
	repo domain.CustomerRepository
}

func (dcs DefaultCustomerService) GetAllCustomers(status string) ([]dto.CustomerResponse, *errs.AppError) {
	switch status {
	case "active", "1":
		status = "1"
	case "inactive", "0":
		status = "0"
	default:
		status = ""
	}
	c, err := dcs.repo.FindAll(status)
	if err != nil {
		return nil, err
	}
	var response []dto.CustomerResponse
	for _, v := range c {
		response = append(response, v.ToDto())
	}
	return response, nil
}

func (dcs DefaultCustomerService) GetCustomer(id string) (*dto.CustomerResponse, *errs.AppError) {
	c, err := dcs.repo.ByID(id)
	if err != nil {
		return nil, err
	}
	response := c.ToDto()

	return &response, nil
}

func NewCustomerService(repository domain.CustomerRepository) DefaultCustomerService {
	return DefaultCustomerService{repository}
}
