package app

import (
	"car-rental-management-system/app/common/request"
	"car-rental-management-system/app/common/response"
	"car-rental-management-system/app/models"
	"car-rental-management-system/app/services"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// GetAllRentOrder 获取全部订单信息

func GetAllRentOrder(c *gin.Context) {

	var pageQuery request.Pagination
	pageQuery.Current, _ = strconv.Atoi(c.Query("current"))
	pageQuery.PageSize, _ = strconv.Atoi(c.Query("pageSize"))
	rentOrder, total, err := services.RentOrderService.GetAllRentOrder(&pageQuery)
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
			"page":     pageQuery.Current,
			"pageSize": pageQuery.PageSize,
		}})
}

// GetRentOrderBy 根据条件获取订单信息

func GetRentOrderBy(c *gin.Context) {

	var query request.GetAllRentOrder
	if err := c.ShouldBindJSON(&query); err != nil {
		response.ValidateFail(c, request.GetErrorMsg(query, err))
		return
	}
	total, rentOrder, err := services.RentOrderService.GetRentOrderBy(&query)
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

// CreateRentIdentity 根据条件生成订单号

func CreateRentIdentity(c *gin.Context) {
	var query request.CreateRentIdentity
	query.CarNumber = c.Query("car_number")
	query.CustomerIdentity = c.Query("customer_identity")
	fmt.Println(query)
	rentOrderIdentity, err := services.RentOrderService.CreateRentIdentity(&query)
	if err != nil {
		response.BusinessFail(c, err.Error())
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code":    200,
		"message": "success",
		"data":    rentOrderIdentity})
}

// AddRentOrder 创建订单
func AddRentOrder(c *gin.Context) {
	var query models.RentOrder
	if err := c.ShouldBindJSON(&query); err != nil {
		response.ValidateFail(c, request.GetErrorMsg(query, err))
		return
	}
	if RentOrder, err := services.RentOrderService.AddRentOrder(&query); err != nil {
		response.BusinessFail(c, err.Error())
	} else {
		response.Success(c, RentOrder)
	}

}

// DeteleRentOrder 删除订单
func DeteleRentOrder(c *gin.Context) {
	RentOrderId := c.Query("id")
	rentOrde, err := services.RentOrderService.DeteleRentOrder(RentOrderId)
	if err != nil {
		response.BusinessFail(c, err.Error())
		return
	}
	response.Success(c, rentOrde)
}

// UpdateRentOrderStatus 修改出租单状态
func UpdateRentOrderStatus(c *gin.Context) {
	var query request.UpdateRentOrderStatus
	if err := c.ShouldBindJSON(&query); err != nil {
		response.ValidateFail(c, request.GetErrorMsg(query, err))
		return
	}

	if rentOrder, err := services.RentOrderService.UpdateRentOrder(&query); err != nil {
		response.BusinessFail(c, err.Error())
	} else {
		response.Success(c, rentOrder)
	}
}
