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
		authRouter.POST("/user/info", app.Info)
		authRouter.POST("/user/logout", app.Logout)

		router.GET("/customer/getallcustomerinfo", app.GetAllCustomerInfo)
		router.GET("/customer/getcustomerinfobyid", app.GetCustomerInfoById)
		router.GET("/customer/getcustomerinfobyidentity", app.GetCustomerInfoByIdentity)
		router.POST("/customer/getcustomerinfobyname", app.GetCustomerInfoByName)
		router.GET("/customer/deletecustomer", app.DeteleCustomer)
		router.POST("/customer/updatecustomerinfo", app.UpdateCustomerInfo)
		router.POST("/customer/addcustomer", app.AddCustomer)
		router.GET("/customer/getcustomerbyrentorder", app.GetCustomerByRentOrder)

		router.GET("/car/getallcarinfo", app.GetAllCarInfo)
		router.GET("/car/getcarinfobyid", app.GetCarInfoById)
		router.GET("/car/deletecar", app.DeteleCar)
		router.POST("/car/updatecarinfo", app.UpdateCar)
		router.POST("/car/addcar", app.AddCar)
		router.POST("/car/getcarinfobynumber", app.GetCarInfoByNumber)

		router.GET("/rentOrder/getallrentorder", app.GetAllRentOrder)
		router.POST("/rentOrder/getrentorderby", app.GetRentOrderBy)
		router.GET("/rentOrder/createrentidentity", app.CreateRentIdentity)
		router.POST("/rentOrder/AddRentOrder", app.AddRentOrder)
		router.GET("/rentOrder/deleteRentOrder", app.DeteleRentOrder)
		router.GET("/rentOrder/getmoneybyuser", app.GetMoneyByUser)
		router.GET("/rentOrder/gettotalmoney", app.GetTotalMoney)

		router.GET("/check/getallcheck", app.GetAllCheck)
		router.POST("/check/addcheck", app.AddCheck)
		router.GET("/check/deletecheck", app.DeteleCheck)
		router.POST("/check/updaterentorderstatus", app.UpdateRentOrderStatus)
		router.POST("/rentOrder/getcheckby", app.GetCheckBy)
		router.POST("/rentOrder/updatecheck", app.UpdateCheck)
		router.GET("/check/getadresscount", app.GetAdressCount)
		router.GET("/check/getoccupationcount", app.GetOccupationCount)

	}

}
