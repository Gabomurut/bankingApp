package domain

import (
	"github.com/bankingApp/dto"
	"github.com/bankingApp/errs"
)

type Account struct {
	AccountId   string
	CustomerId  string
	OpeningDate string
	AccountType string
	Amount      float64
	Status      string
}

func (a Account) ToNewAccountResponseDto() dto.NewAccountResponse {

	return dto.NewAccountResponse{AccountId: a.AccountId}
}

type AccountRepository interface {
	Save(Account) (*Account, *errs.AppError)
}
