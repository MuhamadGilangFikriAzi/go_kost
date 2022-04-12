package middleware

import (
	"github.com/gin-gonic/gin"
	authenticator2 "gokost.com/m/authenticator"
	"net/http"
	"strings"
)

type AuthTokenMiddleware struct {
	acctToken authenticator2.Token
}

func NewAuthTokenMiddleware(configToken authenticator2.Token) *AuthTokenMiddleware {
	return &AuthTokenMiddleware{
		acctToken: configToken,
	}
}

func (a *AuthTokenMiddleware) TokenAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		if c.Request.URL.Path == "/login/admin" || strings.Contains(c.Request.URL.Path, "/files") {
			c.Next()
		} else {
			h := authHeader{}
			err := c.ShouldBindHeader(&h)
			if err != nil {
				c.JSON(http.StatusConflict, gin.H{
					"Message": err,
				})
				c.Abort()
				return
			}
			tokenString := strings.Replace(h.AuthorizationHeader, "Bearer ", "", -1)
			if tokenString == "" {
				c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
					"http_code": http.StatusUnauthorized,
					"meessage":  "Unautherized",
					"service":   "Token-auth",
				})
				//e := commonresp.NewErrorMessage(http.StatusUnauthorized, "Token-auth", "01", "Unautherized")
				//c.Abort()
				//c.Error(fmt.Errorf("%s", e.ToJson()))
				//c.JSON("Error")
				return
			}
			token, errToken := a.acctToken.VerifAccessToken(tokenString)
			if errToken != nil {
				c.JSON(http.StatusUnauthorized, gin.H{
					"message": errToken.Error(),
				})
				c.Abort()
				return
			}
			if token["iss"] == a.acctToken.GetAppName() {
				c.Next()
			} else {
				c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
					"http_code": http.StatusUnauthorized,
					"meessage":  "Unautherized",
					"service":   "Token-auth",
				})
				c.Abort()
				return
			}
			c.Next()
		}
	}
}
