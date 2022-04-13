package manager

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"os"
)

type InfraManager interface {
	PostgreConn() *sqlx.DB
	MysqlConn() *gorm.DB
}

type infraManager struct {
	postgreConn *sqlx.DB
	mysqlConn   *gorm.DB
}

func (i *infraManager) PostgreConn() *sqlx.DB {
	return i.postgreConn
}

func (i *infraManager) MysqlConn() *gorm.DB {
	return i.mysqlConn
}

func NewInfraManager(urlPostgresql string, urlMysql string) InfraManager {
	conn, err := sqlx.Connect("pgx", urlPostgresql)
	connMysql, errMysql := gorm.Open(mysql.Open(urlMysql), &gorm.Config{})
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}
	if errMysql != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", errMysql)
		os.Exit(1)
	}
	return &infraManager{
		postgreConn: conn,
		mysqlConn:   connMysql,
	}
}
