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

// GetAllCarInfo 获取全部车辆信息

func GetAllCarInfo(c *gin.Context) {

	var pageQuery request.Pagination
	pageQuery.Current, _ = strconv.Atoi(c.Query("current"))
	pageQuery.PageSize, _ = strconv.Atoi(c.Query("pageSize"))
	err, total, car := services.CarService.GetAllCarInfo(&pageQuery)
	if err != nil {
		response.BusinessFail(c, err.Error())
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code":    200,
		"message": "success",
		"data": map[string]interface{}{
			"data":     car,
			"total":    total,
			"page":     pageQuery.Current,
			"pageSize": pageQuery.PageSize,
		}})
}

// GetCarInfoById 根据用户id获取车辆信息
func GetCarInfoById(c *gin.Context) {
	carId := c.Query("id")
	err, car := services.CarService.GetCarInfoById(carId)
	if err != nil {
		response.BusinessFail(c, err.Error())
		return
	} else {
		response.Success(c, car)
	}
}

// AddCar 添加车辆
func AddCar(c *gin.Context) {
	var query models.Car
	if err := c.ShouldBindJSON(&query); err != nil {
		response.ValidateFail(c, request.GetErrorMsg(query, err))
		return
	}
	if err, Car := services.CarService.AddCar(&query); err != nil {
		response.BusinessFail(c, err.Error())
	} else {
		response.Success(c, Car)
	}

}

// DeteleCar 删除车辆
func DeteleCar(c *gin.Context) {
	carId := c.Query("id")
	err, car := services.CarService.DeteleCar(carId)
	if err != nil {
		response.BusinessFail(c, err.Error())
		return
	}
	response.Success(c, car)
}

// UpdateCar 修改车辆
func UpdateCar(c *gin.Context) {
	var carInfo models.Car
	if err := c.ShouldBindJSON(&carInfo); err != nil {
		response.ValidateFail(c, request.GetErrorMsg(carInfo, err))
		return
	}

	if err, car := services.CarService.UpdateCarInfo(&carInfo); err != nil {
		response.BusinessFail(c, err.Error())
	} else {
		response.Success(c, car)
	}
}

// GetCarInfoByNumber 根据车牌模糊查询车辆信息
func GetCarInfoByNumber(c *gin.Context) {
	var query request.GetCarInfoByNumber
	if err := c.ShouldBindJSON(&query); err != nil {
		response.ValidateFail(c, request.GetErrorMsg(query, err))
		return
	}
	err, total, car := services.CarService.GetCarInfoByNumber(&query)
	if err != nil {
		response.BusinessFail(c, err.Error())
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code":    200,
		"message": "success",
		"data": map[string]interface{}{
			"data":     car,
			"total":    total,
			"page":     query.Current,
			"pageSize": query.PageSize,
		}})
}
