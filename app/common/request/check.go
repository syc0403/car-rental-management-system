package request

type GetCheckBy struct {
	Current           int    `json:"current" form:"page"`      // 页码
	PageSize          int    `json:"pageSize" form:"pageSize"` // 每页大小
	CheckIdentity     string `json:"identity"`
	RentOrderIdentity string `json:"rentorder_identity"`
}
