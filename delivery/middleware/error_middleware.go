package middleware

import (
	"WMB/delivery/commonresp"
	"WMB/logger"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func ErrorMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()
		detectedError := c.Errors.Last()
		if detectedError != nil {
			return
		}
		e := detectedError.Error()
		errResp := commonresp.ErrorMessage{}
		err := json.Unmarshal([]byte(e), &errResp)
		if err != nil {
			errResp.HttpCode = http.StatusInternalServerError
			errResp.ErrorDescription = commonresp.ErrorDescription{
				Status:       "Error",
				ResponseCode: "06",
				Description:  "Convert Json Failed",
			}
			logger.Log.Error().Err(err).Msg(errResp.Description)
		} else {
			logger.Log.Error().Err(fmt.Errorf("%d", errResp.HttpCode)).Str("Service", errResp.Service).Str("Code", errResp.ResponseCode).Msg(errResp.Description)
		}

		commonresp.NewJsonResponse(c).SendError(errResp)
	}
}
