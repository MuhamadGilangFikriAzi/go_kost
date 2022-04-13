package api

import (
	"github.com/gin-gonic/gin"
	"gokost.com/m/delivery/common_resp"
	"gokost.com/m/delivery/logger"
	"gokost.com/m/usecase"
	"net/http"
)

type productApi struct {
	usecase usecase.AllProductUseCase
}

func (p *productApi) GetAllProduct() gin.HandlerFunc {
	return func(c *gin.Context) {
		data, err := p.usecase.GetAllProduct()
		if err != nil {
			logger.SendLogToDiscord("Get All Customer", err)
			common_resp.NewCommonResp(c).FailedResp(http.StatusInternalServerError, common_resp.FailedMessage(err.Error()))
			return
		}
		common_resp.NewCommonResp(c).SuccessResp(http.StatusOK, common_resp.SuccessMessage("get all data customer", data))
	}
}

func NewProductApo(routeGroup *gin.RouterGroup, usecase usecase.AllProductUseCase) {
	api := &productApi{
		usecase,
	}

	routeGroup.GET("getall", api.GetAllProduct())
}
