package manager

import (
	"github.com/jmoiron/sqlx"
	"gokost.com/m/config"
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

func NewInfraManager() InfraManager {
	return &infraManager{
		postgreConn: config.NewConfig().PostgreConn(),
	}
}
