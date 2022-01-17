package service

import "github.com/bankingApp/domain"

type CustomerService interface {
	GetAllCustomers() ([]domain.Customer, error)
	GetCustomer(string) (*domain.Customer, error)
}

type DefaultCustomerService struct {
	repo domain.CustomerRepository
}

func (s DefaultCustomerService) GetAllCustomers() ([]domain.Customer, error) {

	return s.repo.FindAll()
}

func (s DefaultCustomerService) GetCustomer(id string) (*domain.Customer, error) {

	return s.repo.FindCustomerById(id)
}

func NewCustomerService(repository domain.CustomerRepository) DefaultCustomerService {

	return DefaultCustomerService{repository}
}