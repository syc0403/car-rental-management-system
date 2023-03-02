package models

import (
	"strconv"
)

type User struct {
	ID
	LoginName string `json:"login_name" gorm:"type:varchar(255);not null;comment:用户名称"`
	Password  string `json:"password" gorm:"type:varchar(255);not null;comment:用户密码"`
	Phone     string `json:"phone" gorm:"type:varchar(255);comment:用户手机号"`
	Identity  string `json:"identity" gorm:"type:varchar(255);comment:身份证"`
	RealName  string `json:"real_name" gorm:"type:varchar(20);comment:真实姓名"`
	Sex       int    `json:"sex" gorm:"type:int(1);default:0;comment:性别(0:男,1:女)"`
	Address   string `json:"address" gorm:"type:varchar(255);comment:地址"`
	Position  string `json:"position" gorm:"type:varchar(255);comment:职业"`
	Type      int    `json:"type" gorm:"type:int(1);default:2;comment:性别(1:管理员,2:普通用户)"`
	Available int    `json:"available" gorm:"type:int(1);default:1;comment:性别(0:停止,1:启用)"`
	Timestamps
	SoftDeletes
}

func (user User) GetUid() string {
	return strconv.Itoa(int(user.ID.ID))
}
