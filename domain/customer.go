package domain

import (
	"github.com/bankingApp/dto"
	"github.com/bankingApp/errs"
)

type Customer struct {
	Id          string `db:"customer_id"`
	Name        string
	City        string
	ZipCode     string
	DateOfBirth string `db:"date_of_birth"`
	Status      string
}

func (c Customer) statusAsText() string {
	statusAsText := "active"
	if c.Status == "0" {
		statusAsText = "inactive"
	}

	return statusAsText
}

func (c Customer) ToDto() dto.CustomerResponse {

	return dto.CustomerResponse{
		Id:          c.Id,
		Name:        c.Name,
		City:        c.City,
		ZipCode:     c.ZipCode,
		DateOfBirth: c.DateOfBirth,
		Status:      c.statusAsText(),
	}
}

func (c Customer) ToDtoArray(customers []Customer) []dto.CustomerResponse {

	var customersResponse []dto.CustomerResponse

	for _, c := range customers {

		customersResponse = append(customersResponse, c.ToDto())

	}

	return customersResponse

}

type CustomerRepository interface {
	FindAll(string) ([]Customer, *errs.AppError)
	FindCustomerById(string) (*Customer, *errs.AppError)
	FindAllByStatus(status string) ([]Customer, *errs.AppError)
}
