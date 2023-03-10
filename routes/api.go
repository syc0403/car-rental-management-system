package routes

import (
	"car-rental-management-system/app/controllers/app"
	"car-rental-management-system/app/middleware"
	"car-rental-management-system/app/services"

	"github.com/gin-gonic/gin"
)

// SetApiGroupRoutes 定义 api 分组路由
func SetApiGroupRoutes(router *gin.RouterGroup) {

	router.POST("/user/register", app.Register)
	router.POST("/user/login", app.Login)

	authRouter := router.Group("").Use(middleware.JWTAuth(services.AppGuardName))
	{
		router.POST("/user/updateuserinfo", app.UpdateUserInfo)
		router.DELETE("/user/deleteuser", app.DeteleUser)
		router.GET("/user/getalluserinfo", app.GetAllUserInfo)
		router.GET("/customer/getallcustomerinfo", app.GetAllCustomerInfo)
		router.GET("/customer/getcustomerinfobyid", app.GetCustomerInfoById)
		router.POST("/customer/getcustomerinfobyname", app.GetCustomerInfoByName)
		router.GET("/customer/deletecustomer", app.DeteleCustomer)
		router.POST("/customer/updatecustomerinfo", app.UpdateCustomerInfo)
		router.POST("/customer/addcustomer", app.AddCustomer)

		router.GET("/car/getallcarinfo", app.GetAllCarInfo)
		router.GET("/car/getcarinfobyid", app.GetCarInfoById)
		router.GET("/car/deletecar", app.DeteleCar)
		router.POST("/car/updatecarinfo", app.UpdateCar)
		router.POST("/car/addcar", app.AddCar)
		router.POST("/car/getcarinfobynumber", app.GetCarInfoByNumber)

		authRouter.POST("/user/info", app.Info)
		authRouter.POST("/user/logout", app.Logout)
	}

}
