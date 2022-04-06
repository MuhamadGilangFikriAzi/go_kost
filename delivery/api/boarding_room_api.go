package api

import (
	"github.com/gin-gonic/gin"
	"gokost.com/m/delivery/common_resp"
	"gokost.com/m/delivery/logger"
	"gokost.com/m/usecase"
	"net/http"
)

type boardingRoomApi struct {
	usecase usecase.AvailableRoomUseCase
}

func (b *boardingRoomApi) GetAllAvailableRoom() gin.HandlerFunc {
	return func(c *gin.Context) {
		data, err := b.usecase.AvailableRoom()
		if err != nil {
			logger.SendLogToDiscord("Get All Available Room", err)
			common_resp.NewCommonResp(c).FailedResp(http.StatusInternalServerError, common_resp.FailedMessage(err.Error()))
			return
		}
		common_resp.NewCommonResp(c).SuccessResp(http.StatusOK, common_resp.SuccessMessage("Get All Available Room", data))
	}
}

func NewBoardingApi(routerGroup *gin.RouterGroup, usecase usecase.AvailableRoomUseCase) {
	api := &boardingRoomApi{
		usecase,
	}

	routerGroup.GET("get_available_room", api.GetAllAvailableRoom())
}
