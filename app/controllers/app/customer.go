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

// GetCustomerInfoByName 根据用户姓名模糊查询客户信息
func GetCustomerInfoByName(c *gin.Context) {
	var query request.GetCustomerInfoByName
	if err := c.ShouldBindJSON(&query); err != nil {
		response.ValidateFail(c, request.GetErrorMsg(query, err))
		return
	}
	total, customer, err := services.CustomerService.GetCustomerInfoByName(&query)
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
			"page":     query.Current,
			"pageSize": query.PageSize,
		}})
}

// GetCustomerInfoByIdentity 根据身份证获取客户信息
func GetCustomerInfoByIdentity(c *gin.Context) {
	customerId := c.Query("identity")
	err, customer := services.CustomerService.GetCustomerInfoByIdentity(customerId)
	if err != nil {
		response.BusinessFail(c, err.Error())
		return
	} else {
		response.Success(c, customer)
	}
}

// GetCustomerByRentOrder 根据订单号获取客户信息
func GetCustomerByRentOrder(c *gin.Context) {
	rentOrderIdentity := c.Query("rentorder_identity")
	customer, err := services.CustomerService.GetCustomerByRentOrder(rentOrderIdentity)
	if err != nil {
		response.BusinessFail(c, err.Error())
		return
	} else {
		response.Success(c, customer)
	}
}

// GetAdressCount 查询不同地区的客户数量
func GetAdressCount(c *gin.Context) {
	res, err := services.CustomerService.GetAdressCount()
	if err != nil {
		response.BusinessFail(c, err.Error())
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code":    200,
		"message": "success",
		"data": map[string]interface{}{
			"data": res,
		}})
}

// GetOccupationCount 查询不同职业的客户数量
func GetOccupationCount(c *gin.Context) {
	res, err := services.CustomerService.GetOccupationCount()
	if err != nil {
		response.BusinessFail(c, err.Error())
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code":    200,
		"message": "success",
		"data": map[string]interface{}{
			"data": res,
		}})
}
