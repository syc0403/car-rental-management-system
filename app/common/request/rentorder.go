package request

type GetAllRentOrder struct {
	Current          int    `json:"current" form:"page"`      // 页码
	PageSize         int    `json:"pageSize" form:"pageSize"` // 每页大小
	Identity         string `json:"identity"`
	CarNumber        string `json:"car_number"`
	CustomerIdentity string `json:"customer_identity"`
	Status           int    `json:"status"`
}

type CreateRentIdentity struct {
	CarNumber        string `json:"car_number"`
	CustomerIdentity string `json:"customer_identity"`
}

type UpdateRentOrderStatus struct {
	Identity string `json:"identity"`
	Status   int    `json:"status"`
}
