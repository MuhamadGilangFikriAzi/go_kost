package config

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/spf13/viper"
	"gokost.com/m/authenticator"
	"time"
)

type Config struct {
	dbConn      string
	mysqlConn   string
	ConfigToken authenticator.Token
}

func newTokenConfig() authenticator.Token {
	tokenConfig := authenticator.TokenConfig{
		AplicationName:      "Warung Makan Bahari",
		JwtSignatureKey:     "P@ssw0rd",
		JwtSignatureMethod:  jwt.SigningMethodHS256,
		AccessTokenDuration: 600 * time.Second,
	}
	return authenticator.NewToken(tokenConfig)
}

func (c *Config) MysqlConn() string {
	return c.mysqlConn
}

func (c *Config) PostgreConn() string {
	return c.dbConn
}

func ReadConfigFile(configName string) {
	viper.SetConfigName(configName)
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")
	err := viper.ReadInConfig() // Find and read the Config file
	if err != nil {             // Handle errors reading the Config file
		panic(fmt.Errorf("Fatal error Config file: %w \n", err))
	}
}

func GetConfigValue(configName string) string {
	ReadConfigFile("Config")
	return viper.GetString(configName)
}

func newPostgreConn() string {
	dbName := GetConfigValue("DBNAME")
	dbHost := GetConfigValue("DBHOST")
	dbUsername := GetConfigValue("DBUSERNAME")
	dbPassword := GetConfigValue("DBPASSWORD")
	dbPort := GetConfigValue("DBPORT")
	urlDb := fmt.Sprintf("postgres://%s:%s@%s:%s/%s", dbUsername, dbPassword, dbHost, dbPort, dbName)
	fmt.Println(urlDb)
	// urlExample := "postgres://username:password@localhost:5432/database_name"

	return urlDb
}

func newMysqlConn() string {
	dbUser := GetConfigValue("MYSQLUSER")
	dbPass := GetConfigValue("MYSQLPASS")
	dbUrl := GetConfigValue("MYSQLURL")
	dbPort := GetConfigValue("MYSQLPORT")
	dbName := GetConfigValue("MYSQLDBNAME")
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", dbUser, dbPass, dbUrl, dbPort, dbName)
	return dsn
}

func NewConfig() *Config {
	return &Config{
		dbConn:      newPostgreConn(),
		ConfigToken: newTokenConfig(),
		mysqlConn:   newMysqlConn(),
	}
}
