package response

type GetMoneyByUser struct {
	Money    float64 `json:"money"`
	OperName string  `json:"oper_name"`
}
