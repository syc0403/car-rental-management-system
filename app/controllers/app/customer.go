package app

import (
	"car-rental-management-system/app/common/request"
	"car-rental-management-system/app/common/response"
	"car-rental-management-system/app/models"
	"car-rental-management-system/app/services"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// GetAllCustomerInfo 获取全部客户信息

func GetAllCustomerInfo(c *gin.Context) {

	var pageQuery request.Pagination
	pageQuery.Current, _ = strconv.Atoi(c.Query("current"))
	pageQuery.PageSize, _ = strconv.Atoi(c.Query("pageSize"))
	err, total, customer := services.CustomerService.GetAllCustomerInfo(&pageQuery)
	if err != nil {
		response.BusinessFail(c, err.Error())
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code":    200,
		"message": "success",
		"data": map[string]interface{}{
			"data":     customer,
			"total":    total,
			"page":     pageQuery.Current,
			"pageSize": pageQuery.PageSize,
		}})
}

// GetCustomerInfoById 根据用户id获取客户信息
func GetCustomerInfoById(c *gin.Context) {
	customerId := c.Query("id")
	err, customer := services.CustomerService.GetCustomerInfoById(customerId)
	if err != nil {
		response.BusinessFail(c, err.Error())
		return
	} else {
		response.Success(c, customer)
	}
}

// AddCustomer 添加客户
func AddCustomer(c *gin.Context) {
	var query models.Customer
	if err := c.ShouldBindJSON(&query); err != nil {
		response.ValidateFail(c, request.GetErrorMsg(query, err))
		return
	}
	if err, customer := services.CustomerService.AddCustomer(&query); err != nil {
		response.BusinessFail(c, err.Error())
	} else {
		response.Success(c, customer)
	}

}

// DeteleCustomer 删除用户
func DeteleCustomer(c *gin.Context) {
	customerId := c.Query("id")
	err, customer := services.CustomerService.DeteleCustomer(customerId)
	if err != nil {
		response.BusinessFail(c, err.Error())
		return
	}
	response.Success(c, customer)
}

// UpdateCustomerInfo 修改客户信息
func UpdateCustomerInfo(c *gin.Context) {
	var customerInfo models.Customer
	if err := c.ShouldBindJSON(&customerInfo); err != nil {
		response.ValidateFail(c, request.GetErrorMsg(customerInfo, err))
		return
	}

	if err, customer := services.CustomerService.UpdateCustomerInfo(&customerInfo); err != nil {
		response.BusinessFail(c, err.Error())
	} else {
		response.Success(c, customer)
	}
}
