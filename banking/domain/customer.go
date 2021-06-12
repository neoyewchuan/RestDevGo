package domain

import (
	"github.com/neoyewchuan/RestDevGo/banking/dto"
	"github.com/neoyewchuan/RestDevGo/banking/errs"
)

type Customer struct {
	ID          string `json:"id" xml:"id" db:"customer_id"`
	Name        string `json:"full-name" xml:"full-name" db:"name"`
	City        string `json:"city" xml:"city" db:"city"`
	ZipCode     string `json:"zip-code" xml:"zip-code" db:"zipcode"`
	DateOfBirth string `json:"date-of-birth" xml:"date-of-birth" db:"date_of_birth"`
	Status      string `json:"status" xml:"status" db:"status"`
}

func (c Customer) statusAsText() string {
	statAsTxt := "active"
	if c.Status == "0" {
		statAsTxt = "inactive"
	}
	return statAsTxt
}

func (c Customer) ToDto() dto.CustomerResponse {

	return dto.CustomerResponse{
		ID:          c.ID,
		Name:        c.Name,
		DateOfBirth: c.DateOfBirth,
		City:        c.City,
		ZipCode:     c.ZipCode,
		Status:      c.statusAsText(),
	}
}

type CustomerRepository interface {
	FindAll(string) ([]Customer, *errs.AppError)
	ByID(string) (*Customer, *errs.AppError)
}
