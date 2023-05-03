package app

import (
	"car-rental-management-system/app/common/request"
	"car-rental-management-system/app/common/response"
	"car-rental-management-system/app/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

// GetAllRentOrder 获取全部订单信息

func GetAllRentOrder(c *gin.Context) {

	var query request.GetAllRentOrder
	if err := c.ShouldBindJSON(&query); err != nil {
		response.ValidateFail(c, request.GetErrorMsg(query, err))
		return
	}
	total, rentOrder, err := services.RentOrderService.GetAllRentOrder(&query)
	if err != nil {
		response.BusinessFail(c, err.Error())
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code":    200,
		"message": "success",
		"data": map[string]interface{}{
			"data":     rentOrder,
			"total":    total,
			"page":     query.Current,
			"pageSize": query.PageSize,
		}})
}
