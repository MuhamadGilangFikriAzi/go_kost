package main

import (
	_ "github.com/jackc/pgx/v4/stdlib"
	"gokost.com/m/manager"
	"gokost.com/m/repository"
)

func main() {
	conn := manager.NewInfraManager().PostgreConn()
	repo := repository.NewTransactionRepo(conn)
	//repo.GetAllAdmin()
	//repo.GetAllCustomer()
	//repo.GetCustomerWithTransaction()
	//repo.GetAllAvailableRoom()
	//data := appresponse.TransactionRequest{
	//	"asidnlnasiufbqowmuqf",
	//	"A03",
	//}
	repo.UpdateTransaction("asidnlnasiufbqowmuqf")
	//defer conn.Close(context.Background())
}
