package api

import (
	"github.com/gin-gonic/gin"
	"gokost.com/m/delivery/appresponse"
	"gokost.com/m/delivery/common_resp"
	"gokost.com/m/usecase"
	"net/http"
)

type transactionApi struct {
	usecaseInsert usecase.InsertTransactionUseCase
	usecaseUpdate usecase.UpdateTransactionUseCase
}

func (t *transactionApi) StoreTransaction() gin.HandlerFunc {
	return func(c *gin.Context) {
		var requestData appresponse.TransactionRequest
		errBind := c.ShouldBindJSON(&requestData)
		if errBind != nil {
			common_resp.NewCommonResp(c).FailedResp(http.StatusInternalServerError, common_resp.FailedMessage(errBind.Error()))
			return
		}

		err := t.usecaseInsert.InsertTransaction(requestData)
		if err != nil {
			common_resp.NewCommonResp(c).FailedResp(http.StatusInternalServerError, common_resp.FailedMessage(err.Error()))
			return
		}
		common_resp.NewCommonResp(c).SuccessResp(http.StatusOK, common_resp.SuccessMessage("Success insert transaction", ""))
	}
}

func NewTransactionApi(routerGroup *gin.RouterGroup, usecaseInsert usecase.InsertTransactionUseCase, usecaseUpdate usecase.UpdateTransactionUseCase) {
	api := &transactionApi{
		usecaseInsert,
		usecaseUpdate,
	}

	routerGroup.POST("store_transaction", api.StoreTransaction())
}
