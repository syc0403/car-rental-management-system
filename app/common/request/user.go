package request

type Register struct {
	LoginName string `form:"login_name" json:"login_name" binding:"required"`
	Phone     string `form:"phone" json:"phone" binding:"required,phone"`
	Password  string `form:"password" json:"password" binding:"required"`
}

type Login struct {
	Phone    string `form:"phone" json:"phone" binding:"required,phone"`
	Password string `form:"password" json:"password" binding:"required"`
}

func (register Register) GetMessages() ValidatorMessages {
	return ValidatorMessages{
		"name.required":     "用户名称不能为空",
		"phone.required":    "手机号码不能为空",
		"phone.phone":       "手机号码格式不正确",
		"password.required": "用户密码不能为空",
	}
}
func (login Login) GetMessages() ValidatorMessages {
	return ValidatorMessages{
		"phone.required":    "手机号码不能为空",
		"phone.phone":       "手机号码格式不正确",
		"password.required": "用户密码不能为空",
	}
}
