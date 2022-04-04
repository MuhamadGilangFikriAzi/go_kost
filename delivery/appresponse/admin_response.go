package appresponse

import "github.com/jackc/pgtype"

type CustomerLastTransactionResponse struct {
	Name            string      `json:"customer_name"`
	PhoneNumber     string      `json:"phone_number" db:"phone_number"`
	RoomNumber      string      `db:"room_number" json:"room_number"`
	LastTransaction pgtype.Date `db:"last_transaction" json:"last_transaction"`
	Status          string
}
