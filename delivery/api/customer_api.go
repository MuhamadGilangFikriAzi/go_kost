package api

import (
	"github.com/gin-gonic/gin"
	"gokost.com/m/delivery/common_resp"
	"gokost.com/m/usecase"
	"net/http"
)

type customerApi struct {
	usecase usecase.ListCustomerUseCase
}

func (c *customerApi) GetAllCustomer() gin.HandlerFunc {
	return func(g *gin.Context) {
		data, err := c.usecase.ListCustomer()
		if err != nil {
			common_resp.NewCommonResp(g).FailedResp(http.StatusInternalServerError, common_resp.FailedMessage(err.Error()))
			return
		}
		common_resp.NewCommonResp(g).SuccessResp(http.StatusOK, common_resp.SuccessMessage("get all data customer", data))
	}
}

func (c *customerApi) GetAllCustomerWithTransaction() gin.HandlerFunc {
	return func(g *gin.Context) {
		data, err := c.usecase.ListCustomerWithTransaction()
		if err != nil {
			common_resp.NewCommonResp(g).FailedResp(http.StatusInternalServerError, common_resp.FailedMessage(err.Error()))
			return
		}
		common_resp.NewCommonResp(g).SuccessResp(http.StatusOK, common_resp.SuccessMessage("get all data customer", data))
	}
}

func NewCustomerApi(routeGroup *gin.RouterGroup, useCase usecase.ListCustomerUseCase) {
	customer := &customerApi{
		useCase,
	}

	routeGroup.GET("getall", customer.GetAllCustomer())
	routeGroup.GET("getall/transaction", customer.GetAllCustomerWithTransaction())
}
