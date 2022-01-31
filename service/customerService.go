package service

import (
	"github.com/bankingApp/domain"
	"github.com/bankingApp/dto"
	"github.com/bankingApp/errs"
	"github.com/bankingApp/logger"
)

type CustomerService interface {
	GetAllCustomers(string) (*[]dto.CustomerResponse, *errs.AppError)
	GetCustomer(string) (*dto.CustomerResponse, *errs.AppError)
}

type DefaultCustomerService struct {
	repo domain.CustomerRepository
}

func (s DefaultCustomerService) GetAllCustomers(status string) (*[]dto.CustomerResponse, *errs.AppError) {

	c, err := s.repo.FindAll(status)

	if err != nil {
		return nil, err
	}

	response := c[0].ToDtoArray(c)

	logger.Info("Funciona el DTO para el array")

	return &response, nil
}

func (s DefaultCustomerService) GetCustomer(id string) (*dto.CustomerResponse, *errs.AppError) {

	c, err := s.repo.FindCustomerById(id)
	if err != nil {
		return nil, err
	}

	response := c.ToDto()

	return &response, nil
}

func NewCustomerService(repository domain.CustomerRepository) DefaultCustomerService {

	return DefaultCustomerService{repository}
}
