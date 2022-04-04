package usecase

import (
	"gokost.com/m/delivery/appresponse"
	"gokost.com/m/repository"
)

type InsertTransactionUseCase interface {
	InsertTransaction(dataTransaction appresponse.TransactionRequest) error
}

type insertTransactionUseCase struct {
	repo repository.TransactionRepo
}

func (i *insertTransactionUseCase) InsertTransaction(dataTransaction appresponse.TransactionRequest) error {
	return i.repo.InsertTransaction(dataTransaction)
}

func NewInsertTransactionUseCase(repo repository.TransactionRepo) InsertTransactionUseCase {
	return &insertTransactionUseCase{
		repo,
	}
}
