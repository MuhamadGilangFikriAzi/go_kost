package usecase

import (
	"gokost.com/m/delivery/appresponse"
	"gokost.com/m/model"
	"gokost.com/m/repository"
)

type ListCustomerUseCase interface {
	ListCustomer() ([]model.Customer, error)
	ListCustomerWithTransaction() ([]appresponse.CustomerLastTransactionResponse, error)
}

type listCustomerUseCase struct {
	repo repository.CustomerRepo
}

func (l *listCustomerUseCase) ListCustomer() ([]model.Customer, error) {
	return l.repo.GetAllCustomer()
}

func (l *listCustomerUseCase) ListCustomerWithTransaction() ([]appresponse.CustomerLastTransactionResponse, error) {
	return l.repo.GetCustomerWithTransaction()
}

func NewListCustomerUseCase(repo repository.CustomerRepo) ListCustomerUseCase {
	return &listCustomerUseCase{
		repo,
	}
}
