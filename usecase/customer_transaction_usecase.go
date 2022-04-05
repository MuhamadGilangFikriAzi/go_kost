package usecase

import (
	"gokost.com/m/delivery/appresponse"
	"gokost.com/m/repository"
)

type CustomerTransactionUseCase interface {
	SearchTransactionByCustomerId(customerId string) ([]appresponse.UserTransactionResonse, error)
}

type customerTransactionUseCase struct {
	repo repository.TransactionRepo
}

func (c *customerTransactionUseCase) SearchTransactionByCustomerId(customerId string) ([]appresponse.UserTransactionResonse, error) {
	return c.repo.GetTransactionByCustomerId(customerId)
}

func NewCustomerTransactionUseCase(repo repository.TransactionRepo) CustomerTransactionUseCase {
	return &customerTransactionUseCase{
		repo,
	}
}
