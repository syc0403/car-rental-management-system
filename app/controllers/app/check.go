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

// GetAllCheck 获取全部检查单

func GetAllCheck(c *gin.Context) {

	var pageQuery request.Pagination
	pageQuery.Current, _ = strconv.Atoi(c.Query("current"))
	pageQuery.PageSize, _ = strconv.Atoi(c.Query("pageSize"))
	check, total, err := services.CheckService.GetAllCheck(&pageQuery)
	if err != nil {
		response.BusinessFail(c, err.Error())
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code":    200,
		"message": "success",
		"data": map[string]interface{}{
			"data":     check,
			"total":    total,
			"page":     pageQuery.Current,
			"pageSize": pageQuery.PageSize,
		}})
}

// AddCheck 创建检查单
func AddCheck(c *gin.Context) {
	var query models.Check
	if err := c.ShouldBindJSON(&query); err != nil {
		response.ValidateFail(c, request.GetErrorMsg(query, err))
		return
	}
	if check, err := services.CheckService.AddCheck(&query); err != nil {
		response.BusinessFail(c, err.Error())
	} else {
		response.Success(c, check)
	}

}

// DeteleCheck 删除检查单
func DeteleCheck(c *gin.Context) {
	CheckId := c.Query("id")
	check, err := services.CheckService.DeteleCheck(CheckId)
	if err != nil {
		response.BusinessFail(c, err.Error())
		return
	}
	response.Success(c, check)
}
