package domain

type Customer struct {
	ID          string `json:"id" xml:"id"`
	Name        string `json:"full-name" xml:"full-name"`
	City        string `json:"city" xml:"city"`
	ZipCode     string `json:"zip-code" xml:"zip-code"`
	DateOfBirth string `json:"date-of-birth" xml:"date-of-birth"`
	Status      string `json:"status" xml:"status"`
}

type CustomerRepository interface {
	FindAll() ([]Customer, error)
}
