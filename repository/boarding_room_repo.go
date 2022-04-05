package repository

import (
	"github.com/jmoiron/sqlx"
	"gokost.com/m/model"
)

type BoardingRoomRepo interface {
	GetAllAvailableRoom() ([]model.BoardingRoom, error)
}

type boardingRoomRepo struct {
	sqlDb *sqlx.DB
}

func (b *boardingRoomRepo) GetAllAvailableRoom() ([]model.BoardingRoom, error) {
	var data []model.BoardingRoom
	err := b.sqlDb.Select(&data, "select * from boarding_room where is_available = true")
	if err != nil {
		return nil, err
	}
	return data, nil
}

func NewBoardingRoomRepo(sqlxDb *sqlx.DB) BoardingRoomRepo {
	return &boardingRoomRepo{
		sqlxDb,
	}
}
