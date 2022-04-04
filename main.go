package main

import (
	_ "github.com/jackc/pgx/v4/stdlib"
	"gokost.com/m/manager"
	"gokost.com/m/repository"
)

func main() {
	conn := manager.NewInfraManager().PostgreConn()
	repo := repository.NewAdminRepo(conn)
	//repo.GetAllAdmin()
	repo.GetAllCustomer()
	//defer conn.Close(context.Background())
}
