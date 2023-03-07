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

// Register 用户注册
func Register(c *gin.Context) {
	var form request.Register
	if err := c.ShouldBindJSON(&form); err != nil {
		response.ValidateFail(c, request.GetErrorMsg(form, err))
		return
	}

	if err, user := services.UserService.Register(form); err != nil {
		response.BusinessFail(c, err.Error())
	} else {
		response.Success(c, user)
	}
}

func Info(c *gin.Context) {
	err, user := services.UserService.GetUserInfo(c.Keys["id"].(string))
	if err != nil {
		response.BusinessFail(c, err.Error())
		return
	}
	response.Success(c, user)
}

// UpdateUserInfo 修改用户信息
func UpdateUserInfo(c *gin.Context) {
	var userInfo models.User
	if err := c.ShouldBindJSON(&userInfo); err != nil {
		response.ValidateFail(c, request.GetErrorMsg(userInfo, err))
		return
	}

	if err, user := services.UserService.UpdateUserInfo(&userInfo); err != nil {
		response.BusinessFail(c, err.Error())
	} else {
		response.Success(c, user)
	}
}

// DeteleUser 删除用户
func DeteleUser(c *gin.Context) {
	userId := c.Query("id")
	err, user := services.UserService.DeteleUser(userId)
	if err != nil {
		response.BusinessFail(c, err.Error())
		return
	}
	response.Success(c, user)
}

// GetAllUserInfo 获取全部用户信息

func GetAllUserInfo(c *gin.Context) {

	var pageQuery request.Pagination
	pageQuery.Current, _ = strconv.Atoi(c.Query("current"))
	pageQuery.PageSize, _ = strconv.Atoi(c.Query("pageSize"))
	err, total, user := services.UserService.GetAllUserInfo(&pageQuery)
	if err != nil {
		response.BusinessFail(c, err.Error())
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code":    200,
		"message": "success",
		"data": map[string]interface{}{
			"data":     user,
			"total":    total,
			"page":     pageQuery.Current,
			"pageSize": pageQuery.PageSize,
		}})
}
