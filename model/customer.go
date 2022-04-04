package model

type Customer struct {
	Id          string
	Name        string
	Address     string
	Ktp         string `db:"ktpid"`
	PhoneNumber string `db:"phone_number"`
	StartRentAt string `db:"start_rent_at"`
	EndRentAt   string `db:"end_rent_at"`
}
