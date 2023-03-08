package models

import (
	"time"

	"gorm.io/gorm"
)

type Customer struct {
	ID           int            `json:"id" gorm:"type:int(11);primaryKey"`
	CustomerName string         `json:"customer_name" gorm:"type:varchar(255);not null;comment:客户姓名"`
	Phone        string         `json:"phone" gorm:"type:varchar(255);comment:用户手机号"`
	Identity     string         `json:"identity" gorm:"type:varchar(255);comment:身份证"`
	Sex          int            `json:"sex" gorm:"type:int(1);default:0;comment:性别(0:男,1:女)"`
	Address      string         `json:"address" gorm:"type:varchar(255);comment:地址"`
	Occupation   string         `json:"occupation" gorm:"type:varchar(255);comment:职业"`
	ImgUrl       string         `json:"img_url" gorm:"type:varchar(255)"`
	CreatedAt    time.Time      `json:"created_at"`
	UpdatedAt    time.Time      `json:"updated_at"`
	DeletedAt    gorm.DeletedAt `json:"deleted_at" gorm:"index"`
}
