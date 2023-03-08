package request

type GetCustomerInfoByName struct {
	Current      int    `json:"current" form:"page"`      // 页码
	PageSize     int    `json:"pageSize" form:"pageSize"` // 每页大小
	CustomerName string `json:"customer_name"`
}
