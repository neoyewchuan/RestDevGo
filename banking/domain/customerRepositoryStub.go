package domain

type CustomerRepositoryStub struct {
	customers []Customer
}

func (stub CustomerRepositoryStub) FindAll() ([]Customer, error) {
	return stub.customers, nil
}

func NewCustomerRepositoryStub() CustomerRepositoryStub {
	cust := []Customer{
		{ID: "001", Name: "Ashish", City: "New Delhi", ZipCode: "110075", DateOfBirth: "2000-01-01", Status: "1"},
		{ID: "002", Name: "Rob", City: "New Delhi", ZipCode: "110075", DateOfBirth: "1978-12-14", Status: "1"},
		{ID: "003", Name: "Peter Parker", City: "New York City", ZipCode: "11375", DateOfBirth: "2001-08-10", Status: "1"},
		{ID: "007", Name: "James Bond", City: "San Monique", ZipCode: "00007", DateOfBirth: "1968-04-13", Status: "1"},
	}
	return CustomerRepositoryStub{customers: cust}
}
