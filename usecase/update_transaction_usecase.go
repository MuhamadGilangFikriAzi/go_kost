package usecase

import "gokost.com/m/repository"

type UpdateTransactionUseCase interface {
	UpdateTransaction(customerId string, status int) error
}

type updateTransactionUseCase struct {
	repo repository.TransactionRepo
}

func (u *updateTransactionUseCase) UpdateTransaction(customerId string, status int) error {
	return u.repo.UpdateTransaction(customerId, status)
}

func NewUpdateTransactionUseCase(repo repository.TransactionRepo) UpdateTransactionUseCase {
	return &updateTransactionUseCase{
		repo,
	}
}
