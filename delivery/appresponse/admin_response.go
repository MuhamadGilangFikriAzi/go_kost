package appresponse

type CustomerResponse struct {
	Name        string `json:"name"`
	Address     string `json:"address"`
	Ktp         string `db:"ktpid" json:"ktp"`
	PhoneNumber string `db:"phone_number" json:"phone_number"`
	StartRentAt string `db:"start_rent_at" json:"start_rent_at"`
	EndRentAt   string `db:"end_rent_at" json:"end_rent_at"`
}
