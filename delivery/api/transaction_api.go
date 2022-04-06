package api

import (
	"github.com/gin-gonic/gin"
	"gokost.com/m/delivery/appresponse"
	"gokost.com/m/delivery/common_resp"
	"gokost.com/m/delivery/logger"
	"gokost.com/m/usecase"
	"net/http"
)

type transactionApi struct {
	usecaseInsert usecase.InsertTransactionUseCase
	usecaseUpdate usecase.UpdateTransactionUseCase
	usecaseSearch usecase.CustomerTransactionUseCase
}

func (t *transactionApi) StoreTransaction() gin.HandlerFunc {
	return func(c *gin.Context) {
		var requestData appresponse.TransactionRequest
		errBind := c.ShouldBindJSON(&requestData)
		if errBind != nil {
			logger.SendLogToDiscord("Insert Transaction", errBind)
			common_resp.NewCommonResp(c).FailedResp(http.StatusInternalServerError, common_resp.FailedMessage(errBind.Error()))
			return
		}

		err := t.usecaseInsert.InsertTransaction(requestData)
		if err != nil {
			logger.SendLogToDiscord("Insert Transaction", err)
			common_resp.NewCommonResp(c).FailedResp(http.StatusInternalServerError, common_resp.FailedMessage(err.Error()))
			return
		}
		common_resp.NewCommonResp(c).SuccessResp(http.StatusOK, common_resp.SuccessMessage("Success insert transaction", ""))
	}
}

func (t *transactionApi) UpdateTransaction() gin.HandlerFunc {
	return func(c *gin.Context) {
		var requestData appresponse.TransactionUpdateRequest
		errBind := c.ShouldBindJSON(&requestData)
		if errBind != nil {
			logger.SendLogToDiscord("Update Transaction", errBind)
			common_resp.NewCommonResp(c).FailedResp(http.StatusInternalServerError, common_resp.FailedMessage(errBind.Error()))
			return
		}

		err := t.usecaseUpdate.UpdateTransaction(requestData.CustomerId, requestData.Status)
		if err != nil {
			logger.SendLogToDiscord("Update transaction", err)
			common_resp.NewCommonResp(c).FailedResp(http.StatusInternalServerError, common_resp.FailedMessage(err.Error()))
			return
		}
		common_resp.NewCommonResp(c).SuccessResp(http.StatusOK, common_resp.SuccessMessage("Success update transaction", ""))
	}
}

func (t *transactionApi) GetTransactionByCustomerId() gin.HandlerFunc {
	return func(c *gin.Context) {
		var requestData appresponse.TransactionUpdateRequest
		errBind := c.ShouldBindJSON(&requestData)
		if errBind != nil {
			common_resp.NewCommonResp(c).FailedResp(http.StatusInternalServerError, common_resp.FailedMessage(errBind.Error()))
			logger.SendLogToDiscord("Get Transaction By Customer Id", errBind)
			return
		}

		data, err := t.usecaseSearch.SearchTransactionByCustomerId(requestData.CustomerId)
		if err != nil {
			logger.SendLogToDiscord("Search Transaction By Customer Id", err)
			common_resp.NewCommonResp(c).FailedResp(http.StatusInternalServerError, common_resp.FailedMessage(err.Error()))
			return
		}
		common_resp.NewCommonResp(c).SuccessResp(http.StatusOK, common_resp.SuccessMessage("Success get transaction by customer id", data))
	}
}

func NewTransactionApi(routerGroup *gin.RouterGroup, usecaseInsert usecase.InsertTransactionUseCase, usecaseUpdate usecase.UpdateTransactionUseCase, usecaseSearch usecase.CustomerTransactionUseCase) {
	api := &transactionApi{
		usecaseInsert,
		usecaseUpdate,
		usecaseSearch,
	}

	routerGroup.POST("/store_transaction", api.StoreTransaction())
	routerGroup.POST("/update_transaction", api.UpdateTransaction())
	routerGroup.GET("/search_transaction", api.GetTransactionByCustomerId())
}
