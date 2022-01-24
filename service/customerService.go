package service

import (
	"github.com/bankingApp/domain"
	"github.com/bankingApp/errs"
)

type CustomerService interface {
	GetAllCustomers(string) ([]domain.Customer, *errs.AppError)
	GetCustomer(string) (*domain.Customer, *errs.AppError)
}

type DefaultCustomerService struct {
	repo domain.CustomerRepository
}

func (s DefaultCustomerService) GetAllCustomers(status string) ([]domain.Customer, *errs.AppError) {

	if status != "" {

		return s.repo.FindAllByStatus(status)

	} else {

		return s.repo.FindAll()
	}
}

func (s DefaultCustomerService) GetCustomer(id string) (*domain.Customer, *errs.AppError) {

	return s.repo.FindCustomerById(id)
}

func NewCustomerService(repository domain.CustomerRepository) DefaultCustomerService {

	return DefaultCustomerService{repository}
}
