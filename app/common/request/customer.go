package request

type AddCustomer struct {
	CustomerName string `form:"customer_name" json:"customer_name"`
	Phone        string `form:"phone" json:"phone"`
	Identity     string `form:"identity" json:"identity"`
	Sex          int    `form:"sex" json:"sex"`
	Address      string `form:"address" json:"address"`
	Position     string `form:"position" json:"position"`
}
