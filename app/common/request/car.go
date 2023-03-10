package request

type GetCarInfoByNumber struct {
	Current   int    `json:"current" form:"page"`      // 页码
	PageSize  int    `json:"pageSize" form:"pageSize"` // 每页大小
	CarNumber string `json:"car_number"`
}
