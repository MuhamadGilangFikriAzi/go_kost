package config

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"github.com/spf13/viper"
	"os"
)

type config struct {
	dbConn *sqlx.DB
}

func (c *config) PostgreConn() *sqlx.DB {
	return c.dbConn
}

func newPostgreConn() *sqlx.DB {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")
	err := viper.ReadInConfig() // Find and read the config file
	if err != nil {             // Handle errors reading the config file
		panic(fmt.Errorf("Fatal error config file: %w \n", err))
	}
	dbName := viper.GetString("DBNAME")
	dbHost := viper.GetString("DBHOST")
	dbUsername := viper.GetString("DBUSERNAME")
	dbPassword := viper.GetString("DBPASSWORD")
	dbPort := viper.GetString("DBPORT")
	urlDb := fmt.Sprintf("postgres://%s:%s@%s:%s/%s", dbUsername, dbPassword, dbHost, dbPort, dbName)
	fmt.Println(urlDb)
	// urlExample := "postgres://username:password@localhost:5432/database_name"
	conn, err := sqlx.Connect("pgx", urlDb)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}
	return conn
}

func NewConfig() *config {
	return &config{
		newPostgreConn(),
	}
}
