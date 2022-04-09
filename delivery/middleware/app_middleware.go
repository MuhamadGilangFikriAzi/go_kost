package middleware

import (
	"WMB/logger"
	"github.com/dgrijalva/jwt-go"
	"github.com/jmoiron/sqlx"
	"time"
)

type authHeader struct {
	AuthorizationHeader string `header:"Authorization"`
}

type TokenConfig struct {
	ApplicationName     string
	JwtSignatureKey     string
	JwtSigningMethod    *jwt.SigningMethodHMAC
	AccessTokenLifeTime time.Duration
}

type token struct {
	Config TokenConfig
}

func AuthUser(requiredToken string, sqlxdb *sqlx.DB) bool {
	var countAuthUser int
	err := sqlxdb.Get(&countAuthUser, "select count(*) from users where token = $1", requiredToken)
	if err != nil {
		logger.Log.Err(err).Msg("Database not conected")
		panic("database not connected")
	}
	if countAuthUser == 0 {
		logger.Log.Err(err).Msg("Token Auth not found")
		return true
	}
	return false
}
