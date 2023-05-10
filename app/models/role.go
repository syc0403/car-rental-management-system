package models

import (
	"time"

	"gorm.io/gorm"
)

type Role struct {
	ID          int            `json:"id" gorm:"type:int(11);primaryKey"`
	UserType    int            `json:"user_type" gorm:"type:int(1);not null;comment:权限(1:管理员,2:普通用户)"`
	Roles       string         `json:"roles" gorm:"type:varchar(255);not null;comment:权限"`
	Description string         `json:"description" gorm:"type:varchar(255);comment:描述"`
	CreatedAt   time.Time      `json:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at"`
	DeletedAt   gorm.DeletedAt `json:"deleted_at" gorm:"index"`
}
