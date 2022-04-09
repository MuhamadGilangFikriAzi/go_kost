package authenticator

import (
	"fmt"
	"gokost.com/m/delivery/appresponse"
	"golang.org/x/oauth2/jwt"
	"time"
)

type Token interface {
	CreateToken(dataLogin appresponse.LoginResponse) (string, error)
	VerifAccessToken(tokenString string) (jwt.MapClaims, error)
}

type TokenConfig struct {
	AplicationName      string
	JwtSignatureKey     string
	JwtSignatureMethod  *jwt.SigningMethodHMAC
	AccessTokenDuration time.Duration
}

type token struct {
	config TokenConfig
}

func (t *token) CreateToken(dataLogin appresponse.LoginResponse) (string, error) {
	now := time.Now().UTC()
	//fmt.Println(t.config.AccessTokenDuration)
	end := now.Add(t.config.AccessTokenDuration)
	claims := MyClaims{ // Menyiapkan struct dengan isi yg dibutuhkan
		StandardClaims: jwt.StandardClaims{
			Issuer: t.config.AplicationName,
		},
		Username: dataLogin.Username,
		Name:     dataLogin.Name,
	}
	claims.IssuedAt = now.Unix()
	claims.ExpiresAt = end.Unix()
	token := jwt.NewWithClaims(t.config.JwtSignatureMethod, claims) // membuat jwt dengan format method dan claim berupa struct
	return token.SignedString([]byte(t.config.JwtSignatureKey))     // mendapatkan token
}

func (t *token) VerifAccessToken(tokenString string) (jwt.MapClaims, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if method, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Signid Method invalid")
		} else if method != t.config.JwtSignatureMethod {
			return nil, fmt.Errorf("signid Method invalid")
		}
		return []byte(t.config.JwtSignatureKey), nil
	})
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid {
		return nil, err
	}
	return claims, nil
}

func NewToken(config TokenConfig) Token {
	return &token{
		config,
	}
}
