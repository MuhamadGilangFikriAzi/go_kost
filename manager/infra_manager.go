package manager

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"os"
)

type InfraManager interface {
	PostgreConn() *sqlx.DB
}

type infraManager struct {
	postgreConn *sqlx.DB
}

func (i *infraManager) PostgreConn() *sqlx.DB {
	return i.postgreConn
}

func NewInfraManager(urlPostgresql string) InfraManager {
	conn, err := sqlx.Connect("pgx", urlPostgresql)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}
	return &infraManager{
		postgreConn: conn,
	}
}
