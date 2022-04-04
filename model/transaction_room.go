package model

type TransactionRoom struct {
	Id         string
	AdminId    string `db:"admin_id"`
	CusomerId  string `db:"customer_id"`
	RoomNumber string `db:"room_number"`
	Date       string
	Status     int // 1 = pending, 2 = not paid, 3 = finish
}
