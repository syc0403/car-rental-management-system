package app

import (
	"car-rental-management-system/app/common/request"
	"car-rental-management-system/app/common/response"
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
