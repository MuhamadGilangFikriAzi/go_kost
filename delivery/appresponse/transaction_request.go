package appresponse

type TransactionRequest struct {
	CustomerId string `json:"customer_id"`
	RoomNumber string `json:"room_number"`
}
