package repository

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"gokost.com/m/delivery/appresponse"
	"gokost.com/m/model"
)

type CustomerRepo interface {
	GetAllCustomer() ([]model.Customer, error)
	GetCustomerWithTransaction() ([]appresponse.CustomerLastTransactionResponse, error)
}

type customerRepo struct {
	sqlDb *sqlx.DB
}

func (a *customerRepo) GetCustomerWithTransaction() ([]appresponse.CustomerLastTransactionResponse, error) {
	var data []appresponse.CustomerLastTransactionResponse
	//err := a.sqlxDb.Select(&data, "select name from customer")
	err := a.sqlDb.Select(&data, "select c.name , c.phone_number ,(select tr.room_number from transaction_room tr where customer_id = c.id order by tr.date  desc limit 1), (select tr.date from transaction_room tr where tr.customer_id = c.id order by tr.date desc limit 1) as last_transaction, (select (case when tr.status = 1 then 'Proses' when tr.status = 2 then 'Not Paid' when tr.status = 3 then 'Finish' end ) as status from transaction_room tr where tr.customer_id = c.id order by tr.date  desc limit 1) from customer c")
	if err != nil {
		return nil, err
	}
	return data, nil
}

func (a *customerRepo) GetAllCustomer() ([]model.Customer, error) {
	var data []model.Customer
	err := a.sqlDb.Select(&data, "select name, address, ktpid, phone_number, start_rent_at, end_rent_at from customer")
	if err != nil {
		return nil, err
	}
	fmt.Println(data)
	return data, nil
}

func NewCustomerRepo(sqlDb *sqlx.DB) CustomerRepo {
	return &customerRepo{
		sqlDb,
	}
}
