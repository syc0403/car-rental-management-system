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

// GetAllRole 获取全部角色

func GetAllRole(c *gin.Context) {

	var pageQuery request.Pagination
	pageQuery.Current, _ = strconv.Atoi(c.Query("current"))
	pageQuery.PageSize, _ = strconv.Atoi(c.Query("pageSize"))
	role, total, err := services.RoleService.GetAllRole(&pageQuery)
	if err != nil {
		response.BusinessFail(c, err.Error())
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code":    200,
		"message": "success",
		"data": map[string]interface{}{
			"data":     role,
			"total":    total,
			"page":     pageQuery.Current,
			"pageSize": pageQuery.PageSize,
		}})
}

// UpdateRole 修改角色信息
func UpdateRole(c *gin.Context) {
	var params models.Role
	if err := c.ShouldBindJSON(&params); err != nil {
		response.ValidateFail(c, request.GetErrorMsg(params, err))
		return
	}

	if role, err := services.RoleService.UpdateRole(&params); err != nil {
		response.BusinessFail(c, err.Error())
	} else {
		response.Success(c, role)
	}
}

// AddRole 创建角色
func AddRole(c *gin.Context) {
	var query models.Role
	if err := c.ShouldBindJSON(&query); err != nil {
		response.ValidateFail(c, request.GetErrorMsg(query, err))
		return
	}
	if role, err := services.RoleService.AddRole(&query); err != nil {
		response.BusinessFail(c, err.Error())
	} else {
		response.Success(c, role)
	}

}

// DeteleRole 删除角色
func DeteleRole(c *gin.Context) {
	RoleId := c.Query("id")
	role, err := services.RoleService.DeteleRole(RoleId)
	if err != nil {
		response.BusinessFail(c, err.Error())
		return
	}
	response.Success(c, role)
}
