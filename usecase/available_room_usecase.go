package usecase

import (
	"gokost.com/m/model"
	"gokost.com/m/repository"
)

type AvailableRoomUseCase interface {
	AvailableRoom() ([]model.BoardingRoom, error)
}

type availableRoomUseCase struct {
	repo repository.BoardingRoomRepo
}

func (a *availableRoomUseCase) AvailableRoom() ([]model.BoardingRoom, error) {
	return a.repo.GetAllAvailableRoom()
}

func NewAvailableRoomUseCase(repo repository.BoardingRoomRepo) AvailableRoomUseCase {
	return &availableRoomUseCase{
		repo,
	}
}
