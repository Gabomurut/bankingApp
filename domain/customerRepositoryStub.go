package domain

type CustomerRepositoryStub struct {
	customers []Customer
}

func (s CustomerRepositoryStub) FindAll() ([]Customer, error) {
	return s.customers, nil
}

func NewCustomerRepositoryStub() CustomerRepositoryStub {
	customers := []Customer{
		{"1001", "Fran", "Cordoba", "1420", "1982-12-20", "1"},
		{"1002", "Gabo", "Caba", "1419", "1982-05-04", "1"},
		{"1003", "Max", "Mendoza", "2324", "1982-12-10", "1"},
	}
	return CustomerRepositoryStub{customers}

}
