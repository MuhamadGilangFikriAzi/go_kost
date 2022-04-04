package repository

import (
	"database/sql"
	"fmt"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"gokost.com/m/delivery/appresponse"
	"gokost.com/m/model"
	"gokost.com/m/utility"
)

type TransactionRepo interface {
	InsertTransaction(dataTransaction appresponse.TransactionRequest) (sql.Result, error)
	UpdateTransaction(customer_id string)
}

type transactionRepo struct {
	sqlxDb *sqlx.DB
}

func (t *transactionRepo) InsertTransaction(dataTransaction appresponse.TransactionRequest) (sql.Result, error) {
	uuid := uuid.New().String()
	adminId := "olsdinoinqw"
	thisDay := utility.ThisDay()
	result, err := t.sqlxDb.Exec("insert into transaction_room(id, admin_id, customer_id, room_number, created_at, status) values($1, $2, $3, $4, $5, $6)", uuid, adminId, dataTransaction.CustomerId, dataTransaction.RoomNumber, thisDay, 1)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (t *transactionRepo) UpdateTransaction(customer_id string) {
	//thisDay := utility.ThisDay()
	var getData model.TransactionRoom
	err := t.sqlxDb.Get(&getData, "select * from transaction_room where customer_id = $1 order by created_at desc limit 1", customer_id)
	if err != nil {
		panic(err)
	}
	fmt.Println(getData)
	//result, err := t.sqlxDb.Exec("update transaction_room set status = 3, updated_at = $1  where customer_id = $2 order by created_at desc limit 1", thisDay, customer_id)
	//if err != nil {
	//	panic(err)
	//	return nil, err
	//}
	//return result, nil
}

func NewTransactionRepo(sqlxDb *sqlx.DB) TransactionRepo {
	return &transactionRepo{
		sqlxDb: sqlxDb,
	}
}
