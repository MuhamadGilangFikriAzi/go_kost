package manager

import (
	"github.com/jmoiron/sqlx"
	"gokost.com/m/repository"
)

type RepoManager interface {
	TransactionRepo() repository.TransactionRepo
	BoardingRoomRepo() repository.BoardingRoomRepo
	CustomerRepo() repository.CustomerRepo
}

type repoManager struct {
	SqlxDb *sqlx.DB
}

func (r *repoManager) TransactionRepo() repository.TransactionRepo {
	return repository.NewTransactionRepo(r.SqlxDb)
}

func (r *repoManager) BoardingRoomRepo() repository.BoardingRoomRepo {
	return repository.NewBoardingRoomRepo(r.SqlxDb)
}

func (r *repoManager) CustomerRepo() repository.CustomerRepo {
	return repository.NewCustomerRepo(r.SqlxDb)
}

func NewRepoManager(sqlxDb *sqlx.DB) RepoManager {
	return &repoManager{
		sqlxDb,
	}
}
