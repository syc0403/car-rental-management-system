package models

import (
	"time"

	"gorm.io/gorm"
)

type RentOrder struct {
	ID         int            `json:"id" gorm:"type:int(11);primaryKey"`
	CustomerId int32          `json:"customer_id" gorm:"comment:客户id"`
	OrderPrice string         `json:"order_price" gorm:"type:double(10,2);comment:租赁价格"`
	BeginDate  string         `json:"begin_date" gorm:"comment:开始时间"`
	ReturnDate string         `json:"return_date" gorm:"comment:归还时间"`
	Identity   string         `json:"identity" gorm:"varchar(255);comment:订单编号"`
	CarNumber  string         `json:"car_number" gorm:"type:varchar(255);not null;comment:车牌号"`
	Status     int            `json:"status" gorm:"type:int(1);default:0;comment:出租状态(0:未归还,1:已归还)"`
	CreatedAt  time.Time      `json:"created_at"`
	UpdatedAt  time.Time      `json:"updated_at"`
	DeletedAt  gorm.DeletedAt `json:"deleted_at" gorm:"index"`
}
