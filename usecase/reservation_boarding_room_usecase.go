package usecase

import (
	"gokost.com/m/delivery/apprequest"
	"gokost.com/m/repository"
)

type InsertTransactionUseCase interface {
	InsertTransaction(dataTransaction apprequest.TransactionRequest) error
}

type insertTransactionUseCase struct {
	repo repository.TransactionRepo
}

func (i *insertTransactionUseCase) InsertTransaction(dataTransaction apprequest.TransactionRequest) error {
	return i.repo.InsertTransaction(dataTransaction)
}

func NewInsertTransactionUseCase(repo repository.TransactionRepo) InsertTransactionUseCase {
	return &insertTransactionUseCase{
		repo,
	}
}
