package api

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func Test() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "testing",
		})
	}
}
