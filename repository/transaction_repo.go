package repository

import (
	"fmt"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"gokost.com/m/delivery/appresponse"
	"gokost.com/m/model"
	"gokost.com/m/utility"
)

type TransactionRepo interface {
	InsertTransaction(dataTransaction appresponse.TransactionRequest) error
	UpdateTransaction(customer_id string, status int) error
	GetTransactionByCustomerId(customerId string) ([]appresponse.UserTransactionResonse, error)
}

type transactionRepo struct {
	sqlxDb *sqlx.DB
}

func (t *transactionRepo) InsertTransaction(dataTransaction appresponse.TransactionRequest) error {
	uuid := uuid.New().String()
	adminId := "olsdinoinqw"
	thisDay := utility.ThisDay()
	_, err := t.sqlxDb.Exec("insert into transaction_room(id, admin_id, customer_id, room_number, created_at, status) values($1, $2, $3, $4, $5, $6)", uuid, adminId, dataTransaction.CustomerId, dataTransaction.RoomNumber, thisDay, 1)
	if err != nil {
		return err
	}
	return nil
}

func (t *transactionRepo) UpdateTransaction(customer_id string, status int) error {
	thisDay := utility.ThisDay()
	var getData model.TransactionRoom
	err := t.sqlxDb.Get(&getData, "select * from transaction_room where customer_id = $1 order by created_at desc limit 1", customer_id)
	if err != nil {
		panic(err)
	}

	tx := t.sqlxDb.MustBegin()
	tx.Exec("update transaction_room set status = $1, updated_at = $2  where id = $3", status, thisDay, getData.Id)
	if status == 2 {
		tx.Exec("update boarding_room set is_available = true where boarding_room.room_number = $1", getData.RoomNumber)
	} else if status == 1 || status == 3 {
		tx.Exec("update boarding_room set is_available = false where boarding_room.room_number = $1", getData.RoomNumber)
	}
	errTx := tx.Commit()
	if errTx != nil {
		return errTx
	}
	return nil
}

func (t *transactionRepo) GetTransactionByCustomerId(customerId string) ([]appresponse.UserTransactionResonse, error) {
	var data []appresponse.UserTransactionResonse
	err := t.sqlxDb.Select(&data, "select c.name, br.room_number, br.price, tr.created_at as order_date, (case when tr.status = 1 then 'Pendding' when tr.status = 2 then 'Not Paid' when tr.status = 3 then 'Finish' end) as status from transaction_room tr inner join customer c on c.id = tr.customer_id inner join boarding_room br on br.room_number = tr.room_number where customer_id = $1", customerId)
	if err != nil {
		return nil, err
	}
	fmt.Println(data)
	return data, err
}

func NewTransactionRepo(sqlxDb *sqlx.DB) TransactionRepo {
	return &transactionRepo{
		sqlxDb: sqlxDb,
	}
}
