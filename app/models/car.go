package models

import (
	"time"

	"gorm.io/gorm"
)

type Car struct {
	ID          int            `json:"id" gorm:"type:int(11);primaryKey"`
	CarNumber   string         `json:"car_number" gorm:"type:varchar(255);not null;comment:车牌号"`
	CarType     string         `json:"car_type" gorm:"type:varchar(255);comment:车辆类别"`
	Color       string         `json:"color" gorm:"type:varchar(255);comment:车辆颜色"`
	Price       string         `json:"price" gorm:"type:double(10,2);comment:车辆价格"`
	RentPrice   string         `json:"rent_price" gorm:"type:double(10,2);comment:租赁价格"`
	Deposit     string         `json:"deposit" gorm:"type:double(10,2);comment:押金"`
	IsRenting   int            `json:"is_renting" gorm:"type:int(1);default:0;comment:启用(0:未租出,1:已租出)"`
	Description string         `json:"description" gorm:"type:varchar(255);comment:简介"`
	ImgUrl      string         `json:"img_url" gorm:"type:varchar(255);comment:车辆图片"`
	CreatedAt   time.Time      `json:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at"`
	DeletedAt   gorm.DeletedAt `json:"deleted_at" gorm:"index"`
}
