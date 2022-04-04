package repository

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"gokost.com/m/model"
)

type AdminRepo interface {
	GetAllAdmin() []model.AdminRoom
	GetAllCustomer() []model.Customer
}

type admin_repo struct {
	sqlxDb *sqlx.DB
}

func (a *admin_repo) GetAllAdmin() []model.AdminRoom {
	var data []model.AdminRoom
	a.sqlxDb.Select(&data, "select * from admin")
	return data
}

func (a *admin_repo) GetAllCustomer() []model.Customer {
	var data []model.Customer
	a.sqlxDb.Select(&data, "select name, address, ktpid, phone_number, start_rent_at, end_rent_at from customer")
	fmt.Println(data)
	return data
}

func NewAdminRepo(sqlx *sqlx.DB) AdminRepo {
	return &admin_repo{
		sqlxDb: sqlx,
	}
}
