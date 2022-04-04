package manager

import "gokost.com/m/usecase"

type UseCaseManager interface {
	AvailableRoomUseCase() usecase.AvailableRoomUseCase
	CustomerTransactionUseCase() usecase.CustomerTransactionUseCase
	InsertTransactionUseCase() usecase.InsertTransactionUseCase
	ListCustomerUseCase() usecase.ListCustomerUseCase
	UpdateCustomerUseCase() usecase.UpdateTransactionUseCase
}

type useCaseManager struct {
	repo RepoManager
}

func (u *useCaseManager) AvailableRoomUseCase() usecase.AvailableRoomUseCase {
	return usecase.NewAvailableRoomUseCase(u.repo.BoardingRoomRepo())
}

func (u *useCaseManager) CustomerTransactionUseCase() usecase.CustomerTransactionUseCase {
	return usecase.NewCustomerTransactionUseCase(u.repo.TransactionRepo())
}

func (u *useCaseManager) InsertTransactionUseCase() usecase.InsertTransactionUseCase {
	return usecase.NewInsertTransactionUseCase(u.repo.TransactionRepo())
}

func (u *useCaseManager) ListCustomerUseCase() usecase.ListCustomerUseCase {
	return usecase.NewListCustomerUseCase(u.repo.CustomerRepo())
}

func (u *useCaseManager) UpdateCustomerUseCase() usecase.UpdateTransactionUseCase {
	return usecase.NewUpdateTransactionUseCase(u.repo.TransactionRepo())
}

func NewUseCaseManager(repo RepoManager) UseCaseManager {
	return &useCaseManager{
		repo,
	}
}
